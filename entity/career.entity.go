package entity

type Career struct {
	ID 			int64 	`gorm:"primary_key:auto_increment" json:"-"`
	Year		string	`gorm:"type:varchar(100)" json:"-"`
	Company 	string 	`gorm:"type:varchar(100)" json:"-"`
	Position 	string 	`gorm:"type:varchar(100)" json:"-"`
	Color 		string	`gorm:"type:varchar(100)" json:"-"`
	Description string  `gorm:"type:text" json:"-"`
}