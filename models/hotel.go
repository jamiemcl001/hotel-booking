package models

import (
	"database/sql"
	"time"
)

type Hotel struct {
	ID             int
	OrganisationID int
	AddressID      int
	StarRating     int
	ArchivedAt     time.Time
}

func (h *Hotel) Archived() bool {
	return h.ArchivedAt != time.Time{}
}

type HotelModel struct {
	db *sql.DB
}

func NewHotelModel(db *sql.DB) *HotelModel {
	hm := new(HotelModel)
	hm.db = db
	return hm
}

type CreateHotelRequest struct{}

type CreateHotelResponse struct{}

func (hm *HotelModel) NewHotel(chr *CreateHotelRequest) (*CreateHotelResponse, error) {
	return nil, nil
}

type UpdateHotelRequest struct{}

type UpdateHotelResponse struct{}

func (hm *HotelModel) UpdateHotel(uhr *UpdateHotelRequest) (*UpdateHotelResponse, error) {
	return nil, nil
}

type ArchiveHotelRequest struct{}

type ArchiveHotelResponse struct{}

func (hm *HotelModel) ArchiveHotel(ahr *ArchiveHotelRequest) (*ArchiveHotelResponse, error) {
	return nil, nil
}

type FilterHotelsRequest struct{}

type FilterHotelsResponse struct{}

func (hm *HotelModel) FilterHotels(fhr *FilterHotelsRequest) (*FilterHotelsResponse, error) {
	return nil, nil
}
