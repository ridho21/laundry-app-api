package testing

import (
	"errors"
)
// Positif Test Case -> Ekspetasi Testnya adalah fungsi berjalan sesuai validasi
// Negative Test Case -> Ekspetasi testnya fungsi harus terjadi error

func SayHello(name string)(string,error) {
	if name == ""{
		return "",errors.New("Name can't be empty")
	}
	if len(name) < 5 {
		return "",errors.New("Name minimal 5 Karakter")
	}
	return "Hello " + name,nil
}


// suffix _test file
// prefix function harus mengandung kata Test