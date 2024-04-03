package routes

import (
	"net/http"
	"user/controller"
	"user/middleware"

	"github.com/go-chi/chi"
)

func IntializeRouter(controllers *controller.UserController) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.HandleCORS)
		r.Post("/signup", controllers.Signup)
		r.Get("/auth/google", controllers.GoogleAuth)
		r.Post("/auth/google/callback", controllers.GoogleCallback)
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(405)
			w.Write([]byte("wrong method"))
		})
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Write([]byte("route does not exist"))
		})
	})

	return router
}
