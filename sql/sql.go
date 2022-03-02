package sql

import (
	"com/github/FranklinThree/MyWeb/questionnaire"
	"com/github/FranklinThree/MyWeb/universal"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetMysqlDBNew 初始化mysql8.0
func GetMysqlDBNew(config universal.Config) (db *gorm.DB, err error) {
	fmt.Println(config.Map["userName"] + ":" + config.Map["userKey"] +
		"@(" + config.Map["ip"] + ":" + config.Map["port"] + ")/" +
		config.Map["databaseName"] +
		"?charset=" + config.Map["charset"])
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.Map["userName"] + ":" + config.Map["userKey"] +
			"@(" + config.Map["ip"] + ":" + config.Map["port"] + ")/" +
			config.Map["databaseName"] +
			"?charset=" + config.Map["charset"],
	}),
		&gorm.Config{},
	)
	if !universal.CheckErr(err) {
		universal.ConsolePrint(universal.Error, "sql数据库初始化失败！请检查", "\""+config.Path+"\"", "文件是否正确！")
		return nil, err
	}
	return
}

// GetMysqlDBOld 初始化mysql5.6
func GetMysqlDBOld(config universal.Config) (db *gorm.DB, err error) {
	fmt.Println(config.Map["userName"] + ":" + config.Map["userKey"] +
		"@(" + config.Map["ip"] + ":" + config.Map["port"] + ")/" +
		config.Map["databaseName"] +
		"?charset=" + config.Map["charset"])
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.Map["userName"] + ":" + config.Map["userKey"] +
			"@(" + config.Map["ip"] + ":" + config.Map["port"] + ")/" +
			config.Map["databaseName"] +
			"?charset=" + config.Map["charset"],
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}),
		&gorm.Config{},
	)
	if !universal.CheckErr(err) {
		universal.ConsolePrint(universal.Error, "sql数据库初始化失败！请检查", "\""+config.Path+"\"", "文件是否正确！")
		return nil, err
	}
	return
}
func QuestionSqlStart(db *gorm.DB) (value int, err error) {
	universal.ConsolePrint(universal.Info, "Starting to start all Question tables...")
	value = 0
	if !db.Migrator().HasTable(questionnaire.Questionnaire{}) {
		err = db.Migrator().CreateTable(questionnaire.Questionnaire{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 100
	} else {
		universal.ConsolePrint(universal.Warning, "Table Questionnaires already exists")
	}

	if !db.Migrator().HasTable(questionnaire.Question{}) {
		err = db.Migrator().CreateTable(questionnaire.Question{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 10
	} else {
		universal.ConsolePrint(universal.Warning, "[WARNING]Table Questions already exists")
	}
	if !db.Migrator().HasTable(questionnaire.Choice{}) {
		err = db.Migrator().CreateTable(questionnaire.Choice{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 1
	} else {
		universal.ConsolePrint(universal.Warning, "Table Choices already exists")
	}

	if value == 0 {
		universal.ConsolePrint(universal.Warning, "Nothing to do.")
	} else {
		universal.ConsolePrint(universal.Info, "All Question tables were started.")
	}
	return value, nil
}
func QuestionSqlDrop(db *gorm.DB) (value int, err error) {
	universal.ConsolePrint(universal.Info, "Starting to drop all Question tables...")
	if db.Migrator().HasTable(questionnaire.Choice{}) {
		err = db.Migrator().DropTable(questionnaire.Choice{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 1
	} else {
		universal.ConsolePrint(universal.Warning, "Table Choices was already removed")
	}
	if db.Migrator().HasTable(questionnaire.Question{}) {
		err = db.Migrator().DropTable(questionnaire.Question{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 10
	} else {
		universal.ConsolePrint(universal.Warning, "Table Questions was already removed")
	}
	if db.Migrator().HasTable(questionnaire.Questionnaire{}) {
		err = db.Migrator().DropTable(questionnaire.Questionnaire{})
		if !universal.CheckErr(err) {
			return -1, err
		}
		value += 100
	} else {
		universal.ConsolePrint(universal.Warning, "Table Questionnaires was already removed")
	}
	if value == 0 {
		universal.ConsolePrint(universal.Warning, "Nothing to do.")
	} else {
		universal.ConsolePrint(universal.Info, "All Question tables were dropped")
	}
	return value, nil
}

//func StatisticSqlStart(db *gorm.DB)(value int,err error){
//	ConsolePrint(Info,"Starting to start all Statistic tables...")
//}

func Test01(db *gorm.DB) (err error) {
	universal.ConsolePrint(universal.Info, "starting executing test-1...")
	objs := make([]questionnaire.Question, 2)
	objs[0] = questionnaire.Question{0, 110201, "您谈过恋爱嘛?", "",
		[]questionnaire.Choice{
			{0, 0, "是", 0},
			{0, 0, "否", 0},
		},
		0,
	}
	objs[1] = questionnaire.Question{
		0, 110201, "您的性别是?", "",
		[]questionnaire.Choice{
			{0, 0, "♂", 0},
			{0, 0, "♀", 0},
		},
		0,
	}
	qnn := questionnaire.Questionnaire{0, 0, "测试问卷1", "", objs}
	fmt.Println(&qnn)
	db.Create(&qnn)
	universal.ConsolePrint(universal.Info, "execute test-1 successfully")
	return nil
}
