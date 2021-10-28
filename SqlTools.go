package MyWeb

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getMysqlDB_New(config Config) (db *gorm.DB, err error) {
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
	if !CheckErr(err) {
		return nil, errors.New("sql数据库初始化失败！请检查" + config.Path + "文件是否正确！")
	}
	return
}
