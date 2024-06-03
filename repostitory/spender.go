package repository

type Spender struct {
	ID    int
	Name  string
	Email string
}

type SpenderRepository interface {
	Create(Spender) (*Spender, error)
	GetAll() ([]Spender, error)
}
