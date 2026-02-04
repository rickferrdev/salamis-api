package domain

type PostDomain struct {
	ID      string
	UserID  string // relation N:1
	Title   string
	Content string
}
