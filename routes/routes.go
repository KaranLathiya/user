package routes

import (
	"net/http"
	"user/controller"
	"user/middleware"

	"github.com/go-chi/chi"
)

func InitializeRouter(controllers *controller.UserController) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.HandleCORS)
		r.Post("/signup", controllers.Signup)
		r.Post("/login", controllers.Login)

		r.Route("/auth", func(r chi.Router) {
			r.Get("/google", controllers.GoogleAuth)
			r.Get("/google/callback", controllers.GoogleCallback)
		})

		r.Post("/otp/verify", controllers.VerifyOTP)

		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.Authentication)
			r.Route("/block", func(r chi.Router) {
				r.Get("/", controllers.BlockUserList)
				r.Post("/", controllers.BlockUser)
				r.Delete("/{blocked}", controllers.UnblockUser)
			})
			r.Route("/profile", func(r chi.Router) {
				r.Put("/privacy", controllers.UpdateUserPrivacy)
				r.Put("/name-details", controllers.UpdateUserNameDetails)
			})
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/username", controllers.GetUsernameByID)
				r.Get("/user-details", controllers.GetUserDetailsByID)
			})
			r.Get("/user-list",controllers.GetUserList)
		})

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
