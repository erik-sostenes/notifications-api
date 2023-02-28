package domain

import (
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"
)

// Address represents an entity related to the entity Order
type Address struct {
	AddressId           AddressId
	AddressCountry      AddressCountry
	AddressState        AddressState
	AddressMunicipality AddressMunicipality
	AddressLatitude     AddressLatitude
	AddressLongitude    AddressLongitude
}

func NewAddress(country, state, municipality, latitude, longitude string) (Address, error) {
	addressId, err := NewAddressId(shared.GenerateUuID())
	if err != nil {
		return Address{}, err
	}

	addressCountry, err := NewAddressCountry(country)
	if err != nil {
		return Address{}, err
	}

	addressState, err := NewAddressState(state)
	if err != nil {
		return Address{}, err
	}

	addressMunicipality, err := NewAddressMunicipality(municipality)
	if err != nil {
		return Address{}, err
	}

	addressLatitude, err := NewAddressLatitude(latitude)
	if err != nil {
		return Address{}, err
	}

	addressLongitude, err := NewAddressLongitude(longitude)
	if err != nil {
		return Address{}, err
	}

	return Address{
		AddressId:           addressId,
		AddressCountry:      addressCountry,
		AddressState:        addressState,
		AddressMunicipality: addressMunicipality,
		AddressLatitude:     addressLatitude,
		AddressLongitude:    addressLongitude,
	}, err
}

type AddressId struct {
	Value string
}

func NewAddressId(value string) (AddressId, error) {
	v, err := domain.Identifier(value).Validate()

	return AddressId{v}, err
}

type AddressCountry struct {
	Value string
}

func NewAddressCountry(value string) (AddressCountry, error) {
	v, err := domain.String(value).Validate()

	return AddressCountry{v}, err
}

type AddressState struct {
	Value string
}

func NewAddressState(value string) (AddressState, error) {
	v, err := domain.String(value).Validate()

	return AddressState{v}, err
}

type AddressMunicipality struct {
	Value string
}

func NewAddressMunicipality(value string) (AddressMunicipality, error) {
	v, err := domain.String(value).Validate()

	return AddressMunicipality{v}, err
}

type AddressLatitude struct {
	Value float64
}

func NewAddressLatitude(value string) (AddressLatitude, error) {
	v, err := domain.Float(value).Validate()

	return AddressLatitude{v}, err
}

type AddressLongitude struct {
	Value float64
}

func NewAddressLongitude(value string) (AddressLongitude, error) {
	v, err := domain.Float(value).Validate()

	return AddressLongitude{v}, err
}
