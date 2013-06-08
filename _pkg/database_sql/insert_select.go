package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
	 * Use "mysql" as driverName.
	 * Use a valid DSN as dataSourceName.
	 *
	 * Optional parts marked by squared brackets:
	 *     [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	 *
	 * A DSN in its fullest form:
	 *     username:password@protocol(address)/dbname?param=value
	 *
	 * Except the database name, all values are optional. So the mininal DSN is:
	 *     /dbname
	 *
	 * If you do not want to preselect a database leave dbname empty:
	 *     /
	 *
	 * Protocol - In general you should use an Unix domain socket if available
	 *            and TCP otherwise for best performance.
	 *
	 * Address  - For TCP and UDP networks, addresses have the form host:port.
	 *            If host is a literal IPv6 address, it must be enclosed in
	 *            square brackets.
	 *
	 * Parameters - Parameters are case-sensitive!
	 */
	db, err := sql.Open("mysql", "root:123456@/godemo?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO square_num values(?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM square_num WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err := stmtIns.Exec(i, i*i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error())
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d\n", squareNum)
}
