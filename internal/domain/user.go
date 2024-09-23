package domain

type User struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email" binding:"required"`
	Password string `db:"password_hash" json:"password" binding:"required"`
}

type SignUp struct {
	Email         string `json:"email" binding:"required" validate:"email"`
	Password      string `json:"password" binding:"required"`
	RetryPassword string `json:"retry_password" binding:"required"`
}

type SignIn struct {
	Email    string `db:"email" json:"email" binding:"required"`
	Password string `db:"password_hash" json:"password" binding:"required"`
}
