package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	ROle     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
