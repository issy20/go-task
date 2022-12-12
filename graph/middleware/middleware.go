package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/issy20/go-task/graph/domain/model"
	"github.com/issy20/go-task/graph/pkg/auth"
	"github.com/issy20/go-task/graph/pkg/db"
)

type userContextKey string

var key = userContextKey("user")

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := auth.ExtractToken(r)
		if tokenStr == "" {
			next.ServeHTTP(w, r)
			return
		}

		user_id, err := auth.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		db := db.InitDB()
		user := &model.User{ID: user_id}

		if err := db.Debug().First(&user).Error; err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), key, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *model.User {
	raw := ctx.Value(key)
	if raw == nil {
		log.Println("Cannot find key")
		return nil
	}

	return raw.(*model.User)
}
