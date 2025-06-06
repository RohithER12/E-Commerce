package repo

import (
	"database/sql"
	"fmt"

	"github.com/janeefajjustin/ecommerce/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) SaveUser(user models.User) error {
	//var usr models.User
	query := `INSERT INTO users(username,password,phonenumber,email,firstname,middlename,lastname)
	 VALUES($1,$2,$3,$4,$5,$6,$7)`

	row := u.db.QueryRow(query, user.Username, user.Password, user.PhoneNumber, user.Email, user.FirstName, user.MiddleName, user.LastName)
	if err := row.Err(); err != nil {
		fmt.Println("error is from query")
		return err
	}

	return nil

}

func (u *UserRepo) CheckUniqueUserName(username string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = $1`
	rows, err := u.db.Query(query, username)
	if err != nil {
		fmt.Println("error is from save user repo function check username")
		return -1, err
	}
	rows.Scan(&count)
	
	return count, nil

}

func (u *UserRepo) CheckUniquePhoneNumber(phonenumber string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE phonenumber = $1`
	rows, err := u.db.Query(query, phonenumber)
	if err != nil {
		return -1, err
	}
	
	rows.Scan(&count)
	return count, nil
}

func (u *UserRepo) CheckUniqueEmailID(emailId string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1`
	rows, err := u.db.Query(query, emailId)
	if err != nil {
		return -1, err
	}
	rows.Scan(&count)
	return count, nil

}
