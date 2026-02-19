package models

import "fmt"

type Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"harga"`
}

type ItemKeranjang struct {
	Nama   string
	Harga  int
	Jumlah int
}

type Tampilable interface {
	Tampilkan()
}

func (m Menu) Tampilkan() {
	fmt.Printf("ID: %d | %s - Rp%d\n", m.Id, m.Name, m.Harga)
}

func (i ItemKeranjang) Tampilkan() {
	subtotal := i.Harga * i.Jumlah
	fmt.Printf(" %s x%d = Rp%d\n", i.Nama, i.Jumlah, subtotal)
}
