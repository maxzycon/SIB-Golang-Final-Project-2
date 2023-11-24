package dto

import "time"

type UserRow struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Role     uint    `json:"role"`
	Image    *string `json:"image"`
}

type UserRowDetail struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserRowDetailUpdateRes struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PayloadUpdateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PayloadLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PayloadUpdateProfile struct {
	Password string `json:"password"`
}

type PayloadCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type LoginRes struct {
	AccessToken string `json:"token"`
}

type UserPaginatedRow struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Merchant string `json:"merchant"`
	Role     uint   `json:"role"`
}
