package dto

import "time"

type PhotosRow struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Caption   string        `json:"caption"`
	PhotoURL  string        `json:"photo_url"`
	UserID    uint          `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserPhotoRow `json:"User"`
}

type UserPhotoRow struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PayloadPhotos struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type PhotoCreteResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
