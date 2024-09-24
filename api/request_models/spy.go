package requestmodels

import "github.com/ZiplEix/pixel-espion/models"

type NewSpyRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=50"`
	Color string `json:"color" validate:"required,hexcolor"`
}

type GetAllSpiesResponse struct {
	Spies []models.Spy `json:"spies"`
}
