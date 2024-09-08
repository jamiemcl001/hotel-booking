package models

import "database/sql"

type RoomModel struct {
	db *sql.DB
}

func NewRoomModel(db *sql.DB) *RoomModel {
	rm := new(RoomModel)

	rm.db = db
	return rm
}

type Room struct {
	ID              int
	RoomNumber      string
	HotelID         int
	RoomFeatures    string
	Capacity        int
	BaseNightlyRate float32
}

type CreateRoomRequest struct{}

type CreateRoomResponse struct{}

func (rm *RoomModel) CreateRoom(crr *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, nil
}

type ArchiveRoomRequest struct{}

type ArchiveRoomResponse struct{}

func (rm *RoomModel) ArchiveRoom(arr *ArchiveRoomRequest) (*ArchiveRoomResponse, error) {
	return nil, nil
}

type UpdateRoomRequest struct{}

type UpdateRoomResponse struct{}

func (rm *RoomModel) UpdateRoom(urr *UpdateRoomRequest) (*UpdateRoomResponse, error) {
	return nil, nil
}

type GetRoomsByHotelRequest struct{}

type GetRoomsByHotelResponse struct{}

func (rm *RoomModel) GetRoomsByHotel(grr *GetRoomsByHotelRequest) (*GetRoomsByHotelResponse, error) {
	return nil, nil
}

type FilterAvailableRoomsRequest struct{}

type FilterAvailableRoomsResponse struct{}

func (rm *RoomModel) FilterAvailableRooms(frr *FilterAvailableRoomsRequest) (*FilterAvailableRoomsResponse, error) {
	return nil, nil
}
