package actions

import (
	"github.com/kevincobain2000/go-vercel-template/models"
)

// UserAction
type UserAction struct {
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func (r *UserAction) Get(query string) *models.User {
	user := models.UserModel().First()
	return user
}
