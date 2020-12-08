package main

import (
	_ "./model"
	"./lib/random"
	_ "./service"
)

func main() {
	random.GetId(1234567890, 100000, 8)
	//select{}
	//service.InitChat("666")
}