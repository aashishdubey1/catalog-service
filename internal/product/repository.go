package product

import "github.com/jackc/pgx/v5"

type ProductRepository interface {
	GetById(id string) (*Product, error)
	GetAll() ([]*Product, error)
	Create(product Product) error
}

type pgProductRepo struct {
	db *pgx.Conn
}

func NewPgProductRepo(db *pgx.Conn) ProductRepository {
	return &pgProductRepo{db: db}
}

func (r *pgProductRepo) GetById(id string) (*Product, error) {
	return &Product{}, nil
}

func (r *pgProductRepo) Create(product Product) error {
	return nil
}

func (r *pgProductRepo) GetAll() ([]*Product, error) {
	return []*Product{}, nil
}
