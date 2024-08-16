package types

import (
	"social-network/internal/models"
	"time"
)

type GroupEventMetadata struct {
	Group GroupNotification `json:"group"`
	Event EventDetails      `json:"event"`
}

type GroupNotification struct {
	ID        int              `json:"id"`
	Title     string           `json:"title"`
	InvitedBy models.Requester `json:"invited_by,omitempty"`
}

type GroupRequestNotification struct {
	ID        int              `json:"id"`
	Title     string           `json:"title"`
	Requester models.Requester `json:"requester,omitempty"`
}
type EventDetails struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Creator			string						`json:"creator"`
	Description string           `json:"description,omitempty"`
	EventTime   time.Time        `json:"event_time"`
	Options     []models.Options `json:"options"`
}
