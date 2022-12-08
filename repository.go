package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var DataBase *sql.DB

func connect() {
	d, err := sql.Open("mysql", "root:local@/authentication")

	if err != nil {
		panic(err)
	}

	DataBase = d
}

func GetHash(username *string) []byte {
	row := DataBase.QueryRow("SELECT password_hash FROM authentication_data WHERE username=?", *username)

	var passwordHash []byte

	err := row.Scan(&passwordHash)

	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return passwordHash
}

func Insert(username *string, passwordHash *[]byte) {
	_, _ = DataBase.Exec("INSERT INTO authentication_data(username, password_hash) VALUES(?, ?)", *username, *passwordHash)
}

func Update(username *string, passwordHash *[]byte) {
	_, _ = DataBase.Exec("UPDATE authentication_data SET password_hash=? WHERE username=?", *passwordHash, *username)
}
