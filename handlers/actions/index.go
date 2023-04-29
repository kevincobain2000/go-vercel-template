package actions

import (
	"github.com/k0kubun/pp"
	"github.com/kevincobain2000/go-vercel-template/models"
	log "github.com/sirupsen/logrus"
)

// UserAction
type UserAction struct {
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func (r *UserAction) Get(query string) *models.User {
	user := models.UserModel().First()

	log.Info(pp.Sprint(user))
	return user
}
