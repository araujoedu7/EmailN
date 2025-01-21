package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `json:"ID" validate:"required"`
	Name      string    `json:"Name" validate:"min=5,max=24"`
	CreatedOn time.Time `json:"CreatedOn" validate:"required"`
	Content   string    `json:"Content" validate:"min=5,max=1094"`
	Contacts  []Contact `json:"Contacts" validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	// Validação dos parâmetros de entrada
	if len(name) < 5 || len(name) > 24 {
		return nil, errors.New("name must be between 5 and 24 characters")
	}
	if len(content) < 5 || len(content) > 1094 {
		return nil, errors.New("content must be between 5 and 1094 characters")
	}
	if len(emails) == 0 {
		return nil, errors.New("at least one contact is required")
	}

	// Construção do slice de contatos
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	// Criação da campanha
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}, nil
}
