package ports

// Data transfer objects (DTOs)
// User related DTOs
type UserOutput struct {
	Email    string
	Username string
	Nickname string
}

// Profile related DTOs
type ProfileInput struct {
	Status      string
	AvatarURL   string
	Description string
	UserID      uint
}

type ProfileOutput struct {
	Status      string
	AvatarURL   string
	Description string
	User        UserOutput
}

// Post related DTOs
type PostInput struct {
	Title    string
	Content  string
	AuthorID uint
}

type PostOutput struct {
	ID      uint
	Title   string
	Content string
	Author  UserOutput
}

// Auth related DTOs
type AuthInput struct {
	Email    string
	Password string
	Username string
	Nickname string
}

type AuthOutput struct {
	Token string
	User  UserOutput
}
