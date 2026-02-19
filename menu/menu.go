package menu

import (
	"fmt"
	"koda-b6-golang/models"
)

func TampilkanMenu(menuData []models.Menu) {
	fmt.Println("=== MENU PIZZA ===")
	for i, item := range menuData {
		fmt.Printf("%d. %s - Rp%d\n", i+1, item.Name, item.Harga)
	}
}
