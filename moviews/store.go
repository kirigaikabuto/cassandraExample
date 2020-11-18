package moviews

type MovieStore interface {
	Create(movie *Movie) (*Movie, error)
	Get(id int64) (*Movie, error)
	List() ([]Movie, error)
	Update(movie *MovieUpdate) (*Movie, error)
	Delete(id int64) error
}
