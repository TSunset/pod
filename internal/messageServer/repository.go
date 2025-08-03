package messageserver

import "gorm.io/gorm"

//CRUD

type MessageRepository interface {
	CreateMessage(ms Message) (Message, error)
	GetAllMessages() ([]Message, error)
	GetMessageByID(id string) (Message, error)
	UpdateMessage(ms Message) error
	DeleteMessage(id string) error
}

type msRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &msRepository{db: db}
}

func (r *msRepository) CreateMessage(ms Message) (Message, error) {
	err := r.db.Create(&ms)
	return ms, err.Error
}

func (r *msRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *msRepository) GetMessageByID(id string) (Message, error) {
	var ms Message
	result := r.db.Where("id = ?", id).First(&ms)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return ms, nil
}

func (r *msRepository) UpdateMessage(ms Message) error {
	result := r.db.Save(&ms)
	return result.Error
}

func (r *msRepository) DeleteMessage(id string) error {
	result := r.db.Where("id = ?", id).Delete(&Message{})
	return result.Error
}
