package controllers

import (
	"github.com/jaeyo/personal-archive/repositories"
	"github.com/jaeyo/personal-archive/services"
)

type ControllerModule interface {
	ArticleController() *ArticleController
	ArticleTagController() *ArticleTagController
	NoteController() *NoteController
	SettingController() *SettingController
}

type controllerModule struct {
	articleController    *ArticleController
	articleTagController *ArticleTagController
	noteController       *NoteController
	settingController    *SettingController
}

type appIface interface {
	services.ServiceModule
	repositories.RepositoryModule
}

func NewControllerModule(app appIface) ControllerModule {
	return &controllerModule{
		// TODO IMME: NewArticleController -> newArticleController
		articleController:    NewArticleController(app),
		articleTagController: NewArticleTagController(app),
		noteController:       NewNoteController(app),
		settingController:    NewSettingController(app),
	}
}

func (m *controllerModule) ArticleController() *ArticleController {
	return m.articleController
}

func (m *controllerModule) ArticleTagController() *ArticleTagController {
	return m.articleTagController
}

func (m *controllerModule) NoteController() *NoteController {
	return m.noteController
}

func (m *controllerModule) SettingController() *SettingController {
	return m.settingController
}
