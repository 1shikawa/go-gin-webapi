package utility

// import (
// 	"fmt"
// 	// "os"
// 	"time"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Article struct {
// 	gorm.Model
// 	Title string
// 	Body  string
// }

// type User struct {
// 	gorm.Model
// 	Name string
// 	Email  string
// }

// var Db *gorm.DB

// func init() {
// 	user := "root"
// 	pw := "root_password"
// 	db_name := "test_database"
// 	fmt.Println(user, pw, db_name)
// 	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
// 	dialector := mysql.Open(path)
// 	log.Println("Start Conneting Database...")
// 	var err error
// 	if Db, err = gorm.Open(dialector); err != nil {
// 		connect(dialector, 100)
// 	}
// 	log.Println("Database connected!!")

// 	// Migrate the schema
// 	log.Println("Start Migrating schema")
//   Db.AutoMigrate(&Article{})
//   Db.AutoMigrate(&User{})
// 	log.Println("Migrate completed")

// 	// Create
//   // Db.Create(&Article{Title: "gorm_test", Body: "gorm_test_body"})
// 	// Db.Create(&User{Name: "ishikawa", Email: "ishikawa.toru@gmail.com"})

// }

// func connect(dialector gorm.Dialector, count uint) {
// 	var err error
// 	if Db, err = gorm.Open(dialector); err != nil {
// 		if count > 1 {
// 			time.Sleep(time.Second * 2)
// 			count--
// 			fmt.Printf("retry... count:%v\n", count)
// 			connect(dialector, count)
// 			return
// 		}
// 		panic(err.Error())
// 	}
// }
