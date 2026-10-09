package auth

import (
	"github.com/Yash-programs16/devsync/internal/auth/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes(
	router *mux.Router,
	authHandler *handler.AuthHandler,
) {

	router.HandleFunc(
		"/api/v1/auth/register",
		authHandler.Register,
	).Methods("POST")

	router.HandleFunc(
		"/api/v1/auth/login",
		authHandler.Login,
	).Methods("POST")

	router.HandleFunc(
		"/api/v1/auth/refresh",
		authHandler.Refresh,
	).Methods("POST")

	router.HandleFunc(
		"/api/v1/auth/logout",
		authHandler.Logout,
	).Methods("POST")
	router.HandleFunc(
		"/api/v1/auth/verify-email",
		authHandler.VerifyEmail,
	).Methods("POST")

}
