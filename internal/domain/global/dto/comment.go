package dto

import "time"

type CommentRow struct {
	ID        uint             `json:"id"`
	Message   string           `json:"message"`
	PhotoID   uint             `json:"photo_id"`
	UserID    uint             `json:"user_id"`
	UpdatedAt time.Time        `json:"updated_at"`
	CreatedAt time.Time        `json:"created_at"`
	User      *UserCommentRow  `json:"User"`
	Photo     *PhotoCommentRow `json:"Photo"`
}

type PhotoCommentRow struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type UserCommentRow struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PayloadComment struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
}

type PayloadUpdateComment struct {
	Message string `json:"message"`
}

type CommentCreteResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UpdatedAt time.Time `json:"updated_at"`
}
