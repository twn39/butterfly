package middleware

import (
	"context"
	"github.com/twn39/butterfly/github"
	"net/http"
)

func SetGithubClient(ctx context.Context, client *github.Client) context.Context {
	return context.WithValue(ctx, "github", client)
}

func GithubClientClientMiddleWare(owner string, repo string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client := github.NewGithubClient(owner, repo)
			ctx := SetGithubClient(r.Context(), client)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
