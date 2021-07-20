package seed

import (
  "fmt"
  "github.com/mizu-ryo/echo-sam1/src/db"
  "github.com/mizu-ryo/echo-sam1/src/models"
)

func Seed() {

  db := db.Init()

  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Template{})

  user := models.User{
    Name:  "mizu",
    Email: "mizu@email.com",
    Templates: []models.Template{{ID:1, Title: "sample"},{ID:2,Title:"sample2"}},
  }
  db.Create(&user)

//  template := models.Template{
//    Title: "sample1",
//    UserID: 1,
//  }
//  db.Create(&template)

  fmt.Println("create seed")

  defer db.Close()
}
