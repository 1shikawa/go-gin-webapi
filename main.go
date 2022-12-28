package main

import (
	"embed"
	"html/template"
	"log"
	// "os"
	// "fmt"
	// "time"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"1shikawa.com/m/router"
)

// var DB *gorm.DB

// func Init() {
// 	// user := os.Getenv("MYSQL_USER")
// 	// pw := os.Getenv("MYSQL_PASSWORD")
// 	// db_name := os.Getenv("MYSQL_DATABASE")
// 	user := "root"
// 	pw := "root_password"
// 	db_name := "test_database"
// 	var dsn string = fmt.Sprintf("%s:%s@tcp(mysql)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, db_name)
// 	dialector := mysql.Open(dsn)

// 	var err error
// 	if DB, err = gorm.Open(dialector, &gorm.Config{}); err != nil {
// 		log.Fatal(err)
// 		// panic("Could not connect with database!")
// 		// connect(dialector, 100)
// 	}
// 	log.Println("Database connected!!")
// }
// func connect(dialector gorm.Dialector, count uint) {
// 	var err error
// 	if DB, err = gorm.Open(dialector); err != nil {
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

//go:embed templates/*
var f embed.FS

func main() {

	// user := "root"
	// pw := "root_password"
	// db_name := "test_database"
	// var dsn string = fmt.Sprintf("%s:%s@tcp(mysql)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pw, db_name)
	// dialector := mysql.Open(dsn)

	// var err error
	// log.Println("Start Conneting Database...")
	// if DB, err = gorm.Open(dialector, &gorm.Config{}); err != nil {
	// 	log.Fatal(err)
	// 	panic("Could not connect with database!")
	// 	// connect(dialector, 100)
	// }
	// log.Println("Database connected!!")


	tmpl := template.Must(template.New("").ParseFS(f, "templates/*.html"))
	r := router.GetRouter()
	r.SetHTMLTemplate(tmpl)
	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}


// func connect(dialector gorm.Dialector, count uint) {
// 	var err error
// 	if DB, err = gorm.Open(dialector); err != nil {
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


// type Article struct {
// 	gorm.Model
// 	Id    int 		`json:"id"`
// 	Title string	`json:"title"`
// 	Body  string	`json:"body"`
// }


// func GetAll() (datas []Article) {
// 	result := DB.Find(&datas)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

// func ArticleIndex(ctx *gin.Context) {
// 	datas := GetAll()
// 	ctx.JSON(200, datas)
// }
