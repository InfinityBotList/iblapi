package popltypes

import "github.com/infinitybotlist/eureka/dovewing/dovetypes"

type Partner struct {
	ID     string `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Image  string `json:"image" validate:"required"`
	Short  string `json:"short" validate:"required"`
	Links  []Link `json:"links" validate:"required,min=1,max=2"`
	UserID string `json:"-" validate:"required,numeric"`

	// Internal field
	User *dovetypes.PlatformUser `json:"user"`
}

type PartnerList struct {
	Featured        []*Partner `json:"featured" validate:"required,dive"`
	BotPartners     []*Partner `json:"bot_partners" validate:"required,dive"`
	BotListPartners []*Partner `json:"bot_list_partners" validate:"required,dive"`
}
