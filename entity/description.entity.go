package entity

type Description struct {
	ID            int64  `gorm:"primary_key:auto_increment" json:"-"`
	Description   string `gorm:"type:text" json:"-"`

}