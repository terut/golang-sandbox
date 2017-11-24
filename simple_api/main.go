package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	r.GET("/users", func(c *gin.Context) {
		var (
			users   []User
			items   []Item
			jsonMap map[string]interface{} = make(map[string]interface{})
		)

		db.Find(&users)

		for i, user := range users {
			db.Model(&user).Related(&items)
			user.Items = items
			users[i] = user
		}

		jsonMap["users"] = users
		c.JSON(200, jsonMap)
	})

	r.Run(":8080")
}

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type Item struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

func init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Item{})

	itemNames := []string{
		"Game1",
		"Game2",
		"Game3",
		"Game4",
		"Game5",
		"Game6",
		"Game7",
	}

	userNames := []string{
		"User1",
		"User2",
		"User3",
		"User4",
		"User5",
		"User6",
		"User7",
		"User8",
		"User9",
		"User10",
	}

	users := CreateUsers(userNames, itemNames)

	count := 0
	db.Table("users").Count(&count)
	if count == 0 {
		for _, user := range users {
			db.Create(&user)
		}
	}
}

func CreateUsers(userNames []string, itemNames []string) []User {
	users := make([]User, len(userNames))

	for i, name := range userNames {
		users[i] = NewUser(name, itemNames)
	}

	return users
}

func NewUser(name string, itemNames []string) User {
	return User{
		Name:  name,
		Items: CreateItems(itemNames),
	}
}

func CreateItems(itemNames []string) []Item {
	items := make([]Item, len(itemNames))

	for i, item := range itemNames {
		items[i] = NewItem(item)
	}

	return items
}

func NewItem(item string) Item {
	rand.Seed(time.Now().UnixNano())

	return Item{
		Name:  item,
		Score: rand.Intn(10) + 1,
	}
}
