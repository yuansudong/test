package main

import (
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Articles 文章
type Articles struct {
	ID    int64  `xorm:"id"`
	Title string `xorm:"title"`
	Body  string `xorm:"content"`
}

func main() {

}

// Md2Pdf Markdown 转为 P
func Md2Pdf() {

}

func GetTitle() string {
	str := ""
	for i := 0; i < 4; i++ {
		num := rand.Int31n(33060) + 30000
		if num > 40000 {
			num = 40000
		}
		str += string(rune(num))
	}
	return str
}

func InsertDataset() {
	db, err := xorm.NewEngine("mysql", "root:ysd941018@tcp(127.0.0.1:3306)/cc")
	if err != nil {
		log.Fatalln(err.Error())
	}
	length := 10000
	mA := make([]Articles, length)
	for i := 0; i < 4700; i++ {
		log.Println(i)
		for j := 0; j < length; j++ {
			mA[j].Title = GetTitle()
			mA[j].Body = "固定内容"
		}
		if _, err := db.Insert(mA); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
