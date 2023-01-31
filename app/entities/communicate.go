package entities

type FilmResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Genre         string  `json:"genre"`
	Director_Name string  `json:"director_name"`
	Rate          float32 `json:"rate"`
	Year          int     `json:"year"`
	Minutes       int     `json:"minutes"`
}

type GetAllFilmsResponse struct {
	Data []FilmResponse `json:"data"`
}

type LogInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type QueryParams struct {
	Format string
	Genre  string
	Rate   string
	Sort   string
	Limit  string
	Offset string
}

type ListsComponents struct {
	List  string
	Movie string
}
