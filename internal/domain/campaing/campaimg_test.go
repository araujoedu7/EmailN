package campaing_test

import (
	"time"
	
	"emailn/internal/domain/campaing"
	"testing"


	"github.com/stretchr/testify/assert"

	
)

func Test_NewCampaing_CreateCampaing(t *testing.T){
	assert := assert.New(t)
	name := "Campaing x"
	content := "body"
	contacts := []string{"email1", "email2"}

	campaing := campaing.NewCampaing(name, content, contacts)

	
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

	
}

func Test_NewCampaing_IDIsNotNill(t *testing.T){
	assert := assert.New(t)
	name := "Campaing x"
	content := "body"
	contacts := []string{"email1", "email2"}

	campaing := campaing.NewCampaing(name, content, contacts)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnIsNotNill(t *testing.T){
	assert := assert.New(t)
	name := "Campaing x"
	content := "body"
	contacts := []string{"email1", "email2"}
	 now := time.Now().Add(-time.Minute)
	campaing := campaing.NewCampaing(name, content, contacts)
	assert.Greater(campaing.CreatedOn, now)
}
