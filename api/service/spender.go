package service

type NewSpenderRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SpenderResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SpenderService interface {
	NewSpender(NewSpenderRequest) (*SpenderResponse, error)
	GetAllSpenders() ([]SpenderResponse, error)
}
