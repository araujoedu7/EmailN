package campaign_test

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

type Service struct {
	Repository *repositoryMock
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	if newCampaign.Name == "" {
		return "", errors.New("name is required")
	}
	if newCampaign.Content == "" {
		return "", errors.New("content is required")
	}
	if len(newCampaign.Emails) == 0 {
		return "", errors.New("at least one email is required")
	}

	campaign := &campaign.Campaign{
		Name:    newCampaign.Name,
		Content: newCampaign.Content,
		Contacts: func(emails []string) []campaign.Contact {
			contacts := make([]campaign.Contact, len(emails))
			for i, email := range emails {
				contacts[i] = campaign.Contact{Email: email}
			}
			return contacts
		}(newCampaign.Emails),
	}

	if err := s.Repository.Save(campaign); err != nil {
		return "", err
	}
	return campaign.Name, nil
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:    "teste y",
		Content: "Body",
		Emails:  []string{"teste@.com"},
	}

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}

	repositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	
	newCampaign := contract.NewCampaign{
		Name:    "", 
		Content: "Body",
		Emails:  []string{"teste@.com"},
	}

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}

	_, err := service.Create(newCampaign)
	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}


func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:    "teste y",
		Content: "Body",
		Emails:  []string{"teste@teste.com"},
	}

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}

	
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	id, err := service.Create(newCampaign)

	
	assert.NotNil(id)
	assert.Nil(err)
	repositoryMock.AssertExpectations(t)
}
