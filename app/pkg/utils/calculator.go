package utils

// TambahDuit diawali huruf besar agar bisa di-import (Public)
func TambahDuit(awal float64, tambahan int) float64 {
	return awal + float64(tambahan)
}

func KurangDuit(awal float64, pengurangan int) float64 {
	return awal - float64(pengurangan)
}
