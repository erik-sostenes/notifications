package domain

import "github.com/erik-sostenes/notifications-api/pkg/common"

// Address represents an entity related to the entity Order
type Address struct {
	AddressId           AddressId
	AddressCountry      AddressCountry
	AddressState        AddressState
	AddressMunicipality AddressMunicipality
	AddressLatitude     AddressLatitude
	AddressLongitude    AddressLongitude
}

func NewAddress(country, state, municipality string, latitude, longitude float64) (Address, error) {
	addressId, err := NewAddressId(common.GenerateUuID())
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
	value string
}

func NewAddressId(value string) (AddressId, error) {
	v, err := common.Identifier(value).Validate()

	return AddressId{v}, err
}

func (a AddressId) Value() string {
	return a.value
}

type AddressCountry struct {
	value string
}

func NewAddressCountry(value string) (AddressCountry, error) {
	v, err := common.String(value).Validate()

	return AddressCountry{v}, err
}

func (a AddressCountry) Value() string {
	return a.value
}

type AddressState struct {
	value string
}

func NewAddressState(value string) (AddressState, error) {
	v, err := common.String(value).Validate()

	return AddressState{v}, err
}

func (a AddressState) Value() string {
	return a.value
}

type AddressMunicipality struct {
	value string
}

func NewAddressMunicipality(value string) (AddressMunicipality, error) {
	v, err := common.String(value).Validate()

	return AddressMunicipality{v}, err
}

func (a AddressMunicipality) Value() string {
	return a.value
}

type AddressLatitude struct {
	value float64
}

func NewAddressLatitude(value float64) (AddressLatitude, error) {
	return AddressLatitude{value}, nil
}

func (a AddressLatitude) Value() float64 {
	return a.value
}

type AddressLongitude struct {
	value float64
}

func NewAddressLongitude(value float64) (AddressLongitude, error) {
	return AddressLongitude{value}, nil
}

func (a AddressLongitude) Value() float64 {
	return a.value
}
