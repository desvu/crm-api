package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/auth/session"
	"github.com/qilin/store-api/pkg/apierror"
	"go.uber.org/zap"
)

// context key
type contextKey struct{ name string }

var userCtxKey = &contextKey{"user"}
var fingerprintCtxKey = &contextKey{"fp"}

// ExtractUserContext get user data from context
func ExtractUserContext(ctx context.Context) *User {
	if user, ok := ctx.Value(userCtxKey).(*User); ok {
		return user
	}
	return &User{}
}

// SetUserContext sets user context into http.Request
// we need it because gqlgen don't knows about echo.Context and uses http.Request context
// but there is no easy way to set Request context.
func SetUserContext(ctx echo.Context, user *User) {
	r := ctx.Request()
	newctx := context.WithValue(r.Context(), userCtxKey, user)
	ctx.SetRequest(r.WithContext(newctx))
}

// ExtractFingerprintContext get fingerprint (ip, ua, hwid) from context
func ExtractFingerprintContext(ctx context.Context) *Fingerprint {
	if fp, ok := ctx.Value(fingerprintCtxKey).(*Fingerprint); ok {
		return fp
	}
	return &Fingerprint{}
}

// SetFingerprintContext sets fingerprint (ip, ua, hwid) into http.Request
func SetFingerprintContext(ctx echo.Context, fp *Fingerprint) {
	r := ctx.Request()
	newctx := context.WithValue(r.Context(), fingerprintCtxKey, fp)
	ctx.SetRequest(r.WithContext(newctx))
}

// Middleware returns authorization middleware for http server
func (a *Auth) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		claims, err := a.getSession(ctx)
		if err == http.ErrNoCookie {
			// not authorized
			return next(ctx)
		}

		if err == session.ErrSessionNotFound {
			return apierror.InvalidSession
		}

		if err != nil {
			a.logger.Error("failed to get session", zap.Error(err))
			return next(ctx)
		}

		if claims.UserID == 0 {
			return next(ctx)
		}

		SetUserContext(ctx, &User{
			ID: claims.UserID,
		})
		SetFingerprintContext(ctx, &Fingerprint{
			UA:   ctx.Request().UserAgent(),
			IP:   ctx.RealIP(),
			HWID: "", // todo
		})

		a.logger.Debug("user authorized", zap.Int("user_id", claims.UserID))

		return next(ctx)
	}
}
