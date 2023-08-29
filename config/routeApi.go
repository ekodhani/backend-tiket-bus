package config

import (
	controllers "backend-pertama/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))

	//Portal Dhani
	pd := router.Group("/pd")
	{
		// Version 1
		pd_v1 := pd.Group("/v1")
		{
			// Login
			pd_v1.POST("/signup", controllers.Signup)
			pd_v1.POST("/signin", controllers.Signin)
			pd_v1.POST("/getKursi", controllers.GetKursi)
			pd_v1.POST("/getBus", controllers.GetBus)
			pd_v1.POST("/savetiket", controllers.SaveTiket)
			pd_v1.POST("/getbelumbayar", controllers.GetBelumBayar)
			pd_v1.POST("/getdetailorder", controllers.GetDetailOrder)
		}
	}
	return router
}
