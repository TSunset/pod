package messageserver

import "github.com/google/uuid"

type MessageService interface {
	CreateMessage(info string) (Message, error)
	GetAllMessages() ([]Message, error)
	GetMessageByID(id string) (Message, error)
	UpdateMessage(id string, info string) error
	DeleteMessage(id string) error
}

type msService struct {
	repo MessageRepository
}

// CreateMessage implements MessageService.
func (s *msService) CreateMessage(info string) (Message, error) {
	ms := Message{
		Text: info,
		ID:   uuid.New().String(),
	}

	createdMsg, err := s.repo.CreateMessage(ms)
	if err != nil {
		return Message{}, err
	}

	return createdMsg, nil
}

// DeleteMessage implements MessageService.
func (s *msService) DeleteMessage(id string) error {
	// Проверяем существование сообщения
	if _, err := s.repo.GetMessageByID(id); err != nil {
		return err
	}
	return s.repo.DeleteMessage(id)
}

// GetAllMessages implements MessageService.
func (s *msService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

// GetMessageByID implements MessageService.
func (s *msService) GetMessageByID(id string) (Message, error) {
	return s.repo.GetMessageByID(id)
}

// UpdateMessage implements MessageService.
func (s *msService) UpdateMessage(id string, info string) error {
	// Получаем сообщение по ID
	ms, err := s.repo.GetMessageByID(id)
	if err != nil {
		return err
	}

	// Обновляем текст
	ms.Text = info

	// Обновляем сообщение
	return s.repo.UpdateMessage(ms)
}

func NewMessageService(r MessageRepository) MessageService {
	return &msService{repo: r}
}
