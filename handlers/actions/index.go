package actions

import (
	"github.com/k0kubun/pp"
	"github.com/kevincobain2000/go-vercel-template/pkg"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// UserAction
type UserAction struct {
	db *gorm.DB
}

func NewUserAction() *UserAction {
	return &UserAction{
		db: pkg.DB(),
	}
}

// Users
type Users struct {
	Query string `json:"query"`
}

func (r *UserAction) Get(query string) *Users {
	pkg.DB()
	response := &Users{
		Query: query,
	}
	log.Info(pp.Sprint(response))
	return response
}
