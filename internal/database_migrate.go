package internal

import (
	"github.com/jaeyo/personal-archive/common"
	"github.com/jaeyo/personal-archive/services"
	"github.com/pkg/errors"
	"gorm.io/gorm/migrator"
)

func migrate() error {
	savedVerInfo, err := services.GetAppService().GetSavedVerInfo()
	if err != nil {
		return errors.Wrap(err, "failed to get saved ver info")
	}

	savedVer, err := common.NewVersion(savedVerInfo)
	if err != nil {
		return errors.Wrap(err, "failed to parse saved version")
	}
	v0_9_0, _ := common.NewVersion("0.9.0")

	if savedVer.IsLessThan(v0_9_0) {
		if err := migrateArticleTag(); err != nil {
			return errors.Wrap(err, "failed to migrate article tag")
		}
		// TODO: remove old article_tag table
	}
	return nil
}

func migrateArticleTag() error {
	// TODO: migrate old article_tag
}
