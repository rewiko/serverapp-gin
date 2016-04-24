package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rewiko/gin-app/libs/admin"
	"github.com/rewiko/gin-app/libs/config"
	"github.com/rewiko/gin-app/libs/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	r := gin.Default()
	r.Use(mongo.MapMongo)

	r.GET("/", func(c *gin.Context) {
		db := c.MustGet("mongoSession").(*mgo.Database)
		collection := db.C("people")
		err := collection.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			fmt.Println(err)
		}

		result := Person{}
		err = collection.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			fmt.Println(err)
		}
		//log.Fatal("test")
		//fmt.Println("Phone:", result)

		c.JSON(200, result)
	})

	//admin.Main(r)
	admin.Main(r, r.Group("/admin"))

	fmt.Println("Routes: ", r.Routes())
	config.SetConfig()
	setupDatabase()
	//jsonapi.Run()

	r.Run(":8081") // listen and server on 0.0.0.0:8080
}

func setupDatabase() {
	fmt.Println("Setup Database!")
	mongo.GetSession()
}
