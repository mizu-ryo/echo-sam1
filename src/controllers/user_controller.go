package controllers

import (
  "fmt"
  "strconv"
  "github.com/gin-gonic/gin"
  "github.com/mizu-ryo/echo-sam1/src/db"
  "github.com/mizu-ryo/echo-sam1/src/models"
)

// Controller is user controlller
type UserController struct{}

// Index action: GET /users
func (pc UserController) Index(ctx *gin.Context) {
    db := db.Init()
    var users []models.User
    db.Find(&users)
    defer db.Close()

    ctx.JSON(200, gin.H{
      "users": users,
    })
}


// Create action: POST /user
func (pc UserController) Create(ctx *gin.Context) {
    db := db.Init()
    name := ctx.PostForm("name")
    email := ctx.PostForm("email")
    fmt.Println("create user " + name + " with email " + email)
    db.Create(&models.User{Name: name, Email: email})
    defer db.Close()

    ctx.Redirect(302, "/")
}


func (pc UserController) Delete(ctx *gin.Context) {
  db := db.Init()
  n := ctx.Param("id")
  id, err := strconv.Atoi(n)
  if err != nil {
    panic("id is not a number")
  }
  var user models.User
  db.First(&user, id)
  db.Delete(&user)
  fmt.Println("delete user")
  defer db.Close()

  ctx.Redirect(302, "/")

}

