package models

import (
	"log"

	"example.com/Hotel_Bookings/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}

func (user *User) Save() error {
	query := "INSERT INTO users (name,email,password) VALUES (?,?,?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		log.Printf("Error Preparing Query: %v", err)
		return err
	}

	defer statement.Close()

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		log.Printf("Error While Hashing Password: %v", err)
		return err
	}
	_, err = statement.Exec(user.Name, user.Email, hashedPassword)

	if err != nil {
		log.Printf("Error Executing Query: %v", err)
		return err
	}

	return nil
}

// GETTING A BUNCH OF DATA FROM A PIECE OF DATA IS A FUNCTION

func GetUserByEmail(email string) (*User, error) {
	query := "SELECT id,name , email , password FROM users WHERE email =?"

	row := db.DB.QueryRow(query, email) // db.DB.QueryRow(query,VALUE FOR THE PLACE HOLDER) ; EXECUTES A QUERY THAT RETURNS ATMOST ONE ROW

	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
