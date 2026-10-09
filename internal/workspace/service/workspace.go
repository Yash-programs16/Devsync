package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/Yash-programs16/devsync/internal/workspace/invitation"
	"github.com/Yash-programs16/devsync/internal/workspace/repository"
	"github.com/google/uuid"
)

type WorkspaceService struct {
	repo           *repository.WorkspaceRepository
	memberRepo     *repository.WorkspaceMemberRepository
	invitationRepo *repository.InvitationRepository
}

func NewWorkspaceService(
	repo *repository.WorkspaceRepository,
	memberRepo *repository.WorkspaceMemberRepository,
	invitationRepo *repository.InvitationRepository,
) *WorkspaceService {
	return &WorkspaceService{
		repo:           repo,
		memberRepo:     memberRepo,
		invitationRepo: invitationRepo,
	}
}

type CreateWorkspaceRequest struct {
	Name        string
	Slug        string
	Description *string
	OwnerID     uuid.UUID
}

type UpdateWorkspaceRequest struct {
	Name        string
	Slug        string
	Description *string
}

type CreateInvitationRequest struct {
	WorkspaceID uuid.UUID
	InviterID   uuid.UUID
	Email       string
}

func (s *WorkspaceService) CreateWorkspace(
	ctx context.Context,
	req CreateWorkspaceRequest,
) (*models.Workspace, error) {

	req.Name = strings.TrimSpace(req.Name)
	req.Slug = strings.TrimSpace(strings.ToLower(req.Slug))

	if req.Name == "" {
		return nil, errors.New("workspace name is required")
	}

	if req.Slug == "" {
		return nil, errors.New("workspace slug is required")
	}

	if req.OwnerID == uuid.Nil {
		return nil, errors.New("owner id is required")
	}

	now := time.Now()

	workspace := &models.Workspace{
		ID:          uuid.New(),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		OwnerID:     req.OwnerID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	member := &models.WorkspaceMember{
		ID:          uuid.New(),
		UserID:      req.OwnerID,
		WorkspaceID: workspace.ID,
		JoinedAt:    now,
	}

	role := &models.Role{
		ID:          uuid.New(),
		WorkspaceID: workspace.ID,
		Name:        "owner",
		CreatedAt:   now,
	}

	if err := s.repo.CreateWithOwner(
		ctx,
		workspace,
		member,
		role,
	); err != nil {
		return nil, err
	}

	return workspace, nil
}

func (s *WorkspaceService) ListWorkspaces(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Workspace, error) {

	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}

	return s.repo.ListByUserID(ctx, userID)
}

func (s *WorkspaceService) GetWorkspace(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) (*models.Workspace, error) {

	if id == uuid.Nil {
		return nil, errors.New("workspace id is required")
	}

	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}

	isMember, err := s.memberRepo.IsMember(
		ctx,
		id,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("workspace not found")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *WorkspaceService) GetWorkspaceBySlug(
	ctx context.Context,
	slug string,
) (*models.Workspace, error) {

	slug = strings.TrimSpace(strings.ToLower(slug))

	if slug == "" {
		return nil, errors.New("workspace slug is required")
	}

	return s.repo.GetBySlug(ctx, slug)
}

func (s *WorkspaceService) UpdateWorkspace(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
	req UpdateWorkspaceRequest,
) (*models.Workspace, error) {

	if id == uuid.Nil {
		return nil, errors.New("workspace id is required")
	}

	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}

	workspace, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if workspace.OwnerID != userID {
		return nil, errors.New(
			"only workspace owner can update workspace",
		)
	}

	name := strings.TrimSpace(req.Name)
	slug := strings.TrimSpace(strings.ToLower(req.Slug))

	if name == "" {
		return nil, errors.New("workspace name is required")
	}

	if slug == "" {
		return nil, errors.New("workspace slug is required")
	}

	workspace.Name = name
	workspace.Slug = slug
	workspace.Description = req.Description

	if err := s.repo.Update(ctx, workspace); err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, id)
}

