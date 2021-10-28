package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AwesomeServer struct {
	sqlConfig Config
	netConfig Config
}

func (as *AwesomeServer) Start() (err error) {
	db, err := getMysqlDB_New(as.sqlConfig)
	if !CheckErr(err) {
		return err
	}
	objs := make([]QuestionnaireObject, 16)
	objs[1] = &Question{100, "您谈过恋爱嘛?", "", []Choice{{1, "是"}, {2, "否"}}}
	qnn := Questionnaire{1, "测试问卷1", "", objs}

	//测试模块
	//db.Exec("create table text(id int not null)")

	fmt.Println("this is", qnn)
	fmt.Println(db.Migrator().HasTable(&Questionnaire{}))
	db.Set("gorm:table_options", "AUTO_INCREMENT = 1 ENGINE=InnoDB DEFAULT CHARSET=utf8")
	err = db.Migrator().CreateTable(&Questionnaire{})
	CheckErr(err)
	fmt.Println(&qnn)
	db.Create(&qnn)
	var got = new(Questionnaire)
	db.First(got)
	fmt.Printf("%#v\n", got)
	//
	router := gin.New()

	router.LoadHTMLGlob("static/html/*")
	router.StaticFS("/static", http.Dir("./static"))

	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "666",
		})
	})
	//网页返回数据到服务器
	router.POST("/questionnaire/input", func(c *gin.Context) {

	})

	//发送网页以及题目数据
	router.GET("/questionnaire/output", func(c *gin.Context) {
		//qnnID := c.Query("qnnID")
		//db.
	})

	err = router.Run(as.netConfig.Map["ip"] + ":" + as.netConfig.Map["port"])
	if !CheckErr(err) {
		return errors.New("服务器初始化失败")
	}
	return

}
