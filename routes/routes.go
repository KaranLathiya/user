package routes

import (
	"net/http"
	"go-structure/controller"

	"github.com/go-chi/chi"
)

func IntializeRouter(controllers *controller.UserController) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		// r.Use(middleware.HandleCORS)
		r.Post("/signup", controllers.CreateUserHandler)

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
