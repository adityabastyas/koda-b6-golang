package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"harga"`
}

func main() {

	fmt.Println("Selamat Datang di Pizza Hut Indonesia")

	url := "https://raw.githubusercontent.com/adityabastyas/koda-b6-golang/refs/heads/main/assets/data/menu.json"

	var menuData []Menu

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Body :", err)
		return
	}

	err = json.Unmarshal(body, &menuData)
	if err != nil {
		fmt.Println("Error nmarshal:", err)
	}

	for i := range menuData {
		fmt.Println(
			menuData[i].Id,
			"-",
			menuData[i].Name,
			"-",
			menuData[i].Harga,
		)
	}

}
