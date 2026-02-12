package auth

type RequestLoginDTO struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type ResponseLoginDTO struct {
	Token string `binding:"required"`
}
