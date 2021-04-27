package database

import (
	"fmt"
	"ormt/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = connect()
}

func GetDB() *gorm.DB {
	return db
}

func connect() *gorm.DB {

	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User+":"+config.Pass,
		config.Host+":"+config.Port,
		config.Name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		PrepareStmt: true, //缓存预编译语句
	})
	if err != nil {
		fmt.Println("连接数据库失败: " + err.Error())
	} else {
		fmt.Println("连接数据库" + config.Name + "成功")
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(20)
	// SetConnMaxLifetime 设置了连接可复用的最大时间 分钟。
	sqlDB.SetConnMaxLifetime(5)
	return db
}
