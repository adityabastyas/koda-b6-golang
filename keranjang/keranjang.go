package keranjang

import (
	"fmt"
	"koda-b6-golang/input"
	"koda-b6-golang/menu"
	"koda-b6-golang/models"
	"strconv"
)

func TampilkanKeranjang(keranjang []models.ItemKeranjang) {

}

func TambahKeKeranjang(menuData []models.Menu, keranjang *[]models.ItemKeranjang) {
	menu.TampilkanMenu(menuData)

	jawab := input.TanyaUser("pilih nomor atau 0 untuk bata1: ' ")
	nomor, err := strconv.Atoi(jawab) // strig jdi int
	if err != nil || nomor < 0 || nomor > len(menuData) {
		fmt.Println("nomor salah")
		return
	}
	if nomor == 0 {
		fmt.Println("batal")
		return
	}

	pizza := menuData[nomor-1]

	jawabanJumlah := input.TanyaUser("mau beli berapa? ")
	jumlah, err := strconv.Atoi(jawabanJumlah)
	if err != nil || jumlah < 1 {
		fmt.Println("jumlah minimal 1")
		return
	}

	*keranjang = append(*keranjang, models.ItemKeranjang{
		Nama:   pizza.Name,
		Harga:  pizza.Harga,
		Jumlah: jumlah,
	})

	fmt.Println("berhasil ditambakan")
	fmt.Printf("%s x%d = Rp%d\n", pizza.Name, jumlah, pizza.Harga*jumlah)
}
