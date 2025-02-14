package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infestructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}
	
	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignGet)

	http.ListenAndServe(":3000", r)
}
