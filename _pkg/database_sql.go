package main

import (
	"database/sql"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

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
}
