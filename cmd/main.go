package main

import (
	"profileService/internal/handlers"
)

func main() {
	contr := handlers.GetController()
	contr.Run()
}
