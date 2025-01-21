package endpoints

import (
	"emailn/internal/contract"
	InternalErrors "emailn/internal/imternal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign

	// Decodifica o JSON da requisição
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid JSON: " + err.Error(),
		})
		return
	}

	// Validação dos campos do request
	if err := validate.Struct(request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Validation failed: " + err.Error(),
		})
		return
	}

	// Chama o serviço para criar a campanha
	id, err := h.CampaignService.Create(request)
	if err != nil {
		if errors.Is(err, InternalErrors.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusBadRequest)
		}
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Responde com sucesso e o ID da campanha
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"id": id,
	})
}
