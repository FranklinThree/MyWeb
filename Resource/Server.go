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
	SqlConfig Config
	NetConfig Config
}

func (as *AwesomeServer) Start() (err error) {
	db, err := getMysqlDB_New(as.SqlConfig)
	objs := make([]QuestionnaireObject, 16)
	objs[1] = &Question{100, "您谈过恋爱嘛?", []Choice{{1, "是"}, {2, "否"}}}
	qnn := Questionnaire{1, objs}
	db.Create(&qnn)
	return

}
