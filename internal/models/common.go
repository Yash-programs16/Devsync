package models

type UserStatus string

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusAway    UserStatus = "away"
	UserStatusBusy    UserStatus = "busy"
	UserStatusOffline UserStatus = "offline"
)

type ChannelType string

const (
	ChannelText     ChannelType = "text"
	ChannelCode     ChannelType = "code"
	ChannelDecision ChannelType = "decision"
)

type WorkspaceRole string

const (
	RoleOwner  WorkspaceRole = "owner"
	RoleAdmin  WorkspaceRole = "admin"
	RoleMember WorkspaceRole = "member"
)
