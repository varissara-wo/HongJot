package service

type NewSpenderRequest struct {
	Name  string
	Email string
}

type SpenderResponse struct {
	ID    int
	Name  string
	Email string
}

type SpenderService interface {
	NewSpender(NewSpenderRequest) (*SpenderResponse, error)
	GetAllSpenders() ([]SpenderResponse, error)
}
