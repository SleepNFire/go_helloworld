package data

import (
	"database/sql"
	"errors"
	"fmt"
	"go_bootstrap/project/config"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlAccessor struct {
	Db     *sql.DB
	Config config.MySql
}

func NewMySqlAccessor(globalConf *config.Config) (*MySqlAccessor, error) {
	MySqlAccessor := &MySqlAccessor{
		Config: globalConf.MySql,
	}
	Db, err := MySqlAccessor.connectToMySql()
	if err != nil {
		fmt.Println("error during connecting to mysql")
		return nil, errors.New("error during connecting to mysql")
	}

	MySqlAccessor.Db = Db
	return MySqlAccessor, nil
}

func (mysql *MySqlAccessor) connectToMySql() (Db *sql.DB, err error) {
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", mysql.Config.User, mysql.Config.Password, mysql.Config.Adress, mysql.Config.Database))
	if err != nil {
		return nil, err
	}

	return Db, nil
}

func (mysql *MySqlAccessor) RegisterEndpoints(router *gin.Engine) {
	router.GET("/mysql/health", mysql.Ping)
}

func (mysql *MySqlAccessor) Ping(c *gin.Context) {
	err := mysql.Db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal error: connection mysql")
	}
	c.JSON(http.StatusOK, "mysql is good")
}
