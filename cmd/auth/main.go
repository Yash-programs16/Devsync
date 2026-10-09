package main

import (
	"log"
	"net/http"

	"github.com/Yash-programs16/devsync/configs"
	"github.com/Yash-programs16/devsync/internal/auth"
	"github.com/Yash-programs16/devsync/internal/auth/handler"
	"github.com/Yash-programs16/devsync/internal/auth/repository"
	"github.com/Yash-programs16/devsync/internal/auth/service"
	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/email"
	"github.com/Yash-programs16/devsync/internal/token"

	"github.com/gorilla/mux"
)

func main() {

	// Config
	cfg, err := configs.Load()

	if err != nil {
		log.Fatal(err)
	}

	// Database
	db, err := database.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Repositories

	userRepo :=
		repository.NewUserRepository(
			db.Pool,
		)

	refreshRepo :=
		repository.NewRefreshTokenRepository(
			db.Pool,
		)

	emailVerificationRepo :=
		repository.NewEmailVerificationRepository(
			db.Pool,
		)

	// Token maker

	tokenMaker, err :=
		token.NewPasetoMaker(
			cfg.PasetoKey,
		)

	if err != nil {
		log.Fatal(err)
	}

	// Email Sender

	emailSender :=
		email.NewSMTPClient(
			cfg.SMTPHost,
			cfg.SMTPPort,
			cfg.SMTPUsername,
			cfg.SMTPPassword,
		)

	// Service

	authService :=
		service.NewAuthService(
			userRepo,
			refreshRepo,
			emailVerificationRepo,
			tokenMaker,
			emailSender,
		)

	// Handler

	authHandler :=
		handler.NewAuthHandler(
			authService,
		)

	// Router

	router :=
		mux.NewRouter()

	auth.RegisterRoutes(
		router,
		authHandler,
	)

	log.Println("server running on :8080")

	err = http.ListenAndServe(
		":8080",
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}
