package repositories

import (
	"github.com/jaeyo/personal-archive/internal"
	"github.com/jaeyo/personal-archive/models"
	"sync"
)

type NoteRepository interface {
	Save(note *models.Note) error
	FindAllWithPage(offset, limit int) (models.Notes, int64, error)
	FindByIDsWithPage(ids []int64, offset, limit int) (models.Notes, int64, error)
	FindByIDs(ids []int64) (models.Notes, error)
	FindTitles() (models.Notes, error)
	GetByID(id int64) (*models.Note, error)
	ExistByTitle(title string) (bool, error)
	DeleteByIDs(ids []int64) error
}

type noteRepository struct {
	database             *internal.DB
	noteSearchRepository NoteSearchRepository
}

func newNoteRepository(app appIface) NoteRepository {
	return &noteRepository{}
}

var GetNoteRepository = func() func() NoteRepository {
	var instance NoteRepository
	var once sync.Once

	return func() NoteRepository {
		once.Do(func() {
			instance = &noteRepository{
				database:             internal.GetDatabase(),
				noteSearchRepository: GetNoteSearchRepository(),
			}
		})
		return instance
	}
}()

func (r *noteRepository) Save(note *models.Note) error {
	isInsert := note.ID == 0

	if err := r.database.Save(note).Error; err != nil {
		return err
	}

	if isInsert {
		return r.noteSearchRepository.Insert(note)
	} else {
		return r.noteSearchRepository.Update(note)
	}
}

func (r *noteRepository) FindAllWithPage(offset, limit int) (models.Notes, int64, error) {
	var notes []*models.Note
	if err := r.database.
		Preload("Paragraphs").
		Preload("Paragraphs.ReferenceArticles").
		Preload("Paragraphs.ReferenceWebs").
		Order("created DESC").
		Offset(offset).
		Limit(limit).
		Find(&notes).Error; err != nil {
		return nil, -1, err
	}

	var cnt int64
	if err := r.database.
		Model(&models.Note{}).
		Count(&cnt).Error; err != nil {
		return nil, -1, err
	}

	ensureNoteAssociationNotNil(notes)
	return notes, cnt, nil
}

func (r *noteRepository) FindByIDsWithPage(ids []int64, offset, limit int) (models.Notes, int64, error) {
	var notes []*models.Note
	if err := r.database.
		Preload("Paragraphs").
		Preload("Paragraphs.ReferenceArticles").
		Preload("Paragraphs.ReferenceWebs").
		Where("id IN ?", ids).
		Offset(offset).
		Limit(limit).
		Find(&notes).Error; err != nil {
		return nil, -1, err
	}

	ensureNoteAssociationNotNil(notes)
	return notes, int64(len(ids)), nil
}

func (r *noteRepository) FindByIDs(ids []int64) (models.Notes, error) {
	var notes []*models.Note
	if err := r.database.
		Preload("Paragraphs").
		Preload("Paragraphs.ReferenceArticles").
		Preload("Paragraphs.ReferenceWebs").
		Where("id IN ?", ids).
		Find(&notes).Error; err != nil {
		return nil, err
	}

	ensureNoteAssociationNotNil(notes)
	return notes, nil
}

func (r *noteRepository) FindTitles() (models.Notes, error) {
	var notes []*models.Note
	if err := r.database.
		Select("id", "title").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	ensureNoteAssociationNotNil(notes)
	return notes, nil
}

func (r *noteRepository) GetByID(id int64) (*models.Note, error) {
	var note models.Note
	err := r.database.
		Preload("Paragraphs").
		Preload("Paragraphs.ReferenceArticles").
		Preload("Paragraphs.ReferenceWebs").
		First(&note, id).Error

	ensureNoteAssociationNotNil(models.Notes{&note})
	return &note, err
}

func (r *noteRepository) ExistByTitle(title string) (bool, error) {
	var cnt int64
	err := r.database.
		Model(&models.Note{}).
		Where("title = ?", title).
		Count(&cnt).Error
	return cnt > 0, err
}

func (r *noteRepository) DeleteByIDs(ids []int64) error {
	if err := r.database.Where("id IN ?", ids).Delete(&models.Note{}).Error; err != nil {
		return err
	}

	return r.noteSearchRepository.Deletes(ids)
}

func ensureNoteAssociationNotNil(notes models.Notes) {
	for _, note := range notes {
		if note.Paragraphs == nil {
			note.Paragraphs = models.Paragraphs{}
		}
		for _, paragraph := range note.Paragraphs {
			if paragraph.ReferenceArticles == nil {
				paragraph.ReferenceArticles = models.ReferenceArticles{}
			}
			if paragraph.ReferenceWebs == nil {
				paragraph.ReferenceWebs = models.ReferenceWebs{}
			}
		}
	}
}
