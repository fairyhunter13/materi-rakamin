package main

import (
	"fmt"
	"time"
)

type Product struct {
	barcodeNumber string
	price         int64
	name          string
}

func (p *Product) information() string {
	return fmt.Sprintf("Informasi Produk, nomor barcode: %s, harga: %d, nama: %s.", p.barcodeNumber, p.price, p.name)
}

type SamyangNoodle struct {
	Product
	variant     string
	isFried     bool
	color       string
	ExpiredDate time.Time
}

func (sn *SamyangNoodle) PrintInformation() {
	fmt.Printf("%s\n", sn.information())
	fmt.Printf("Varian: %s.\n", sn.variant)
	fmt.Printf("Goreng: %t.\n", sn.isFried)
	fmt.Printf("Warna: %s.\n", sn.color)
	expired := sn.ExpiredDate
	expiredString := fmt.Sprintf("%v.%v.%v", expired.Day(), int(expired.Month()), expired.Year())
	fmt.Printf("Tanggal Kedaluwarsa: %s.\n", expiredString)
}

func inheritance() {
	cheeseSamyang := SamyangNoodle{
		Product: Product{
			barcodeNumber: "1239089213",
			price:         25000,
			name:          "Mi Samyang Cheese Extra Hot",
		},
		variant:     "cheese",
		isFried:     true,
		color:       "kuning",
		ExpiredDate: time.Now().Add(12 * 30 * 24 * time.Hour),
	}
	cheeseSamyang.PrintInformation()
}
