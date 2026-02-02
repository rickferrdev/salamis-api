package domain

type UserDomain struct {
	ID       uint
	Email    string
	Nickname string
	Username string
	Password string
}

type ProfileDomain struct {
	Status      string
	UserID      uint
	AvatarURL   string
	Description string
}
