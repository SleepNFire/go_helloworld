package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessorMySQl(t *testing.T) {
	if os.Getenv("FUNCTIONNEL_TEST") != "1" {
		return
	}
	t.Run("connection", Test_Connection)
}

func Test_Connection(t *testing.T) {
	err := Mysql.Db.Ping()
	assert.NoError(t, err)
}
