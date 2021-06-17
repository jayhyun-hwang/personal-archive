package repositories

type RepositoryModule interface {
	ArticleRepository() ArticleRepository
	ArticleSearchRepository() ArticleSearchRepository
	ArticleTagRepository() ArticleTagRepository
	MiscRepository() MiscRepository
	NoteRepository() NoteRepository
	NoteSearchRepository() NoteSearchRepository
	ParagraphRepository() ParagraphRepository
	ReferenceArticleRepository() ReferenceArticleRepository
	ReferenceWebRepository() ReferenceWebRepository
}

type repositoryModule struct {
	articleRepository          ArticleRepository
	articleSearchRepository    ArticleSearchRepository
	articleTagRepository       ArticleTagRepository
	miscRepository             MiscRepository
	noteRepository             NoteRepository
	noteSearchRepository       NoteSearchRepository
	paragraphRepository        ParagraphRepository
	referenceArticleRepository ReferenceArticleRepository
	referenceWebRepository     ReferenceWebRepository
}

type appIface interface {
}

func NewRepositoryModule(app appIface) RepositoryModule {
	return &repositoryModule{
		articleRepository:          newArticleRepository(app),
		articleSearchRepository:    newArticleSearchRepository(app),
		articleTagRepository:       newArticleTagRepository(app),
		miscRepository:             newMiscRepository(app),
		noteRepository:             newNoteRepository(app),
		noteSearchRepository:       newNoteSearchRepository(app),
		paragraphRepository:        newParagraphRepository(app),
		referenceArticleRepository: newReferenceArticleRepository(app),
		referenceWebRepository:     newReferenceWebRepository(app),
	}
}

func (m *repositoryModule) ArticleRepository() ArticleRepository {
	return m.articleRepository
}

func (m *repositoryModule) ArticleSearchRepository() ArticleSearchRepository {
	return m.articleSearchRepository
}

func (m *repositoryModule) ArticleTagRepository() ArticleTagRepository {
	return m.articleTagRepository
}

func (m *repositoryModule) MiscRepository() MiscRepository {
	return m.miscRepository
}

func (m *repositoryModule) NoteRepository() NoteRepository {
	return m.noteRepository
}

func (m *repositoryModule) NoteSearchRepository() NoteSearchRepository {
	return m.noteSearchRepository
}

func (m *repositoryModule) ParagraphRepository() ParagraphRepository {
	return m.paragraphRepository
}

func (m *repositoryModule) ReferenceArticleRepository() ReferenceArticleRepository {
	return m.referenceArticleRepository
}

func (m *repositoryModule) ReferenceWebRepository() ReferenceWebRepository {
	return m.referenceWebRepository
}
