package actions

import (
	"github.com/charmbracelet/log"
	"github.com/kevincobain2000/go-vercel-template/pkg"
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
	response := &Users{
		Query: query,
	}
	log.Debug(response)
	return response
}
