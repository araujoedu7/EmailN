package campaign_test

import (
	"time"

	"emailn/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign x"
	content  = "body"
	contacts = []string{"email1", "email2"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaign, _ := campaign.NewCampaign(name, content, contacts)
	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign("", content, contacts)
	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, "", contacts)
	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, []string{})
	assert.Equal("contact is required", err.Error())
}
