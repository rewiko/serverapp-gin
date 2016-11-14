package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/rewiko/app/libs/admin"
	"github.com/rewiko/app/libs/config"
	"github.com/gocql/gocql"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	r := gin.Default()
	//r.Use(mongo.MapMongo)

	r.GET("/", func(c *gin.Context) {
		//fmt.Println("Phone:", result)

		c.JSON(200, "test")
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
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "myproject"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
	"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)

	}

	//mongo.GetSession()
}
