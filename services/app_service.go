package services

import (
	"github.com/jaeyo/personal-archive/common"
	"github.com/jaeyo/personal-archive/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io/ioutil"
	"sync"
)

const (
	AppVer = "app.version"
)

type AppService interface {
	PreserveVerInfo() error
	GetSavedVerInfo() (string, error)
	GetCurrentVerInfo() (string, error)
}

type appService struct {
	miscRepository repositories.MiscRepository
}

var GetAppService = func() func() AppService {
	var instance AppService
	var once sync.Once

	return func() AppService {
		once.Do(func() {
			instance = &appService{
				miscRepository: repositories.GetMiscRepository(),
			}
		})
		return instance
	}
}()

func (s *appService) PreserveVerInfo() error {
	ver, err := s.GetCurrentVerInfo()
	if err != nil {
		return errors.Wrap(err, "failed to get current ver info")
	}

	if err := s.miscRepository.CreateOrUpdate(AppVer, ver); err != nil {
		return errors.Wrap(err, "failed to create / update")
	}
	return nil
}

func (s *appService) GetSavedVerInfo() (string, error) {
	verInfo, err := s.miscRepository.GetValue(AppVer)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "0.0.0", nil
		}
		return "", errors.Wrap(err, "failed to get ver info")
	}
	return verInfo, nil
}

func (s *appService) GetCurrentVerInfo() (string, error) {
	verFile := "/app/VERSION.txt"
	if common.IsLocal() {
		verFile = "./VERSION.txt"
	}

	ver, err := ioutil.ReadFile(verFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to read version.txt file")
	}
	return string(ver), nil
}
