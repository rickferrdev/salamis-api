package mongo

import (
	"github.com/rickferrdev/salamis-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Nickname string             `bson:"nickname"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
	Profile  ProfileSchema      `bson:"profile"`
}

func (UserSchema) TableName() string {
	return "users"
}

func (u *UserSchema) FromDomain() *domain.UserDomain {
	return &domain.UserDomain{
		ID:       u.ID.Hex(),
		Nickname: u.Nickname,
		Password: u.Password,
		Username: u.Username,
		Email:    u.Email,
		Profile: domain.ProfileDomain{
			Status:      u.Profile.Status,
			AvatarURL:   u.Profile.AvatarURL,
			Description: u.Profile.Description,
		},
	}
}

type ProfileSchema struct {
	Status      string `bson:"status"`
	AvatarURL   string `bson:"avatar_url"`
	Description string `bson:"description"`
}

func UserFromSchema(user *domain.UserDomain) *UserSchema {
	return &UserSchema{
		Username: user.Username,
		Nickname: user.Nickname,
		Password: user.Password,
		Email:    user.Email,
		Profile: ProfileSchema{
			Status:      user.Profile.Status,
			AvatarURL:   user.Profile.AvatarURL,
			Description: user.Profile.Description,
		},
	}
}
