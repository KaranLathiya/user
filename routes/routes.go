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

		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", controllers.Signup)
			r.Post("/login", controllers.Login)
			r.Route("/google", func(r chi.Router) {
				r.Get("/", controllers.GoogleAuth)
				r.Get("/login", controllers.GoogleLogin)
			})
		})

		r.Post("/otp/verify", controllers.VerifyOTP)

		r.Route("/user/profile", func(r chi.Router) {
			r.Use(middleware.Authentication)
			r.Get("/", controllers.GetCurrentUserDetails)
			r.Put("/privacy", controllers.UpdateUserPrivacy)
			r.Put("/basic", controllers.UpdateBasicDetails)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(middleware.Authentication)
			r.Get("/", controllers.GetUserList)
			r.Get("/{id}/id", controllers.GetUserDetailsByID)
			r.Get("/{username}/username", controllers.GetUserDetailsByUsername)

			r.Route("/block", func(r chi.Router) {
				r.Get("/", controllers.BlockedUserList)
				r.Post("/", controllers.BlockUser)
			})

			r.Delete("/{blocked}/unblock", controllers.UnblockUser)

		})

		r.Route("/internal", func(r chi.Router) {
			r.Post("/users/details", controllers.GetUsersDetailsByIDs)
			r.Post("/user/otp", controllers.CreateOTPForDeleteOrganization)
			r.Post("/otp/verify", controllers.VerifyOTPForDeleteOrganization)
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
