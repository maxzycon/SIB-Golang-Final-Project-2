package dto

import "time"

type SocialMediaWrappper struct {
	SocialMedias []*SocialMediaRow `json:"social_medias"`
}

type SocialMediaRow struct {
	ID             uint                `json:"id"`
	Name           string              `json:"name"`
	SocialMediaURL string              `json:"social_media_url"`
	UserID         uint                `json:"UserId"`
	UpdatedAt      time.Time           `json:"updated_at"`
	CreatedAt      time.Time           `json:"created_at"`
	User           *UserSocialMediaRow `json:"User"`
}

type UserSocialMediaRow struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PayloadSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_medias_url"`
}

type PayloadUpdateSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_medias_url"`
}

type SocialMediaCreteResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
