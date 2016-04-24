package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rewiko/gin-app/libs/admin/model"
	"github.com/rewiko/gin-app/libs/admin/resource"
	"github.com/rewiko/gin-app/libs/admin/storage"
	"github.com/rewiko/gin-app/libs/jsonapi/api2go"
	"github.com/rewiko/gin-app/libs/jsonapi/api2go-adapter/gingonic"
)

//func Main(g *gin.Engine) {
//g.GET("/admin", func(c *gin.Context) {
//c.JSON(200, gin.H{
//"message": "admin",
//})
//})
//fmt.Printf("pass administration")
//}

func Main(g *gin.Engine, r *gin.RouterGroup) {

	api := api2go.NewAPIWithRouting(
		"v1/admin",
		api2go.NewStaticResolver("/"),
		gingonic.New(g),
	)

	db, err := storage.InitDB()
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	userStorage := storage.NewUserStorage(db)
	chocStorage := storage.NewChocolateStorage(db)
	api.AddResource(model.User{}, resource.UserResource{ChocStorage: chocStorage, UserStorage: userStorage})
	api.AddResource(model.Chocolate{}, resource.ChocolateResource{ChocStorage: chocStorage, UserStorage: userStorage})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "admin",
		})
	})
	//admin(r)

	fmt.Println("Administration!")

}

func admin(r *gin.RouterGroup) {
	r.GET("/test/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": id,
		})
	})
}

//func admin(inner func(*gin.Context), name string) func(*gin.Context) {
//return func(c *gin.Context) {
//start := time.Now()

//inner(c)

//log.Printf(
//"%s\t%s\t%s\t%s",
//c.Request.Method,
//c.Request.RequestURI,
//name,
//time.Since(start),
//)
//}
//}
