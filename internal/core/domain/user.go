package domain

type UserDomain struct {
	ID       uint
	Email    string
	Nickname string
	Username string
	Password string
	Profile  ProfileDomain // relation 1:1
	Posts    []PostDomain  // relation 1:N
}

type ProfileDomain struct {
	Status      string
	AvatarURL   string // relation 1:1
	Description string
}
