package MyWeb

import (
	_ "database/sql"
	"errors"
	"fmt"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
	"net/http"
	_ "net/http"
	_ "strconv"
)

type AwesomeServer struct {
	SqlConfig Config
	NetConfig Config
}

func (as *AwesomeServer) Start() (err error) {
	db, err := getMysqlDB_New(as.SqlConfig)
	if !CheckErr(err) {
		return err
	}
	objs := make([]QuestionnaireObject, 16)
	objs[1] = &Question{100, "您谈过恋爱嘛?", "", []Choice{{1, "是"}, {2, "否"}}}
	qnn := Questionnaire{1, "测试问卷1", "", objs}

	//测试模块
	fmt.Println(&qnn)
	db.Create(&qnn)
	var got = new(Questionnaire)
	db.First(got)
	fmt.Printf("%#v\n", got)
	//
	router := gin.New()

	router.LoadHTMLGlob("./static/html/*")
	router.StaticFS("/static", http.Dir("./static"))

	//网页返回数据到服务器
	router.POST("/questionnaire/input", func(c *gin.Context) {

	})

	//发送网页以及题目数据
	router.GET("/questionnaire/output", func(c *gin.Context) {
		//qnnID := c.Query("qnnID")
		//db.
	})

	err = router.Run(as.NetConfig.Map["ip"] + ":" + as.NetConfig.Map["port"])
	if !CheckErr(err) {
		return errors.New("服务器初始化失败")
	}
	return

}
