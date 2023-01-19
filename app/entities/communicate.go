package entities

type QueryParams struct {
	Filter string
	Genre string
	Rate string
	Sort   string
	Limit  string
	Offset string
}

type FilmFromDB struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Genre       string  `json:"genre"`
	Director_id int     `json:"director_id"`
	Rate        float32 `json:"rate"`
	Year        int     `json:"year"`
	Minutes     int     `json:"minutes"`
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
	Data []FilmFromDB `json:"data"`
}
