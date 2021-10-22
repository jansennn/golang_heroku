package entity

type Project struct {
	ID     		int64 	`gorm:"primary_key:auto_increment" json:"-"`
	Name   		string 	`gorm:"type:varchar(100)" json:"-"`
	Image  		string 	`gorm:"type:varchar(100)" json:"-"`
	Description string 	`gorm:"type:text" json:"-"`
	LinkCode    string 	`gorm:"type:varchar(100)" json:"-"`
}