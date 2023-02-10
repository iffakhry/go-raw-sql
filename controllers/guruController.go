package controllers

import (
	"be15/rawsql/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetAllGuru(db *sql.DB) {
	// SELECT DATA
	// select * from guru
	rows, errSelect := db.Query("SELECT id, nama, telepon, email FROM guru")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allGuru []entities.Guru
	for rows.Next() {
		var datarow entities.Guru
		errScan := rows.Scan(&datarow.Id, &datarow.Nama, &datarow.Telepon, &datarow.Email)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allGuru = append(allGuru, datarow)
		// fmt.Println(datarow)
	}

	// fmt.Println(allGuru)
	for i, v := range allGuru {
		fmt.Println("i:", i, "v:", v.Id, v.Nama)
	}
}

func InsertGuru(db *sql.DB, newGuru entities.Guru) {
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
}
