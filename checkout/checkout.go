package checkout

import (
	"fmt"
	"koda-b6-golang/input"
	"koda-b6-golang/keranjang"
	"koda-b6-golang/models"
)

func Checkout(cekKeranjang *[]models.ItemKeranjang) {
	if len(*cekKeranjang) == 0 {
		fmt.Println("keranjang kosong, silahkan berberlanja")
		return
	}

	keranjang.TampilkanKeranjang(*cekKeranjang)

	jawaban := input.TanyaUser("lanjut bayar? (y/n): ")
	if jawaban == "Y" || jawaban == "y" {
		totalBelanja := 0
		for _, item := range *cekKeranjang {
			subtotal := item.Harga * item.Jumlah
			totalBelanja += subtotal
		}

		fmt.Println("pembayaran berhasil")

		*cekKeranjang = []models.ItemKeranjang{}
	} else {
		fmt.Println("checkout dibatalkan")
	}

}
