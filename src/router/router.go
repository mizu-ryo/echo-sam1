package router

import (
  "github.com/gin-gonic/gin"
  "github.com/mizu-ryo/echo-sam1/src/controllers"
)

func Router() {
  router := gin.Default()

  router.GET("/", func(ctx *gin.Context){
    ctx.JSON(200, gin.H{
      "message":"Hello, World!",
    })
  })

  u := router.Group("/users")
  {
    ctrl := controllers.UserController{}
    u.GET("/index", ctrl.Index)
    u.POST("/new", ctrl.Create)
    u.DELETE("/delete/:id", ctrl.Delete)
  }

  router.Run()
}
