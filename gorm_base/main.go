package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dataSourceName := "root:123456@tcp(localhost:3306)/mybatis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//自动创建表 schema
	// CREATE TABLE `products` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`code` longtext,`price` bigint unsigned,PRIMARY KEY (`id`),INDEX idx_products_deleted_at (`deleted_at`))
	db.AutoMigrate(&Product{})
	// INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2021-08-08 22:19:11.125','2021-08-08 22:19:11.125',NULL,'d42',100)
	//db.Create(&Product{
	//	Code:  "d42",
	//	Price: 100,
	//})
	var product Product
	//SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	db.First(&product, 1)

	//db.First(&product, "code=?", "d42")

	//更新字段 UPDATE `products` SET `price`=200,`updated_at`='2021-08-08 22:25:50.243' WHERE `id` = 1
	db.Model(&product).Update("Price", 200)

	//更新更新多个字段 UPDATE `products` SET `updated_at`='2021-08-08 22:25:50.387',`code`='F42',`price`=200 WHERE `id` = 1
	db.Model(&product).Updates(Product{
		Code:  "F42",
		Price: 200,
	})

	//UPDATE `products` SET `code`='F42',`price`=200,`updated_at`='2021-08-08 22:25:50.432' WHERE `id` = 1
	db.Model(&product).Updates(map[string]interface{}{
		"Price": 200,
		"Code":  "F42",
	})

	db.Delete(&product, 1)
}
