package internal

type InternalModule interface {
	Database() *DB
}

type internalModule struct {
	database *DB
}

func NewInternalModule() InternalModule {
	return &internalModule{
		database: newDatabase(),
	}
}

func (m *internalModule) Database() *DB {
	return m.database
}
