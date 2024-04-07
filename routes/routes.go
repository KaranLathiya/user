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
		r.Post("/login", controllers.Login)

		r.Route("/auth", func(r chi.Router) {
			r.Get("/google", controllers.GoogleAuth)
			r.Post("/google/callback", controllers.GoogleCallback)
		})

		r.Post("/otp/verify", controllers.VerifyOTP)
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
