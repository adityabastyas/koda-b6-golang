package models

import "fmt"

type Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Harga int    `json:"harga"`
}

type Tampilable interface {
	Tampilkan()
}

func (m Menu) Tampilkan() {
	fmt.Printf("ID: %d | %s - Rp%d\n", m.Id, m.Name, m.Harga)
}
