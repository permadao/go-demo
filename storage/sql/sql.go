package sql

import (
	"github.com/permadao/go-demo/storage/sql/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQL struct {
	db *gorm.DB
}

func New(dsn string) *SQL {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return &SQL{db}
}

func (s *SQL) Migrate() {
	s.db.AutoMigrate(&schema.TestTable{})
}

func (s *SQL) CreateTestTable(t *schema.TestTable, tx *gorm.DB) error {
	if tx == nil {
		tx = s.db
	}

	return tx.Create(&t).Error
}

func (s *SQL) FindTestTables(tx *gorm.DB) (ts []*schema.TestTable, err error) {
	if tx == nil {
		tx = s.db
	}

	err = tx.Find(&tx).Error
	return
}

func (s *SQL) UpdateTestTable(t *schema.TestTable, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = s.db
	}

	return tx.First(t).Update("test_tables", t.TestField).Error
}
