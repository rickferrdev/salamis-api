package auth

type RequestUserRegisterDTO struct {
	Nickname string `json:"nickname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestUserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseUserLoginDTO struct {
	Nickname string `json:"nickname" binding:"required"`
	Token    string `json:"token"`
}

type ResponseUserRegisterDTO struct {
	Nickname string `json:"nickname" binding:"required"`
}
