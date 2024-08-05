package main

import (
	_ "profileService/docs"
	"profileService/internal/handlers"
)

func main() {
	contr := handlers.GetController()
	contr.Run()
}
