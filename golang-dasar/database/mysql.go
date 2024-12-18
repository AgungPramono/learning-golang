package database

var connection string

// otomatis di jalankan ketika package digunakan
func init() {
	connection = "MYSQL"
}

func GetConnection() string {
	return connection
}
