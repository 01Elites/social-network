package types

import "social-network/internal/models"

type GroupRequestMetadata struct {
	UserDetails UserDetails       `json:"requester"`
	Group       GroupNotification `json:"group"`
}

type GroupEventMetadata struct {
	Group GroupNotification `json:"group"`
	Event EventDetails      `json:"event"`
}

type GroupNotification struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	InvitedBy models.Requester `json:"invited_by,omitempty"`
}

type EventDetails struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Options     []string `json:"options"`
	Description string   `json:"description,omitempty"`
}