func (s *WorkspaceService) DeleteWorkspace(
	ctx context.Context,
	id uuid.UUID,
	ownerID uuid.UUID,
) error {

	if id == uuid.Nil {
		return errors.New("workspace id is required")
	}

	if ownerID == uuid.Nil {
		return errors.New("owner id is required")
	}

	workspace, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if workspace.OwnerID != ownerID {
		return errors.New(
			"only workspace owner can delete workspace",
		)
	}

	return s.repo.Delete(ctx, id)
}

func (s *WorkspaceService) ListMembers(
	ctx context.Context,
	workspaceID uuid.UUID,
	userID uuid.UUID,
) ([]models.WorkspaceMember, error) {

	if workspaceID == uuid.Nil {
		return nil, errors.New("workspace id is required")
	}

	if userID == uuid.Nil {
		return nil, errors.New("user id is required")
	}

	isMember, err := s.memberRepo.IsMember(
		ctx,
		workspaceID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("workspace not found")
	}

	return s.memberRepo.ListByWorkspaceID(
		ctx,
		workspaceID,
	)
}

func (s *WorkspaceService) GetMember(
	ctx context.Context,
	workspaceID uuid.UUID,
	requesterID uuid.UUID,
	memberUserID uuid.UUID,
) (*models.WorkspaceMember, error) {

	if workspaceID == uuid.Nil ||
		requesterID == uuid.Nil ||
		memberUserID == uuid.Nil {
		return nil, errors.New("invalid id")
	}

	isMember, err := s.memberRepo.IsMember(
		ctx,
		workspaceID,
		requesterID,
	)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, errors.New("workspace not found")
	}

	return s.memberRepo.GetByWorkspaceAndUser(
		ctx,
		workspaceID,
		memberUserID,
	)
}

func (s *WorkspaceService) RemoveMember(
	ctx context.Context,
	workspaceID uuid.UUID,
	requesterID uuid.UUID,
	memberUserID uuid.UUID,
) error {

	if workspaceID == uuid.Nil ||
		requesterID == uuid.Nil ||
		memberUserID == uuid.Nil {
		return errors.New("invalid id")
	}

	workspace, err := s.repo.GetByID(
		ctx,
		workspaceID,
	)
	if err != nil {
		return err
	}

	if workspace.OwnerID != requesterID {
		return errors.New(
			"only workspace owner can remove members",
		)
	}

	if memberUserID == workspace.OwnerID {
		return errors.New(
			"workspace owner cannot be removed",
		)
	}

	return s.memberRepo.Delete(
		ctx,
		workspaceID,
		memberUserID,
	)
}

func (s *WorkspaceService) InviteMember(
	ctx context.Context,
	req CreateInvitationRequest,
) (*models.Invitation, string, error) {

	req.Email = strings.TrimSpace(
		strings.ToLower(req.Email),
	)

	if req.WorkspaceID == uuid.Nil {
		return nil, "", errors.New("workspace id is required")
	}

	if req.InviterID == uuid.Nil {
		return nil, "", errors.New("inviter id is required")
	}

	if req.Email == "" {
		return nil, "", errors.New("email is required")
	}

	workspace, err := s.repo.GetByID(
		ctx,
		req.WorkspaceID,
	)
	if err != nil {
		return nil, "", errors.New("workspace not found")
	}

	if workspace.OwnerID != req.InviterID {
		return nil, "", errors.New(
			"only workspace owner can invite members",
		)
	}

	rawToken, err := invitation.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	tokenHash := invitation.HashToken(rawToken)

	invite := &models.Invitation{
		ID:          uuid.New(),
		WorkspaceID: req.WorkspaceID,
		InvitedBy:   req.InviterID,
		Email:       req.Email,
		ExpiresAt:   time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
	}

	err = s.invitationRepo.Create(
		ctx,
		invite,
		tokenHash,
	)
	if err != nil {
		return nil, "", err
	}

	invite.Token = rawToken

	return invite, rawToken, nil
}
