package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"database/sql"
	"mini-social-network-api/config"
	"mini-social-network-api/internal/middleware"
	v1follow "mini-social-network-api/internal/v1/follow"
	v1post "mini-social-network-api/internal/v1/post"
	v1user "mini-social-network-api/internal/v1/user"
)

func NewRouter(cfg *config.Config, db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestLoggerMiddleware)

	// Dependency injection
	userRepo := v1user.NewRepository(db)
	userSvc := v1user.NewService(userRepo)
	userHandler := v1user.NewHandler(userSvc)

	postRepo := v1post.NewRepository(db)
	postSvc := v1post.NewService(postRepo)
	postHandler := v1post.NewHandler(postSvc)

	followRepo := v1follow.NewRepository(db)
	followSvc := v1follow.NewService(followRepo)
	followHandler := v1follow.NewHandler(followSvc)

	// Base health route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Versioned API
	r.Route("/v1", func(r chi.Router) {
		// Public
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)

		// Protected
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuthMiddleware)

			r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
				userID := r.Context().Value(middleware.ContextUserIDKey).(int64)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]interface{}{
					"user_id": userID,
				})
			})

			r.Post("/posts", postHandler.CreatePost)
			r.Post("/follow", followHandler.Follow)
		})
	})

	return r
}
