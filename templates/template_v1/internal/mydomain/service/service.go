package service

import (
	"goproposal/internal/models"
	"goproposal/internal/mydomain"
	"goproposal/internal/mydomain/repository/mocked"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/san-services/apicore/apierror"
	"github.com/san-services/apicore/apisettings"
)

var (
	PersonNotFoundErr = apierror.New("Person not found", apierror.CodeNotFoundErr)
)

// Service is a struct able to access all data required
// to perform business logic functions
type Service struct {
	repo  mydomain.Repository
	cache *cache.Cache
}

// New constructs and returns a Service struct
func New(configCache apisettings.Cache, repo mydomain.Repository) (*Service, error) {
	var c *cache.Cache

	if configCache.Enabled {
		// Create a cache with a default expiration time of 5 minutes, and which
		// purges expired items every 10 minutes
		cacheTime := time.Minute * configCache.ExpirationMinutes
		purgeTime := time.Minute * configCache.PurgeMinutes

		c = cache.New(cacheTime, purgeTime)
	}

	s := &Service{
		repo:  repo,
		cache: c,
	}
	return s, nil
}

func (s Service) AddPerson(name string, age int32) error {
	return s.repo.SavePerson(name, age)
}

func (s Service) Persons() ([]models.Person, error) {
	return s.repo.GetAllPersons()
}

func (s Service) Person(id int64) (*models.Person, error) {
	p, err := s.repo.GetPerson(id)
	if err != nil {
		if err == mocked.NotFoundErr {
			return nil, PersonNotFoundErr
		}

		return nil, err
	}

	return p, nil
}
