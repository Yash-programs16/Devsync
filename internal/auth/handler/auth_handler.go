package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Yash-programs16/devsync/internal/auth/service"
	"github.com/Yash-programs16/devsync/internal/models"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(
	authService *service.AuthService,
) *AuthHandler {

	return &AuthHandler{
		authService: authService,
	}
}
func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.RegisterRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	user, err := h.authService.Register(
		r.Context(),
		req,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusCreated,
	)

	json.NewEncoder(w).Encode(
		user,
	)
}
func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.LoginRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	response, err := h.authService.Login(
		r.Context(),
		req,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusUnauthorized,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		response,
	)
}
func (h *AuthHandler) Refresh(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.RefreshRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	response, err := h.authService.Refresh(
		r.Context(),
		req.RefreshToken,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusUnauthorized,
		)

		return
	}

	json.NewEncoder(w).Encode(
		response,
	)
}
func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.RefreshRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	err = h.authService.Logout(
		r.Context(),
		req.RefreshToken,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}
func (h *AuthHandler) VerifyEmail(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req models.VerifyEmailRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	err = h.authService.VerifyEmail(
		r.Context(),
		req,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "email verified successfully",
		},
	)
}
