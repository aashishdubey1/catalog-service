package product

type ProductService interface {
	GetProductById(id string) (*Product, error)
	CreateProduct(product Product) error
	GetAllProduct() ([]*Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product Product) error {
	return nil
}

func (s *productService) GetProductById(id string) (*Product, error) {
	return &Product{}, nil
}

func (s *productService) GetAllProduct() ([]*Product, error) {
	return []*Product{}, nil
}
