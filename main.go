package main

import (
	"log"
	"os"

	"github.com/ishinu/controllers"
	"github.com/ishinu/database"
	"github.com/ishinu/middleware"
	"github.com/ishinu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.GET("/addaddress", controllers.AddAddress())
	router.GET("/edithomeaddress", controllers.EditHomeAddress())
	router.GET("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}
