package domain

type ProductCategory int

const (
	ElectronicCategory ProductCategory = iota
	FashionCategory
	HomeCategory
	HealthCategory
	SportCategory
	BooksCategory
)

type Product struct {
	ID       ID
	Name     string
	Description string
	Price    int64
	Category ProductCategory
	PhotoUrl string
}
