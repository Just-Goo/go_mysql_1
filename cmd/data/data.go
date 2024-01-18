package data

import (
	"github.com/Just-Goo/Go-MySql-1/cmd/config"
	"github.com/Just-Goo/Go-MySql-1/cmd/models"
)

func InsertStudent(firstName, lastName string, age int) error {

	stmt, err := config.MyApp.DB.Prepare("INSERT INTO students (firstname, lastname, age) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(firstName, lastName, age)
	if err != nil {
		return err
	}

	rowsAffec, err := result.RowsAffected()
	if err != nil || rowsAffec != 1 {
		return err
	}

	return nil
}

func InsertStudentReturnID(firstName, lastName string, age int) (Id int, err error) {

	query := "INSERT INTO students (firstname, lastname, age) VALUES (?, ?, ?) RETURNING id"

	err = config.MyApp.DB.QueryRow(query, firstName, lastName, age).Scan(&Id)
	if err != nil {
		return Id, err
	}

	return Id, nil
}

func GetAllStudents() ([]models.User, error) {
	var users []models.User

	stmt, err := config.MyApp.DB.Prepare("SELECT * FROM students")
	if err != nil {
		return users, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var u models.User

		err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	if rows.Err() != nil {
		return users, err
	}

	return users, nil
}

func UpdateStudent(firstName, lastName string, id, age int) error {
	stmt, err := config.MyApp.DB.Prepare("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(firstName, lastName, age, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return err
	}

	return nil

}

func DeleteStudent(id int) error {
	stmt, err := config.MyApp.DB.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return err
	}

	return nil
}
