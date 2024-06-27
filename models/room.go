package models

import (
	"log"

	"github.com/adityagoyal597/hotel-bookings/db"
)

type Room struct {
	ID         int     `json:"id"`
	RoomNumber string  `json:"room_number"`
	Type       string  `json:"type"`
	Price      float64 `json:"price"`
}

func (room *Room) Save() error {
	query := "INSERT INTO rooms(room_number,type,price) VALUES(?,?,?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Preparing Query: %v", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(room.RoomNumber, room.Type, room.Price)

	if err != nil {
		log.Printf("Error Executing Query: %v", err)
		return err
	}

	return nil
}

func GetAllRooms() ([]Room, error) {
	query := "SELECT id,room_number,type,price FROM rooms"

	rows, err := db.DB.Query(query)

	if err != nil {
		log.Printf("Error Querying Rooms: %v", err)
		return nil, err
	}

	defer rows.Close()

	var rooms []Room // CREATING A SLICE OF ROOM FOR GETTING ALL ROOMS

	for rows.Next() { // RETURNS TRUE AS LONG AS THERE ARE ROWS LEFT
		var room Room
		err := rows.Scan(&room.ID, &room.RoomNumber, &room.Type, &room.Price)

		if err != nil {
			log.Printf("Error Scanning Rooms: %v", err)
			return nil, err
		}

		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (room *Room) Update() error {
	query := `
	UPDATE rooms 
	SET room_number=?,type=?,price=?
	WHERE id=?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Preparing Query: %v", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(room.RoomNumber, room.Type, room.Price, room.ID)

	if err != nil {
		log.Printf("Error Executing Error: %v", err)
		return err
	}

	return nil
}

func Delete(id int) error {
	query := "DELETE FROM rooms WHERE id=?"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Deleting room: %v", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		log.Printf("Error Executing Query: %v", err)
		return err
	}

	return nil
}
