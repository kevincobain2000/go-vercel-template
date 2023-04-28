package models

type User struct {
	ID      int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	OrgName string `json:"org_name" gorm:"column:org_name;type:varchar(255); NOT NULL"`
}
