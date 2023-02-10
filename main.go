package main

/*
jalankan:
	go mod init namaproject

download mysql db driver:
	go get -u github.com/go-sql-driver/mysql


*/

import (
	"be15/rawsql/config"
	"be15/rawsql/controllers"
	"be15/rawsql/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.ConnectToDB()
	defer db.Close()

	fmt.Println("Pilih Menu:\n1. Read Data\n2. Add Data")
	fmt.Println("Input pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		controllers.GetAllGuru(db)

	// CASE 2
	case 2:
		//INSERT DATA
		newGuru := entities.Guru{}
		fmt.Println("Masukkan ID Guru:")
		fmt.Scanln(&newGuru.Id)
		fmt.Println("Masukkan Nama Guru:")
		fmt.Scanln(&newGuru.Nama)
		fmt.Println("Masukkan Telepon Guru:")
		fmt.Scanln(&newGuru.Telepon)
		fmt.Println("Masukkan Email Guru:")
		fmt.Scanln(&newGuru.Email)

		controllers.InsertGuru(db, newGuru)
		

	// CASE 3
	case 3:
		fmt.Println("update")

	// CASE 4
	case 4:
		fmt.Println("delete")

	// CASE 5
	case 5:
		fmt.Println("get by id")

	// DEFAULT
	default:
		fmt.Println("Input tidak sesuai")

	}

}
