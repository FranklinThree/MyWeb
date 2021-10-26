package Resource

import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
	_ "net/http"
	_ "strconv"
)

type AwesomeServer struct {
}

func (as *AwesomeServer) Start() (e error) {
	return e

}
