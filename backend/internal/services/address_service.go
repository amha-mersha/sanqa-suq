package services

import (
	"context"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type AddressService struct {
	addressRepo *repositories.AddressRepository
}

func NewAddressService(addressRepo *repositories.AddressRepository) *AddressService {
	return &AddressService{
		addressRepo: addressRepo,
	}
}

func (s *AddressService) CreateAddress(ctx context.Context, userID string, req *dtos.CreateAddressRequestDTO) (*dtos.AddressResponseDTO, error) {
	address := &models.Address{
		UserID:     userID,
		Street:     req.Street,
		City:       req.City,
		State:      req.State,
		PostalCode: req.PostalCode,
		Country:    req.Country,
		Type:       req.Type,
	}

	if err := s.addressRepo.CreateAddress(ctx, address); err != nil {
		return nil, err
	}

	return &dtos.AddressResponseDTO{
		AddressID:  address.AddressID,
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		Type:       address.Type,
	}, nil
}

func (s *AddressService) GetAddressByID(ctx context.Context, addressID int, userID string) (*dtos.AddressResponseDTO, error) {
	address, err := s.addressRepo.GetAddressByID(ctx, addressID, userID)
	if err != nil {
		return nil, err
	}

	return &dtos.AddressResponseDTO{
		AddressID:  address.AddressID,
		Street:     address.Street,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		Type:       address.Type,
	}, nil
}

func (s *AddressService) GetUserAddresses(ctx context.Context, userID string) ([]*dtos.AddressResponseDTO, error) {
	addresses, err := s.addressRepo.GetUserAddresses(ctx, userID)
	if err != nil {
		return nil, err
	}

	var response []*dtos.AddressResponseDTO
	for _, address := range addresses {
		response = append(response, &dtos.AddressResponseDTO{
			AddressID:  address.AddressID,
			Street:     address.Street,
			City:       address.City,
			State:      address.State,
			PostalCode: address.PostalCode,
			Country:    address.Country,
			Type:       address.Type,
		})
	}

	return response, nil
}

func (s *AddressService) UpdateAddress(ctx context.Context, addressID int, userID string, req *dtos.UpdateAddressRequestDTO) (*dtos.AddressResponseDTO, error) {
	existingAddress, err := s.addressRepo.GetAddressByID(ctx, addressID, userID)
	if err != nil {
		return nil, err
	}

	// Update only the fields that are provided
	if req.Street != "" {
		existingAddress.Street = req.Street
	}
	if req.City != "" {
		existingAddress.City = req.City
	}
	if req.State != "" {
		existingAddress.State = req.State
	}
	if req.PostalCode != "" {
		existingAddress.PostalCode = req.PostalCode
	}
	if req.Country != "" {
		existingAddress.Country = req.Country
	}
	if req.Type != "" {
		existingAddress.Type = req.Type
	}

	if err := s.addressRepo.UpdateAddress(ctx, existingAddress); err != nil {
		return nil, err
	}

	return &dtos.AddressResponseDTO{
		AddressID:  existingAddress.AddressID,
		Street:     existingAddress.Street,
		City:       existingAddress.City,
		State:      existingAddress.State,
		PostalCode: existingAddress.PostalCode,
		Country:    existingAddress.Country,
		Type:       existingAddress.Type,
	}, nil
}

func (s *AddressService) DeleteAddress(ctx context.Context, addressID int, userID string) error {
	return s.addressRepo.DeleteAddress(ctx, addressID, userID)
}
