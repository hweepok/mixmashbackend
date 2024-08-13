package db

import (
	"database/sql"
	"fmt"

	"github.com/hweepok/mixmashbackend/pkg/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil

}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, email, password) VALUES(?,?,?)",
		user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetRecipeByName(name string) (*types.Recipe, error) {
	return nil, nil
}

func (s *Store) GetRecipeByID(id string) (*types.Recipe, error) {
	return nil, nil
}

func (s *Store) CreateRecipe(recipe types.Recipe) error {
	_, err := s.db.Exec("INSERT INTO recipes(id, name, description, imageURL, source, alterations, time, servings) VALUES(?,?,?,?,?,?,?,?)",
		recipe.ID, recipe.Name, recipe.Description, recipe.ImageUrl, recipe.Source, recipe.Alterations,
		recipe.TotalTime, recipe.Servings)
	if err != nil {
		return err
	}
	return nil
}
