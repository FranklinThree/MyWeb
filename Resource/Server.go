package Resource

import (
	_ "gorm.io/gorm"

	_ "database/sql"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http"
	_ "strconv"
)

type AwesomeServer struct {
}

func (as *AwesomeServer) Start() (e error) {
	return e

}
