package main

import (
	"log"
	"net/http"

	"github.com/Yash-programs16/devsync/configs"
	"github.com/Yash-programs16/devsync/internal/database"
	"github.com/Yash-programs16/devsync/internal/middleware"
	"github.com/Yash-programs16/devsync/internal/token"
	"github.com/Yash-programs16/devsync/internal/workspace/handler"
	"github.com/Yash-programs16/devsync/internal/workspace/repository"
	"github.com/Yash-programs16/devsync/internal/workspace/service"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := configs.Load()
	if err != nil {
		log.Fatal("config error:", err)
	}

	// Connect to PostgreSQL
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("database error:", err)
	}
	defer db.Close()

	// Token maker
	tokenMaker, err := token.NewPasetoMaker(cfg.PasetoKey)
	if err != nil {
		log.Fatal("token maker error:", err)
	}

	// Repositories
	workspaceRepo := repository.NewWorkspaceRepository(db)

	memberRepo := repository.NewWorkspaceMemberRepository(db)

	invitationRepo := repository.NewInvitationRepository(db)

	// Service
	workspaceService := service.NewWorkspaceService(
		workspaceRepo,
		memberRepo,
		invitationRepo,
	)

	// Handler
	workspaceHandler := handler.NewWorkspaceHandler(
		workspaceService,
	)

	// Router
	router := mux.NewRouter()

	protected := router.
		PathPrefix("/api/v1/workspaces").
		Subrouter()

	protected.Use(
		middleware.RequireAuth(tokenMaker),
	)

	// Create workspace
	protected.HandleFunc(
		"",
		workspaceHandler.CreateWorkspace,
	).Methods("POST")

	// List user's workspaces
	protected.HandleFunc(
		"",
		workspaceHandler.ListWorkspaces,
	).Methods("GET")

	// Get workspace
	protected.HandleFunc(
		"/{id}",
		workspaceHandler.GetWorkspace,
	).Methods("GET")

	// Update workspace
	protected.HandleFunc(
		"/{id}",
		workspaceHandler.UpdateWorkspace,
	).Methods("PATCH")

	// Delete workspace
	protected.HandleFunc(
		"/{id}",
		workspaceHandler.DeleteWorkspace,
	).Methods("DELETE")

	// List members
	protected.HandleFunc(
		"/{id}/members",
		workspaceHandler.ListMembers,
	).Methods("GET")

	// Get member
	protected.HandleFunc(
		"/{id}/members/{user_id}",
		workspaceHandler.GetMember,
	).Methods("GET")

	// Remove member
	protected.HandleFunc(
		"/{id}/members/{user_id}",
		workspaceHandler.RemoveMember,
	).Methods("DELETE")

	// Invitation
	protected.HandleFunc(
		"/{workspace_id}/invites",
		workspaceHandler.InviteMember,
	).Methods("POST")

	log.Println(
		"workspace service running on",
		cfg.WorkspaceHTTPPort,
	)

	if err := http.ListenAndServe(
		":"+cfg.WorkspaceHTTPPort,
		router,
	); err != nil {
		log.Fatal("server error:", err)
	}
}
