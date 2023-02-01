package entities

type User struct {
	ID       int    `json:"id" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      string `json:"age" binding:"required"`
	UserRole string `json:"user_role" binding:"required"`
}

type Film struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Genre        string `json:"genre"`
	DirectorName string `json:"director_name"`
	Rate         string `json:"rate"`
	Year         string `json:"year"`
	Minutes      string `json:"minutes"`
}
