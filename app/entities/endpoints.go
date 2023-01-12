package entities

type User struct {
	ID        int    `json:"id" db:"id"`
	Login     string `json:"login" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Age       string `json:"age" binding:"required"`
	User_Role string `json:"user_role" binding:"required"`
}

type Film struct {
	ID            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	Genre         string `json:"genre" binding:"required"`
	Director_Name string `json:"director_name" binding:"required"`
	Rate          int    `json:"rate" binding:"required"`
	Year          int    `json:"year" binding:"required"`
	Minutes       int    `json:"minutes" binding:"required"`
}

type Director struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Date_of_birth string `json:"date_of_birth"`
}

type Favourites struct {
	User_ID int `json:"user_id"`
	Film_ID int `json:"film_id"`
}

type Wishlist struct {
	User_ID int `json:"user_id"`
	Film_ID int `json:"film_id"`
}
