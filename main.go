package main

import (
	"fmt"
	"log"

	"github.com/giovane-aG/imersaofsfc2-simulator/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

}

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	jsonString, _ := route.ExportJsonPositions()
	fmt.Println(jsonString[0])
}
