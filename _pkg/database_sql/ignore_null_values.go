package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       uint
	Username string
	Age      uint
	Country  string
}

func (u *User) String() string {
	return fmt.Sprintf("Username=%s, Age=%d, Country=%s", u.Username, u.Age, u.Country)
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/godemo?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	statIns, err := db.Prepare("INSERT INTO user (`username`, `age`, `country`) values (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statIns.Close()

	data := []*User{
		&User{Username: "xingyuli", Age: 24, Country: "China"},
		&User{Username: "god", Age: 100, Country: "Paradise"},
		&User{Username: "lucy", Country: "America"}, // with no age
		&User{Username: "king", Age: 33},            // with no country
	}
	for _, user := range data {
		_, err := statIns.Exec(user.Username, user.Age, user.Country)
		if err != nil {
			fmt.Errorf("Failed to insert: %s\n", user)
		} else {
			fmt.Println("Inserted:", user)
		}
	}

	// Reading the whole table
	rows, err := db.Query("SELECT id, username, age, country from user")
	for rows.Next() {
		var id, age uint
		var username, country string
		if err := rows.Scan(&id, &username, &age, &country); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id=%d, username=%s, age=%d, country=%s\n", id, username, age, country)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
