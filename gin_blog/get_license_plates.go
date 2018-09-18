package main

import "github.com/jinzhu/gorm"
import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"os"
)


type SentryMessage struct {
	ID int
	Message string
	Datetime time.Time
	Data string
	Group_id int
	Message_id string
	Project_id int
	Time_spent int
	Platform string

}

func (SentryMessage) TableName() string {
	return "sentry_message"
}


func main() {
	var messages []SentryMessage
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"user",
		"sentry",
		"127.0.01",
		"sentry"))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.Where("message LIKE ?", "%car_license_plate%").Find(&messages)
	f, err := os.OpenFile("license.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, msg := range messages {
		//fmt.Println(msg.Message)
		start := strings.Index(msg.Message, "QK")
		fmt.Println(msg.Message[start:start+14])
		f.WriteString(msg.Message[start:start+14] + "\n")
	}
	fmt.Println("finish")


}
