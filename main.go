package main

import (
	"encoding/json"
	"fmt"
	"io"
	"koda-b6-golang/checkout"
	"koda-b6-golang/input"
	"koda-b6-golang/keranjang"
	"koda-b6-golang/menu"
	"koda-b6-golang/models"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Selamat Datang di Pizza Hut Indonesia")

	url := "https://raw.githubusercontent.com/adityabastyas/koda-b6-golang/refs/heads/main/assets/data/menu.json"

	var menuData []models.Menu
	var kranjang []models.ItemKeranjang

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Body :", err)
		os.Exit(1)
	}

	err = json.Unmarshal(body, &menuData)
	if err != nil {
		fmt.Println("Error nmarshal:", err)
		os.Exit(1)
	}

	if len(menuData) == 0 {
		panic("data tidak ada")
	}

	menuUtama(menuData, &kranjang)

}

func menuUtama(menuData []models.Menu, kranjang *[]models.ItemKeranjang) {
	for {
		fmt.Println("\n== MENU UTAMA ==")
		fmt.Println("1. menu pizza")
		fmt.Println("2. tambah ke keranjang")
		fmt.Println("3. lihat keranjang")
		fmt.Println("4. hapus item keranjang")
		fmt.Println("5. checkout")
		fmt.Println("6. lihat histori")
		fmt.Println("7. keluar")

		pilihan := input.TanyaUser("pilih nomer menu: ")

		switch pilihan {
		case "1":
			menu.TampilkanMenu(menuData)
		case "2":
			keranjang.TambahKeKeranjang(menuData, kranjang)
		case "3":
			keranjang.TampilkanKeranjang(*kranjang)
		case "4":
			keranjang.HapusKeranjang(kranjang)
		case "5":
			checkout.Checkout(kranjang)
		}
	}
}
