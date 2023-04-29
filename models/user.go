package models

import "github.com/kevincobain2000/go-vercel-template/pkg"

type User struct {
	ID      int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	OrgName string `json:"org_name" gorm:"column:org_name;type:varchar(255); NOT NULL"`
}

func UserModel() *User {
	return &User{}
}

// get first user
func (u *User) First() *User {
	var user User
	pkg.DB().First(&user)
	return &user
}
