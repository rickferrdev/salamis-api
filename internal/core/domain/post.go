package domain

type PostDomain struct {
	ID       uint
	AuthorID uint // relation N:1
	Title    string
	Content  string
}
