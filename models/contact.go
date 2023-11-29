package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerId   uint
	ContactId uint
	Type      int // type 0 is friend, type 1 is group
	Desc      string
}

func (table *Contact) TableName() string {
	return "contact"
}
