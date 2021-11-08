package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AwesomeServer struct {
	sqlConfig         Config
	netConfig         Config
	wsConfig          Config
	configs           map[string]Config
	database          *gorm.DB
	httpServer        *gin.Engine
	websocketUpgrader *websocket.Upgrader
	websocketServer   net.Listener
	isReady           int
	maxRetryTimes     int
	retryTime         int64
}

func (as *AwesomeServer) New() (err error) {
	as.maxRetryTimes = 3
	as.retryTime = 3
	as.isReady = 10000
	go func() {
		err = as.StartSql()
		if !CheckErr(err) {
			ConsolePrint(Error, "数据库初始化失败")
			as.isReady -= 100
			return
		}
		ConsolePrint(Info, "数据库初始化成功")
		as.isReady += 1
	}()
	go func() {
		err = as.StartHttpServer()
		if !CheckErr(err) {
			ConsolePrint(Error, "http服务器初始化失败")
			as.isReady -= 200
			return
		}
		ConsolePrint(Info, "http服务器初始化成功")
		as.isReady += 2
	}()
	go func() {
		err = as.StartWebsocketServer()
		if !CheckErr(err) {
			ConsolePrint(Error, "websocket服务器初始化失败:httpserver")
			as.isReady -= 500
			return
		}
		err = as.StartWebsocketUpdater()
		if !CheckErr(err) {
			ConsolePrint(Error, "websocket服务器初始化失败:upgrader")
			as.isReady -= 500
			return
		}
		ConsolePrint(Info, "websocket服务器初始化成功")
		as.isReady += 5
	}()
	return err
}
func (as *AwesomeServer) Start() (err error) {
	ConsolePrint(Info, "程序初始化中...")
	if as.isReady != 5233 {
		if as.isReady == 0 {
			ConsolePrint(Warning, "总服务器未初始化！")
			ConsolePrint(Info, "正在尝试初始化服务器...")
			err = as.New()
			CheckErr(err)
		}
		waitTime := 0
		for as.isReady != 10008 {
			time.Sleep(time.Second * time.Duration(as.retryTime))
			waitTime++
			if waitTime >= as.maxRetryTimes {
				ConsolePrint(Error, "服务器初始化多次失败，不再重试。", "value", as.isReady, "times", waitTime)
				break
			}
			if as.isReady < 0 {
				return errors.New("错误:初始化失败")
			}
			ConsolePrint(Warning, "服务器初始化失败，正在重试... ", "value", as.isReady, "times", waitTime)
		}
		ConsolePrint(Info, "服务器初始化完成")
		as.isReady = 5233
	}
	ConsolePrint(Info, "程序初始化完成")

	//启动http服务器
	go func() {
		err = as.httpServer.Run(as.netConfig.Map["ip"] + ":" + as.netConfig.Map["port"])
		if !CheckErr(err) {
			fmt.Println("Sever stopped with error.")
			return
		}

	}()
	//tcp/websocket服务器

	CheckErr(err)
	{

	}

	go func() {

	}()
	var massage string
	var result string
	for {
		if _, err = fmt.Scanf("%s", &massage); !CheckErr(err) {
			result = "Sever stopped."
			fmt.Println(result)
			return nil
		} else {
			switch strings.ToLower(massage) {
			case "stop":
				ConsolePrint(Info, "Sever stopped.")
				return nil
			case "db.reset":
				if Confirm("reset database") {
					ConsolePrint(Info, "Starting to reset database...")
					if _, err = QuestionSqlDrop(as.database); !CheckErr(err) {
						ConsolePrint(Error, "Failed to reset database:", "step 1:drop")
						continue
					}

					if _, err = QuestionSqlStart(as.database); !CheckErr(err) {
						ConsolePrint(Error, "Failed to reset database:", "step 2:start")
						continue
					}
					ConsolePrint(Info, "Reset database successfully.")
				}
			case "db.drop":
				if Confirm("drop all tables") {
					if _, err = QuestionSqlDrop(as.database); !CheckErr(err) {
						ConsolePrint(Error, "Failed to drop tables:", err)
					}
				}
			case "db.start":
				if Confirm("start all tables") {
					if _, err = QuestionSqlStart(as.database); !CheckErr(err) {

					}

				}
			case "db.test-1":
				err = Test01(as.database)
				CheckErr(err)
			}
		}

	}

}
func (as *AwesomeServer) StartSql() (err error) {
	as.database, err = GetMysqlDBNew(as.sqlConfig)
	if !CheckErr(err) {
		return err
	}
	as.database.Set("gorm:table_options", "AUTO_INCREMENT = 1 ENGINE=InnoDB DEFAULT CHARSET=utf8")
	return
}
func (as *AwesomeServer) StartHttpServer() (err error) {
	as.httpServer = gin.New()
	//http服务器
	{
		as.httpServer.LoadHTMLGlob("static/html/*")
		as.httpServer.StaticFS("/static", http.Dir("./static"))
		as.httpServer.GET("/index.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Title": "666",
			})
		})
		as.httpServer.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Title": "666",
			})
		})
		//网页返回数据到服务器
		as.httpServer.POST("/questionnaire/input", func(c *gin.Context) {

		})

		//发送网页以及题目数据
		as.httpServer.GET("/questionnaire/output", func(c *gin.Context) {
			//qnnID := c.Query("qnnID")
			//db.
		})

	}
	return err
}

