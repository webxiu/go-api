package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UserList struct {
	gorm.Model
	// ID       uint   `gorm:"primary_key" json:"id"`
	UserName string `gorm:"type:varchar(20); not null" json:"username" binding:"required"`
	Status   int    `gorm:"type:int(1); not null" json:"status" binding:"required"`
	Phone    string `gorm:"type:char(11); not null" json:"phone" binding:"required"`
	Email    string `gorm:"type:varchar(30); not null" json:"email" binding:"required"`
	Address  string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	// CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	// UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
	// DeletedAt time.Time `gorm:"column:deleted_at;type:datetime"`
}

func Connect() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 创建表没有s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	fmt.Println("==1===", db)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return db
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库初始化失败", err)
		return db
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 自动创建数据库 https://gorm.io/zh_CN/docs/migration.html#AutoMigrate
	db.AutoMigrate(&UserList{})

	fmt.Println("数据库创建成功")

	return db
}
