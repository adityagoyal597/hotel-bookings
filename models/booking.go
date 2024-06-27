package models

import (
	"log"
	"time"

	"example.com/Hotel_Bookings/db"
)

type Booking struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	RoomID       int       `json:"room_id"`
	CheckInDate  time.Time `json:"checkin_date"`
	CheckOutDate time.Time `json:"checkout_date"`
}

func (booking *Booking) Save() error {
	query := `
	INSERT INTO bookings(id,user_id,room_id,checkin_date,checkout_date)
	VALUES (?,?,?,?,?)
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Preparing Query: %v", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(booking.ID, booking.UserID, booking.RoomID, booking.CheckInDate, booking.CheckOutDate)

	if err != nil {
		log.Printf("Error Executing Query: %v", err)
		return err
	}

	return nil
}

func GetBookingsByUserID(userID int) ([]Booking, error) {
	query := "SELECT * FROM bookings WHERE user_id=?"

	rows, err := db.DB.Query(query)

	if err != nil {
		log.Printf("Error Querying Bookings: %v", err)
		return nil, err
	}

	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.RoomID, &booking.CheckInDate, &booking.CheckOutDate)

		if err != nil {
			log.Printf("Error Scanning Bookings: %v", err)
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	return bookings, nil
}
