package campaign

import (
	"emailn/internal/contract"
    "emailn/internal/imternal-errors"

	
)

type Service struct {
	Repository Repository
}



func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err

	}

	err = s.Repository.Save(campaign)
	if err != nil {
		
		return "",  InternalErrors.ErrInternal
	}
	return campaign.ID, nil
}
