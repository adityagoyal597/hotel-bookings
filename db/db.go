package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "hotel_booking.db")

	if err != nil {
		log.Fatalf("Could not connect to the Database: %v", err)
	}

	err = DB.Ping() // ENSURES THE CONNECTION IS VALID
	// RETURN AN ERROR IF THERE IS AN ISSUE WITH THE CONNECTION

	if err != nil { // CONNECTION NOT VALID
		log.Fatalf("Connection isn't Established: %v", err)
	}

	fmt.Println("Connection Successfully Established!")

	createTables()
}

func createTables() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable) // Exec RETURNS RESULT

	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}

	createRoomTable := `
	CREATE TABLE IF NOT EXISTS rooms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		room_number TEXT NOT NULL,
		type TEXT NOT NULL,
		price FLOAT
	)
	`
	_, err = DB.Exec(createRoomTable)

	if err != nil {
		log.Fatalf("Could not create rooms table: %v", err)
	}

	createBookingTable := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		room_id INTEGER,
		checkin_date DATE,
		checkout_date DATE,
		FOREIGN KEY (user_id) REFERENCES user(id),
		FOREIGN KEY (room_id) REFERENCES rooms(id)
	)
	`
	_, err = DB.Exec(createBookingTable)

	if err != nil {
		log.Fatalf("Could not create bookings table: %v", err)
	}

}
