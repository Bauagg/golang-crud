package models

type Users struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
}

type CreateRegisterUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserGetLogin struct {
	Id       *uint32 `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
	Token    *string
}

type CretateLoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
