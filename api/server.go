package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/spinales/quiz-api/token"
	"github.com/spinales/quiz-api/util"
	"gorm.io/gorm"
)

type Server struct {
	// store      *gorm.DB
	config     *util.Config
	tokenMaker token.Maker
	router     *chi.Mux
	service    *service
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *gorm.DB, config *util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		service:    NewService(store),
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"https://*", "http://*"},
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// }))

	router.Use(cors.AllowAll().Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Post("/login", server.login)
	router.Post("/register", server.register)

	router.Route("/api", func(r chi.Router) {
		// r.Use(authMiddleware2(r, server.tokenMaker))
		r.Use(server.authMiddleware2)
		r.Route("/question", func(r chi.Router) {
			r.Get("/", server.GetQuestions)
			r.Post("/", server.AddQuestion)
			r.Route("/{id:^[0-9]}", func(r chi.Router) {
				r.Get("/", server.GetQuestion)
				r.Delete("/", server.DeleteQuestion)
				r.Put("/", server.UpdateQuestion)
			})
		})
		r.Route("/answer", func(r chi.Router) {
			r.Get("/", server.GetAnswers)
			r.Post("/", server.AddAnswer)
			r.Route("/{id:^[0-9]}", func(r chi.Router) {
				r.Get("/", server.GetAnswer)
				r.Delete("/", server.DeleteAnswer)
				r.Put("/", server.UpdateAnswer)
			})
		})
	})

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return http.ListenAndServe(address, server.router)
}
