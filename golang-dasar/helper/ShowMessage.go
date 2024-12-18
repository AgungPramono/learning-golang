package helper

/**
Exported (Public)
Dimulai dengan huruf kapital.
Bisa diakses dari luar package.
Contoh: FuncName, VariableName'

Unexported (Private)
Dimulai dengan huruf kecil.
Hanya bisa diakses dari dalam package yang sama.
Contoh: funcName, variableName.
*/

var version string = "1.0.1"           //tidak bisa di akses dari luar
var Application = "golang application" // bisa di akses dari luar

func ViewMessage(message string) string {
	return "Your Message: " + message
}

func showVersion() string { //tidak di akses dari luar package, hanya bisa dari package yang sama
	return "version: " + version
}
