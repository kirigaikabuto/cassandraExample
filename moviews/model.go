package moviews

type Movie struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Rating int64 `json:"rating"`
}

type MovieUpdate struct {
	Id     int64   `json:"id"`
	Name   *string  `json:"name"`
	Rating *int64 `json:"rating"`
}
