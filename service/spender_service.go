package service

import repository "github.com/varissara-wo/hongjot/repostitory"

type spenderService struct {
	repo repository.SpenderRepository
}

func (ss spenderService) NewSpenderService(r repository.SpenderRepository) SpenderService {
	return spenderService{repo: r}
}

func (ss spenderService) NewSpender(r NewSpenderRequest) (*SpenderResponse, error) {
	s, err := ss.repo.Create(repository.Spender{Name: r.Name, Email: r.Email})
	if err != nil {
		return nil, err
	}
	sr := SpenderResponse{ID: s.ID, Name: s.Name, Email: s.Email}
	return &sr, nil
}

func (ss spenderService) GetAllSpenders() ([]SpenderResponse, error) {
	s, err := ss.repo.GetAll()
	if err != nil {
		return nil, err
	}
	sr := []SpenderResponse{}
	for _, v := range s {
		r := SpenderResponse{ID: v.ID, Name: v.Name, Email: v.Email}
		sr = append(sr, r)
	}
	return sr, nil
}
