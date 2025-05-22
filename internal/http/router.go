package http

import (
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

	userRepo := v1user.NewRepository(db)
	userSvc := v1user.NewService(userRepo)
	userHandler := v1user.NewHandler(userSvc)

	postRepo := v1post.NewRepository(db)
	postSvc := v1post.NewService(postRepo)
	postHandler := v1post.NewHandler(postSvc)

	followRepo := v1follow.NewRepository(db)
	followSvc := v1follow.NewService(followRepo)
	followHandler := v1follow.NewHandler(followSvc)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// User routes
	r.Route("/v1/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})

	// Post routes
	r.Route("/v1/posts", func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware)
		r.Post("/", postHandler.CreatePost)
		r.Put("/{id}", postHandler.UpdatePost)
		r.Delete("/{id}", postHandler.DeletePost)
	})

	// Follow routes
	r.Route("/v1/follows", func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware)
		r.Post("/", followHandler.Follow)
	})

	return r
}
