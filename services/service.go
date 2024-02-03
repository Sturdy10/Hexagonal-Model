package services

import (
	"Hexagonal-Model/models"
	"Hexagonal-Model/repositories"
	"log"
)

type ServicesPort interface {
	PotsRegisterSer(register models.RequestRegister) error
}

type serviceAdapter struct {
	r repositories.RepositoryPort
}

func NewServices(r repositories.RepositoryPort) ServicesPort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) PotsRegisterSer(register models.RequestRegister) error {
	err := s.r.PotsRegister(register)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
