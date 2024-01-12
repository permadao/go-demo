package schema

import "time"

type TestTable struct {
	ID        int64      `gorm:"primary_key;auto_increment" json:"id"`
	UpdatedAt *time.Time `gorm:"ASSOCIATION_AUTOUPDATE" json:"-"`
	CreatedAt *time.Time `gorm:"ASSOCIATION_AUTOCREATE" json:"-"`
	TestField string
}
