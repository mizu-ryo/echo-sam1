package main

import (
  "github.com/mizu-ryo/echo-sam1/src/seed"
  "github.com/mizu-ryo/echo-sam1/src/router"
)

func main() {
  seed.Seed()
  router.Router()
}
