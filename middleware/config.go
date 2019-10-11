package middleware

import (
	"context"
	"net/http"
)

type Settings struct {
	Server struct {
		Host string
		Port string
	}
	Logger struct {
		Path  string
		Level int
	}

	Github struct {
		Owner string
		Repo  string
	}

	Site struct {
		Title string
	}

	Redis struct {
		Host     string
		Port     string
		Database int
		Password string
	}
}

func SetConfig(ctx context.Context, settings *Settings) context.Context {
	return context.WithValue(ctx, "settings", settings)
}

func ConfigMiddleWare(settings *Settings) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := SetConfig(r.Context(), settings)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
