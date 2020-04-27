package auth

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/qilin/crm-api/internal/auth/session"
	"github.com/qilin/store-api/pkg/apierror"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	redirectURLParam = "redirect"
)

type claims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PhoneNumber   string `json:"phone_number"`
	PhoneVerified bool   `json:"phone_verified"`
	Picture       string `json:"picture"`
	Blocked       bool   `json:"blocked"`
}

func (a *Auth) InitRoutes(g *echo.Group) {
	g.GET("/login", a.login)
	g.GET("/callback", a.callback)
	g.GET("/logout", a.logout)
	g.GET("/logoutcb", a.logoutCallback)
}

func (a *Auth) login(c echo.Context) error {
	rURL := c.QueryParam(redirectURLParam)

	if _, ok := a.checkAuthorized(c); ok {
		return a.redirectSuccess(c, rURL, nil)
	}

	stateCode := uuid.New().String()
	a.setStateCookie(c, stateCode)
	state, err := EncodeState(&State{Redirect: rURL, Code: a.secureStateCode(stateCode)})
	if err != nil {
		return a.redirectSuccess(c, rURL, err)
	}

	var opts = make([]oauth2.AuthCodeOption, 0, 10)
	if c.QueryParam("prompt") == "none" {
		opts = append(opts, oauth2.SetAuthURLParam("prompt", "none"))
	}

	var url = a.oauth2.AuthCodeURL(state, opts...)
	return c.Redirect(http.StatusFound, url)
}

func (a *Auth) logout(c echo.Context) error {
	claims, ok := a.checkAuthorized(c)
	err := a.removeSession(c)
	if err != nil {
		a.logger.Warn("Can't destroy session", zap.Error(err))
	}

	// TODO add action to log history here or leave all history in auth1

	if !ok {
		// redirect to sso logout url
		// without id token it's not possible to define post logout redirect url
		return c.Redirect(http.StatusTemporaryRedirect, a.cfg.LogoutRedirect)
	}

	// build logout url
	u, _ := url.Parse(a.cfg.LogoutRedirect)
	var query = u.Query()
	if claims.IDToken != "" {
		query.Set("id_token_hint", claims.IDToken)

		query.Set("post_logout_redirect_uri", a.cfg.LogoutCallback)
		st, err := EncodeState(&State{Redirect: c.QueryParam("redirect")})
		if err != nil {
			return err
		}
		query.Set("state", st)
	}

	u.RawQuery = query.Encode()

	return c.Redirect(http.StatusTemporaryRedirect, u.String())
}

func (a *Auth) logoutCallback(c echo.Context) error {
	var st State
	if err := DecodeState(c.QueryParam("state"), &st); err != nil {
		return err
	}

	if st.Redirect != "" {
		return c.Redirect(http.StatusTemporaryRedirect, st.Redirect)
	}
	return c.Redirect(http.StatusTemporaryRedirect, a.cfg.SuccessRedirectURL)
}

func (a *Auth) redirectSuccess(c echo.Context, url string, err error) error {
	if url == "" {
		url = a.cfg.SuccessRedirectURL
	}

	if err != nil {
		a.logger.Error("Authentication failed", zap.Error(err))
		url, err = urlWithParams(url, "error", "internal_error")
		if err != nil {
			a.logger.Error("Redirect url parsing failed", zap.Error(err))
		}
	}

	return c.Redirect(http.StatusFound, url)
}

func (a *Auth) callback(c echo.Context) error {
	// Verify state param
	var state State
	err := DecodeState(c.FormValue("state"), &state)
	if err != nil {
		return a.redirectSuccess(c, "", err)
	}
	a.logger.Debug("Oauth callback state", zap.Object("state", state))
	if err := a.validateState(c, state.Code); err != nil {
		return a.redirectSuccess(c, "", err)
	}
	a.removeStateCookie(c)

	err = a.callbackError(c)
	return a.redirectSuccess(c, state.Redirect, err)
}

func (a *Auth) callbackError(c echo.Context) error {

	if c.FormValue("error") != "" {
		return nil
	}

	// exchange code to tokens
	var code = c.FormValue("code")
	var ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	oauthToken, err := a.oauth2.Exchange(ctx, code)
	if err != nil {
		return apierror.AuthCodeExchangeFail
	}

	// parse and verify id_token
	rawIDToken, ok := oauthToken.Extra("id_token").(string)
	if !ok {
		return apierror.IdTokenNotProvided
	}

	idToken, err := a.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return apierror.IdTokenVerificationFail
	}

	a.logger.Debug("Authenticate user", zap.String("ext_user_id", idToken.Subject))

	var claims claims
	if err := idToken.Claims(&claims); err != nil {
		return err
	}

	fp := &Fingerprint{
		UA:   c.Request().UserAgent(),
		IP:   c.RealIP(),
		HWID: "", // todo
	}

	uid, err := a.authorizeUser(c.Request().Context(), idToken.Subject, claims, fp)
	if err != nil {
		return err
	}

	a.logger.Info("User logged in", zap.Int("user_id", uid))

	if err := a.setSession(c, uuid.New().String(), &session.SessionClaims{
		UserID:     uid,
		ExternalID: idToken.Subject,
		CreatedAt:  time.Now(),
		IDToken:    rawIDToken,
	}); err != nil {
		return apierror.SessionSetFailed
	}

	// TODO add action to log history here or leave all history in auth1

	return err
}

func (a *Auth) authorizeUser(ctx context.Context, externalId string, claims claims, fp *Fingerprint) (int, error) {
	// TODO implementation
	return 0, nil
}

func urlWithParams(uri string, params ...string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return uri, err
	}
	q := u.Query()
	for k, v := 0, 1; v < len(params); k, v = k+2, v+2 {
		q.Set(params[k], params[v])
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
