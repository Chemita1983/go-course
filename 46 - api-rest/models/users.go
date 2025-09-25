package models

import (
	"apirest/db"
	"database/sql"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UsersTableName string = "users"

const UserSchema string = `CREATE TABLE IF NOT EXISTS users (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// Inserta o actualiza un registro de usuario
func (user *User) Save() {

	// Cuando se crea un usuario nuevo se crea con id=0
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}

}

// Crea usuario e inserta en bdd
func CreateUser(username, password, email string) *User {
	newUser := newUser(username, password, email)
	newUser.insert()
	return newUser
}

// Obtenemos todos los usuarios
func ListUsers() (Users, error) {

	query := "SELECT id, username, password, email from users"
	users := Users{}
	if rows, err := db.Query(query); err != nil {
		return nil, err
	} else {

		for rows.Next() {
			user := User{}
			var valueEmail sql.NullString // Para valores nulos

			err := rows.Scan(&user.Id, &user.Username, &user.Password, &valueEmail)
			if err != nil {
				log.Println(err)
			}

			if valueEmail.Valid {
				user.Email = valueEmail.String
			} else {
				user.Email = "NONE"
			}

			users = append(users, user)
		}

		return users, nil
	}
}

func GetUserById(id int) (*User, error) {
	query := "SELECT id, username, password, email FROM users WHERE id=?"
	user := newUser("", "", "")

	if rows, err := db.Query(query, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var valueEmail sql.NullString // Para valores nulos

			rows.Scan(&user.Id, &user.Username, &user.Password, &valueEmail)

			if valueEmail.Valid {
				user.Email = valueEmail.String
			} else {
				user.Email = "NONE"
			}
		}
		return user, nil
	}
}

func (user *User) Delete() {
	query := "DELETE FROM users WHERE Id=?"
	_, err := db.Exec(query, user.Id)
	if err != nil {
		log.Println(err)
	}
}

// Construir usuario
func newUser(username, password, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// Inserta un registro
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	// Asignamos el ultimo registro insertado
	user.Id, _ = result.LastInsertId()
}
