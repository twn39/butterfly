package middleware

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/twn39/butterfly/cache"
	"net/http"
)

func SetCache(ctx context.Context, cache cache.Cache) context.Context {
	return context.WithValue(ctx, "cache", cache)
}

func CacheMiddleWare(client *redis.Client) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := cache.NewRedisCache(client)
			ctx := SetCache(r.Context(), c)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
