package schema

import "gorm.io/gorm"

type TestTable struct {
	gorm.Model
	TestField string
}
