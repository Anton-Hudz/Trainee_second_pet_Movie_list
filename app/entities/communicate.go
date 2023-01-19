package entities

type QueryParams struct {
	Filter string
	Genre  string
	Rate   string
	Sort   string
	Limit  string
	Offset string
}

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
