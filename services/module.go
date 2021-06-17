package services

type ServiceModule interface {
	AppService() AppService
	ArticleService() ArticleService
	NoteService() NoteService
	PocketService() PocketService
	PocketSyncService() PocketSyncService
}

type serviceModule struct {
	appService        AppService
	articleService    ArticleService
	noteService       NoteService
	pocketService     PocketService
	pocketSyncService PocketSyncService
}

type appIface interface {
}

func NewServiceModule(app appIface) ServiceModule {
	return &serviceModule{
		appService:        newAppService(app),
		articleService:    newArticleService(app),
		noteService:       newNoteService(app),
		pocketService:     newPocketService(app),
		pocketSyncService: newPocketSyncService(app),
	}
}

func (m *serviceModule) AppService() AppService {
	return m.appService
}

func (m *serviceModule) ArticleService() ArticleService {
	return m.articleService
}

func (m *serviceModule) NoteService() NoteService {
	return m.noteService
}

func (m *serviceModule) PocketService() PocketService {
	return m.pocketService
}

func (m *serviceModule) PocketSyncService() PocketSyncService {
	return m.pocketSyncService
}
