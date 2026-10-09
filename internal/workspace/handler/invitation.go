package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Yash-programs16/devsync/internal/middleware"
	"github.com/Yash-programs16/devsync/internal/workspace/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type createInvitationRequest struct {
	Email string `json:"email"`
}

func (h *WorkspaceHandler) InviteMember(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get authenticated user.
	inviterID, ok := middleware.UserIDFromContext(
		r.Context(),
	)

	if !ok || inviterID == uuid.Nil {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "authentication required",
			},
		)
		return
	}

	// Get workspace ID from URL.
	vars := mux.Vars(r)

	workspaceID, err := uuid.Parse(
		vars["workspace_id"],
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

	// Decode request body.
	var req createInvitationRequest

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

	// Call service.
	invite, _, err := h.service.InviteMember(
		r.Context(),
		service.CreateInvitationRequest{
			WorkspaceID: workspaceID,
			InviterID:   inviterID,
			Email:       req.Email,
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
		invite,
	)
}
