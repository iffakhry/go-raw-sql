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
	"log"

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

		var query = "INSERT INTO guru(id, nama, telepon, email) VALUES(?,?,?,?)"
		statement, errPrepare := db.Prepare(query)
		if errPrepare != nil {
			log.Fatal("error prepare insert", errPrepare.Error())
		}

		result, errInsert := statement.Exec(newGuru.Id, newGuru.Nama, newGuru.Telepon, newGuru.Email)
		if errInsert != nil {
			log.Fatal("error exec insert", errInsert.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("proses berhasil dijalankan")
			} else {
				fmt.Println("proses gagal")
			}
		}

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