//以下为ws模块

// ConnHandle 接收并处理收到的ws请求
func (as *AwesomeServer) ConnHandle(conn *websocket.Conn) (err error) {
	defer func() {
		err = conn.Close()
		CheckErr(err)
	}()
	stopCh := make(chan int)
	go func() {
		err = as.Send(conn, stopCh)
		CheckErr(err)
	}()
	for {
		err = conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		if !CheckErr(err) {
			return errors.New("ws:设置超时日期时发生错误")
		}
		var msg []byte
		_, msg, err = conn.ReadMessage()
		if err != nil {
			close(stopCh)
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					fmt.Printf("ReadMessage timeout remote %v\n", conn.RemoteAddr())
					return
				}
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v\n", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Println("收到消息:", string(msg))
	}

}

// Send10w 发送10万条信息
func (as *AwesomeServer) Send10w(conn *websocket.Conn) (err error) {
	for i := 0; i < 100000; i++ {
		data := fmt.Sprintf("hello websocket tast form server %v", time.Now().UnixNano())
		err = conn.WriteMessage(1, []byte(data))
		if err != nil {
			fmt.Println("Send message failed", err)
			return
		}

	}
	return
}

// Send 发送信息
func (as *AwesomeServer) Send(conn *websocket.Conn, stopCh chan int) (err error) {
	err = as.Send10w(conn)
	CheckErr(err)
	for {
		select {
		case <-stopCh:
			fmt.Println("connection closed")
			return
		case <-time.After(time.Second * 1):
			data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
			err = conn.WriteMessage(1, []byte(data))
			fmt.Println("sending...")
			if err != nil {
				fmt.Println("Send message failed", err)
				return
			}
		}
	}
}

// StartWebsocketServer 开启ws的服务器
func (as *AwesomeServer) StartWebsocketServer() (err error) {
	as.websocketServer, err = net.Listen("tcp", as.wsConfig.Map["ip"]+":"+as.wsConfig.Map["port"])
	if !CheckErr(err) {
		return errors.New("ws:tcp服务器初始化失败")
	}
	err = http.Serve(as.websocketServer, as)
	if !CheckErr(err) {
		return errors.New("ws:http服务器初始化失败")
	}
	return nil
}

//实现handler接口
func (as *AwesomeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != as.wsConfig.Map["urlPath"] {
		httpCode := http.StatusInternalServerError
		reactPhrase := http.StatusText(httpCode)
		fmt.Println("path error ", reactPhrase)
		http.Error(w, reactPhrase, httpCode)
		return
	}
	// 收到 http 请求后 升级 协议
	conn, err := as.websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())
	go func() {
		err = as.ConnHandle(conn)
		CheckErr(err)
	}()

}

// StartWebsocketUpdater 初始化ws协议升级器
func (as *AwesomeServer) StartWebsocketUpdater() (err error) {
	rbs, err := strconv.Atoi(as.wsConfig.Map["readBufferSize"])
	wbs, err := strconv.Atoi(as.wsConfig.Map["writeBufferSize"])
	as.websocketUpgrader = &websocket.Upgrader{
		ReadBufferSize:  rbs,
		WriteBufferSize: wbs,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != as.wsConfig.Map["urlPath"] {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return err
}

func (as *AwesomeServer) HttpWork(relativePath string, a gin.HandlerFunc) {

}
