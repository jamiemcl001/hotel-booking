package models

import (
	"database/sql"
	"time"
)

type Organisation struct {
	ID         int
	Name       string
	AddressID  int
	ArchivedAt time.Time
}

func (o *Organisation) Archived() bool {
	return o.ArchivedAt != time.Time{}
}

type OrganisationModel struct {
	db *sql.DB
}

func NewOrganisationModel(db *sql.DB) *OrganisationModel {
	om := new(OrganisationModel)

	om.db = db
	return om
}

type CreateOrganisationRequest struct{}

type CreateOrganisationResponse struct{}

func (om *OrganisationModel) CreateOrganisation(cor *CreateOrganisationRequest) (*CreateOrganisationResponse, error) {
	return nil, nil
}

type UpdateOrganisationRequest struct{}

type UpdateOrganisationResponse struct{}

func (om *OrganisationModel) UpdateOrganisation(uor *UpdateOrganisationRequest) (*UpdateOrganisationResponse, error) {
	return nil, nil
}

type ArchiveOrganisationRequest struct{}

type ArchiveOrganisationResponse struct{}

func (om *OrganisationModel) ArchiveOrganisation(aor *ArchiveOrganisationRequest) (*ArchiveOrganisationResponse, error) {
	return nil, nil
}
