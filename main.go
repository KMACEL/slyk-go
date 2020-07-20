package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	// ehaIILDIbu0kzNi5CqWYOlYZru3yHXX0
	resp, err := client.R().
		SetHeader("apikey", "ehaIILDIbu0kzNi5CqWYOlYZru3yHXX0").
		Get("https://api.slyk.io/users")

	if err != nil {
		fmt.Println("Error : ", err.Error())
		return
	}

	if resp.IsError() {
		fmt.Println("Error 2 : ", resp.StatusCode())
		return
	}

	fmt.Println("Response : ", string(resp.Body()))

}
