package models

type GroupFeed struct {
	ID       int    `json:"id,omitempty"`
	Members  []User `json:"members,omitempty"`
	Posts    []Post `json:"posts,omitempty"`
	IsMember bool   `json:"ismember,omitempty"`
}


type Create_Group struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type GroupInvite struct {
	ID			    int    `json:"invite_id,omitempty"`
	GroupID 	  int    `json:"group_id,omitempty"`
	ReceiverID  string `json:"receiver_id,omitempty"`
}

type InviteResponse struct {
	ID			    int    `json:"invite_id,omitempty"`
	GroupID 		int		 `json:"group_id,omitempty"`
	Status      string `json:"status,omitempty"`
}