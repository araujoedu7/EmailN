package campaing_test

import (
	"time"

	"emailn/internal/domain/campaing"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing x"
	content  = "body"
	contacts = []string{"email1", "email2"}
)

func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := campaing.NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaing_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := campaing.NewCampaing(name, content, contacts)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaing, _ := campaing.NewCampaing(name, content, contacts)
	assert.Greater(campaing.CreatedOn, now)
}

func Test_NewCampaing_MustValidateName(t *testing.T) {
		assert := assert.New(t)
		
	_, err := campaing.NewCampaing("", content, contacts)		
		assert.Equal("name is required", err.Error())

}

func Test_NewCampaing_MustValidateContent(t *testing.T) {
	assert := assert.New(t)
	
_, err := campaing.NewCampaing(name, "", contacts)		
	assert.Equal("content is required", err.Error())

}

func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)
	
_, err := campaing.NewCampaing(name, content, []string{})		
	assert.Equal("contact is required", err.Error())

}