package impl

import (
	. "github.com/guidomantilla/bookstore_users-api/common/db"
	. "github.com/guidomantilla/bookstore_users-api/common/sql"
	. "github.com/guidomantilla/bookstore_users-api/core/model"
)

const (
	STATEMENT_CREATE       = "INSERT INTO user (first_name, last_name, email, date, status, password) VALUES (?, ?, ?, ?, ?, ?)"
	STATEMENT_UPDATE       = "UPDATE user SET first_name = ?, last_name = ?, date = ? WHERE id = ?"
	STATEMENT_DELETE       = "DELETE FROM user WHERE id = ?"
	STATEMENT_FIND_BY_ID   = "SELECT id, first_name, last_name, email, date, status FROM user WHERE id = ?"
	STATEMENT_FIND         = "SELECT id, first_name, last_name, email, date, status FROM user"
	STATEMENT_EXISTS_BY_ID = "SELECT EXISTS(SELECT id from user WHERE id = ?)"
)

type DefaultUserRepository struct {
	mysqlDataSource DataSource
}

func NewDefaultUserRepository(mysqlDataSource DataSource) *DefaultUserRepository {
	return &DefaultUserRepository{
		mysqlDataSource: mysqlDataSource,
	}
}

func (userRepository *DefaultUserRepository) Create(user *User) error {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_CREATE)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.Date, user.Status, user.Password)
	if err != nil {
		return err
	}

	user.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *DefaultUserRepository) CreateBulk(users []*User) error {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_CREATE)
	if err != nil {
		return err
	}
	defer statement.Close()

	for _, user := range users {

		result, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.Date)
		if err != nil {
			return err
		}

		user.Id, err = result.LastInsertId()
		if err != nil {
			return err
		}
	}

	return nil
}

func (userRepository *DefaultUserRepository) Update(user *User) error {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_UPDATE)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Date, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *DefaultUserRepository) Delete(id int64) error {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_DELETE)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *DefaultUserRepository) FindById(id int64) (*User, error) {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_FIND_BY_ID)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var user User
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date, &user.Status); err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *DefaultUserRepository) Find(paramMap map[string][]string) (*[]User, error) {

	database := userRepository.mysqlDataSource.GetDatabase()

	whereCondition := BuildWhere(paramMap)
	statement, err := database.Prepare(STATEMENT_FIND + whereCondition)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	for rows.Next() {

		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Date, &user.Status)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (userRepository *DefaultUserRepository) ExistsById(id int64) (bool, error) {

	database := userRepository.mysqlDataSource.GetDatabase()

	statement, err := database.Prepare(STATEMENT_EXISTS_BY_ID)
	if err != nil {
		return false, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var exists int
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists == 1, nil
}
