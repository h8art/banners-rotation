package api

import (
	"banner-otus/handlers"
	"github.com/gin-gonic/gin"
)

func RunApi() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		banner := v1.Group("/banner")
		{
			banner.GET("/", handlers.GetBanners)
			banner.POST("/", handlers.AddBanner)
			banner.DELETE("/:id", handlers.DeleteBanner)
			banner.GET("/:slotId/:groupId", handlers.GetBanner)
		}
		v1.GET("/group", handlers.GetGroups)
		v1.GET("/slot", handlers.GetSlots)
		v1.POST("/slot", handlers.AddSlot)
		v1.POST("/group", handlers.AddGroup)
		v1.GET("/visit/:bannerId/:groupId", handlers.VisitBanner)
	}
	_ = router.Run()
}
