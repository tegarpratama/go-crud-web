package entities

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	CreatedAt   string
	UpdatedAt   string
}
