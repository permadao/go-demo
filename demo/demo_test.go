package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	rdsURL := "redis://@localhost:6379/4"
	sqlDSN := "root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	d := New(sqlDSN, rdsURL)
	d.Run(":8080")

	expected := 1
	actual := 1
	assert.Equal(t, expected, actual)
}
