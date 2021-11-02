package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AwesomeServer struct {
	sqlConfig Config
	netConfig Config
}

func (as *AwesomeServer) Start() (err error) {

	db, err := GetMysqlDB_New(as.sqlConfig)
	if !CheckErr(err) {
		return err
	}
	db.Set("gorm:table_options", "AUTO_INCREMENT = 1 ENGINE=InnoDB DEFAULT CHARSET=utf8")

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

	go func() {
		err = router.Run(as.netConfig.Map["ip"] + ":" + as.netConfig.Map["port"])
		if !CheckErr(err) {
			fmt.Println("Sever stopped with error.")
			return
		}
	}()
	var massage string
	var result string
	for {
		if _, err := fmt.Scanf("%s", &massage); !CheckErr(err) {
			result = "Sever stopped."
			fmt.Println(result)
			return nil
		} else {
			switch strings.ToLower(massage) {
			case "stop":
				result = "Sever stopped."
				fmt.Println(result)
				return nil
			case "db.reset":
				if Confirm("reset database") {
					fmt.Println("Starting to reset database...")
					if _, err = SqlDrop(db); !CheckErr(err) {
						fmt.Println("Failed to reset database:", "step 1:drop")
						continue
					}

					if _, err = SqlStart(db); !CheckErr(err) {
						fmt.Println("Failed to reset database:", "step 2:start")
						continue
					}
					fmt.Println("Reset database successfully.")
				}
			case "db.drop":
				if Confirm("drop all tables") {
					if _, err = SqlDrop(db); !CheckErr(err) {
						fmt.Println("Failed to drop tables")
					}
				}
			case "db.start":
				if Confirm("start all tables") {
					if _, err = SqlStart(db); !CheckErr(err) {

					}

				}
			case "db.test-1":
				err = Test01(db)
				CheckErr(err)
			}
		}

	}

}
