package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Yash-programs16/devsync/internal/middleware"
	"github.com/Yash-programs16/devsync/internal/workspace/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type WorkspaceHandler struct {
	service *service.WorkspaceService
}

func NewWorkspaceHandler(
	service *service.WorkspaceService,
) *WorkspaceHandler {
	return &WorkspaceHandler{
		service: service,
	}
}

type createWorkspaceRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
}

type updateWorkspaceRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
}

func (h *WorkspaceHandler) CreateWorkspace(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	var req createWorkspaceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	workspace, err := h.service.CreateWorkspace(
		r.Context(),
		service.CreateWorkspaceRequest{
			Name:        req.Name,
			Slug:        req.Slug,
			Description: req.Description,
			OwnerID:     userID,
		},
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		workspace,
	)
}

func (h *WorkspaceHandler) ListWorkspaces(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	workspaces, err := h.service.ListWorkspaces(
		r.Context(),
		userID,
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{
				"error": "failed to get workspaces",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		workspaces,
	)
}

func (h *WorkspaceHandler) GetWorkspace(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	workspace, err := h.service.GetWorkspace(
		r.Context(),
		id,
		userID,
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": "workspace not found",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		workspace,
	)
}

func (h *WorkspaceHandler) UpdateWorkspace(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	var req updateWorkspaceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	workspace, err := h.service.UpdateWorkspace(
		r.Context(),
		id,
		userID,
		service.UpdateWorkspaceRequest{
			Name:        req.Name,
			Slug:        req.Slug,
			Description: req.Description,
		},
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		workspace,
	)
}

func (h *WorkspaceHandler) DeleteWorkspace(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	if err := h.service.DeleteWorkspace(
		r.Context(),
		id,
		userID,
	); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "workspace deleted successfully",
		},
	)
}

func (h *WorkspaceHandler) ListMembers(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	workspaceID, err := uuid.Parse(
		vars["id"],
	)
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	members, err := h.service.ListMembers(
		r.Context(),
		workspaceID,
		userID,
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": "workspace not found",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		members,
	)
}

func (h *WorkspaceHandler) GetMember(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	workspaceID, err := uuid.Parse(vars["id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	memberUserID, err := uuid.Parse(vars["user_id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid user id",
			},
		)
		return
	}

	member, err := h.service.GetMember(
		r.Context(),
		workspaceID,
		userID,
		memberUserID,
	)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": "member not found",
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		member,
	)
}

func (h *WorkspaceHandler) RemoveMember(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || userID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	vars := mux.Vars(r)

	workspaceID, err := uuid.Parse(vars["id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid workspace id",
			},
		)
		return
	}

	memberUserID, err := uuid.Parse(vars["user_id"])
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid user id",
			},
		)
		return
	}

	if err := h.service.RemoveMember(
		r.Context(),
		workspaceID,
		userID,
		memberUserID,
	); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "member removed successfully",
		},
	)
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
