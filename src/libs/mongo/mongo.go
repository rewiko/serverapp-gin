package mongo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func main(g *gin.RouterGroup) {

}

func GetSession() *mgo.Session {

	hostname := viper.GetString("database.mongodb.host")



	if session == nil {
		conn, err := mgo.Dial(hostname)
		if err != nil {
			// Only warn since we'll retry later for each request
			fmt.Println("Could not connect to Mongo DB. Error: %s", err)
			panic(err.Error())
		} else {
			session = conn
		}
	}
	return session.Clone()

}

func MapMongo(c *gin.Context) {
	//s := session.Clone()
	s := GetSession()

	defer s.Close()

	c.Set("mongoSession", s.DB("db"))
	c.Next()
}
