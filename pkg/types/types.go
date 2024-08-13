package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type RecipeStore interface {
	GetRecipeByName(name string) (*Recipe, error)
	GetRecipeByID(id string) (*Recipe, error)
	CreateRecipe(Recipe) error
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

// User data stuff
type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Recipe data stuff
type Recipe struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageURL"`
	//Ingredients	[]string	`json:"ingredients"`
	//Steps		[]string	`json:"steps"`
	//VidLink		string		`json:"vidlink"`
	Source      string `json:"source"`
	Alterations string `json:"alterations"`
	TotalTime   string `json:"time"`
	Servings    int    `json:"servings"`
}

type PushRecipePayload struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"imageURL" validate:"required"`
	//Ingredients	[]string	`json:"ingredients"`
	//Steps		[]string	`json:"steps"`
	Source      string `json:"source" validate:"required"`
	Alterations string `json:"alterations"`
	TotalTime   string `json:"time" validate:"required"`
	Servings    int    `json:"servings" validate:"required"`
}
