package main

import (
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/sharding"
)

const (
	DSN_LOCAL = "root:root@tcp(localhost:13306)/test"
	//dsn := "mysql://localhost:13306/test?sslmode=disable"
)

type Order struct {
	ID        int64 `gorm:"primarykey"`
	UserID    int64
	ProductID int64
}

func main() {

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: DSN_LOCAL}))
	if err != nil {
		panic(err)
	}

	//createTable(db)

	registerAlgorithm(db)

	// this record will insert to orders_02
	err = db.Create(&Order{UserID: 2}).Error
	if err != nil {
		fmt.Println(err)
	}

	// this record will insert to orders_03
	err = db.Exec("INSERT INTO orders(user_id) VALUES(?)", int64(3)).Error
	if err != nil {
		fmt.Println(err)
	}

	// this will throw ErrMissingShardingKey error
	err = db.Exec("INSERT INTO orders(product_id) VALUES(1)").Error
	fmt.Println(err)

	// this will redirect query to orders_02
	var orders []Order
	err = db.Model(&Order{}).Where("user_id", int64(2)).Find(&orders).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", orders)

	// Raw SQL also supported
	db.Raw("SELECT * FROM orders WHERE user_id = ?", int64(3)).Scan(&orders)
	fmt.Printf("%#v\n", orders)

	// this will throw ErrMissingShardingKey error
	err = db.Model(&Order{}).Where("product_id", "1").Find(&orders).Error
	fmt.Println(err)

	// Update and Delete are similar to create and query
	err = db.Exec("UPDATE orders SET product_id = ? WHERE user_id = ?", 2, int64(3)).Error
	fmt.Println(err) // nil
	err = db.Exec("DELETE FROM orders WHERE product_id = 3").Error
	fmt.Println(err) // ErrMissingShardingKey
}

func registerAlgorithm(db *gorm.DB) {
	middleware := sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		TableShards:         64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "orders")
	db.Use(middleware)
}

func createTable(db *gorm.DB) {
	for i := 0; i < 64; i += 1 {
		table := fmt.Sprintf("orders_%02d", i)
		db.Exec(`DROP TABLE IF EXISTS ` + table)
		db.Exec(`CREATE TABLE ` + table + ` (
			id int PRIMARY KEY,
			user_id int,
			product_id int
		)`)
	}
}
