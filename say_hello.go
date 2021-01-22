package go_say_hello

// kita ubah func SayHello agar memiliki parameter, agar codingannya rusak
// karena orang yang memanggil func SayHello itu error semua, karena func tersebut diwajibkan memasukkan parameter
// sebelumnya tidak menggunakan parameter name
func SayHello(name string) string { // jangan lupa ini uppercase agar public
	return "Hello, I'm"+ name + "\nBerikut adalah code terbarunya"
}
