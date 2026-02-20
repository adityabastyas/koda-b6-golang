package keranjang

import (
	"fmt"
	"koda-b6-golang/input"
	"koda-b6-golang/menu"
	"koda-b6-golang/models"
	"strconv"
)

func TampilkanKeranjang(keranjang []models.ItemKeranjang) {
	fmt.Println("=== KERANJANG BELANJA ===")
	if len(keranjang) == 0 {
		fmt.Println("keranjang kosong")
		return
	}

	totalBelanja := 0

	for i, item := range keranjang {
		subtotal := item.Harga * item.Jumlah
		totalBelanja += subtotal
		fmt.Printf("%d. %s\n", i+1, item.Nama)
		fmt.Printf("  %dx Rp%d = Rp%d\n", item.Jumlah, item.Harga, subtotal)
	}
	fmt.Printf("TOTAL: Rp%d\n", totalBelanja)

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

func HapusKeranjang(keranjang *[]models.ItemKeranjang) {
	if len(*keranjang) == 0 {
		fmt.Println("keranjang kosong")
		return
	}

	TampilkanKeranjang(*keranjang)

	jawaban := input.TanyaUser("pilih nomor atau 0 untuk batal: ")
	nomor, err := strconv.Atoi(jawaban)
	if err != nil || nomor < 0 || nomor > len(*keranjang) {
		fmt.Println("nomor salah")
		return
	}

	if nomor == 0 {
		fmt.Println("batal menghapus")
		return
	}

	namaItem := (*keranjang)[nomor-1].Nama

	*keranjang = append((*keranjang)[:nomor-1], (*keranjang)[nomor:]...)
	fmt.Printf("berhasil hapus: %s\n", namaItem)
}
