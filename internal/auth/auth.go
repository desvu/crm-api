package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/qilin/crm-api/internal/auth/session"

	"github.com/coreos/go-oidc"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Config struct {
	OAuth2 struct {
		Provider     string `required:"true"`
		ClientId     string `required:"true"`
		ClientSecret string `required:"true"`
		RedirectUrl  string `required:"true"`
	}
	LogoutRedirect string // optional if provider doesn't support it in openid-configuration
	LogoutCallback string

	// cookies rules
	SessionCookieName string `default:"ssid"`
	Domain            string
	SecureCookie      bool

	AutoSignIn         bool
	Secret             string
	SuccessRedirectURL string

	Sessions session.Config
}

type Auth struct {
	cfg         Config
	oauth2      *oauth2.Config
	provider    *oidc.Provider
	verifier    *oidc.IDTokenVerifier
	stateSecret []byte
	session     *session.Session
	logger      *zap.Logger
}

// New
func New(cfg *Config) (*Auth, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://id.tst.qilin.super.com/")
	if err != nil {
		return nil, err
	}
	// advanced claims from openid configuration
	var claims struct {
		EndSessionURL string `json:"end_session_endpoint"`
	}
	if err := provider.Claims(&claims); err != nil {
		return nil, err
	}
	if claims.EndSessionURL != "" {
		cfg.LogoutRedirect = claims.EndSessionURL
	}

	return &Auth{
		cfg:      *cfg,
		provider: provider,
		oauth2: &oauth2.Config{
			RedirectURL:  cfg.OAuth2.RedirectUrl,
			ClientID:     cfg.OAuth2.ClientId,
			ClientSecret: cfg.OAuth2.ClientSecret,
			Scopes:       []string{"openid"},
			Endpoint:     provider.Endpoint(),
		},
		verifier: provider.Verifier(&oidc.Config{
			ClientID: cfg.OAuth2.ClientId,
		}),
		stateSecret: []byte(cfg.Secret),
		session:     session.New(&cfg.Sessions),
		logger:      zap.L().Named("auth"),
	}, nil
}

// Session ====================================================================

func (a *Auth) checkAuthorized(c echo.Context) (*session.SessionClaims, bool) {
	claims, err := a.getSession(c)
	if err == session.ErrSessionNotFound {
		return nil, false
	}
	if err != nil {
		// TODO log error
		return nil, false
	}

	if claims.UserID == 0 {
		return nil, false
	}

	return claims, true
}

// getSession returns current session claims
// errors http.ErrNoCookie session.ErrSessionNotFound
func (a *Auth) getSession(c echo.Context) (*session.SessionClaims, error) {
	cookie, err := c.Cookie(a.cfg.SessionCookieName)
	if err != nil {
		return nil, err
	}
	if cookie.Value == "" {
		return nil, http.ErrNoCookie
	}
	return a.session.Get(cookie.Value)
}

func (a *Auth) removeSessionCookie(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     a.cfg.SessionCookieName,
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
		Domain:   a.cfg.Domain,
		Path:     "/",
		Secure:   a.cfg.SecureCookie,
	})
	return nil
}

func (a *Auth) removeSession(c echo.Context) error {
	cookie, err := c.Cookie(a.cfg.SessionCookieName)
	a.removeSessionCookie(c)
	if err == nil && cookie.Value != "" {
		a.session.Drop(cookie.Value)
	}
	return nil
}

func (a *Auth) setSession(c echo.Context, sessID string, claims *session.SessionClaims) error {
	// special cookie for front logic
	c.SetCookie(&http.Cookie{
		Name:     "has_session",
		Value:    "true",
		Domain:   a.cfg.Domain,
		Path:     "/",
		Secure:   a.cfg.SecureCookie,
		SameSite: http.SameSiteStrictMode,
	})

	c.SetCookie(&http.Cookie{
		Name:     a.cfg.SessionCookieName,
		Value:    sessID,
		HttpOnly: true,
		Domain:   a.cfg.Domain,
		Path:     "/",
		Secure:   a.cfg.SecureCookie,
		SameSite: http.SameSiteLaxMode,
		// MaxAge:   0,
	})
	return a.session.Set(sessID, claims)
}

// State Cookie =============================================================

func (a *Auth) removeStateCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    "",
		Domain:   a.cfg.Domain,
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/",
		Secure:   a.cfg.SecureCookie,
	})
}

func (a *Auth) setStateCookie(c echo.Context, value string) {
	c.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    value,
		Domain:   a.cfg.Domain,
		MaxAge:   int((30 * time.Minute).Seconds()),
		HttpOnly: true,
		Path:     "/",
		Secure:   a.cfg.SecureCookie,
		SameSite: http.SameSiteNoneMode,
	})
}

func (a *Auth) validateState(c echo.Context, code string) error {
	var cookie *http.Cookie
	for _, cke := range c.Cookies() {
		if cke.Name == "state" && cke.Value != "" {
			cookie = cke
		}
	}
	if cookie == nil {
		return errors.New("invalid auth state")
	}

	if code != a.secureStateCode(cookie.Value) {
		return errors.New("invalid auth state")
	}

	return nil
}

func (a *Auth) secureStateCode(code string) string {
	var h = sha256.New()
	h.Write([]byte(code))
	h.Write(a.stateSecret)
	return hex.EncodeToString(h.Sum(nil))
}
