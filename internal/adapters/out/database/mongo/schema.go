package mongo

import "github.com/rickferrdev/salamis-api/internal/core/domain"

type UserSchema struct {
	ID       string        `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Nickname string        `bson:"nickname"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Profile  ProfileSchema `bson:"profile"`
	Posts    []PostSchema  `bson:"posts"`
}

type ProfileSchema struct {
	Status      string `bson:"status"`
	AvatarURL   string `bson:"avatar_url"`
	Description string `bson:"description"`
}

type PostSchema struct {
	ID      string `bson:"_id,omitempty"`
	UserID  string `bson:"user_id"`
	Title   string `bson:"title"`
	Content string `bson:"content"`
}

// Profile Mapper
func (u *ProfileSchema) ProfileSchemaToDomain() *domain.ProfileDomain {
	return &domain.ProfileDomain{
		Status:      u.Status,
		AvatarURL:   u.AvatarURL,
		Description: u.Description,
	}
}
func ProfileDomainToSchema(profile domain.ProfileDomain) *ProfileSchema {
	return &ProfileSchema{
		Status:      profile.Status,
		AvatarURL:   profile.AvatarURL,
		Description: profile.Description,
	}
}

// Post Mapper
func (u *PostSchema) PostSchemaToDomain() *domain.PostDomain {
	return &domain.PostDomain{
		ID:      u.ID,
		UserID:  u.UserID,
		Title:   u.Title,
		Content: u.Content,
	}
}
func PostDomainToSchema(post domain.PostDomain) *PostSchema {
	return &PostSchema{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}
func PostsMapperToSchema(posts []domain.PostDomain) []PostSchema {
	schemaPosts := make([]PostSchema, len(posts))
	for i, p := range posts {
		schemaPosts[i] = *PostDomainToSchema(p)
	}
	return schemaPosts
}
func PostsMapperToDomain(posts []PostSchema) []domain.PostDomain {
	domainPosts := make([]domain.PostDomain, len(posts))
	for i, p := range posts {
		domainPosts[i] = *p.PostSchemaToDomain()
	}
	return domainPosts
}

// User Mapper
func UserDomainToSchema(user domain.UserDomain) *UserSchema {
	return &UserSchema{
		ID:       user.ID,
		Nickname: user.Nickname,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Profile:  *ProfileDomainToSchema(user.Profile),
		Posts:    PostsMapperToSchema(user.Posts),
	}
}
func (u *UserSchema) UserSchemaToDomain() *domain.UserDomain {
	return &domain.UserDomain{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Profile:  *u.Profile.ProfileSchemaToDomain(),
		Posts:    PostsMapperToDomain(u.Posts),
	}
}
