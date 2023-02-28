package domain

import "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"

// CreatePriceRequestMother returns instances of domain.Price
type CreatePriceRequestMother struct{}

// createPriceRequestMother create a new instance of domain.Price
func createPriceRequestMother(amount, currency string) (domain.Price, error) {
	return domain.NewPrice(amount, currency)
}

// Random generate an instance of domain.Price by random values
func (p CreatePriceRequestMother) Random() (domain.Price, error) {
	return createPriceRequestMother("34.45", "MX")
}

// CreateAddressRequestMother returns instances of domain.Address
type CreateAddressRequestMother struct{}

// createPriceRequestMother create a new instance of domain.Address
func createAddressRequestMother(country, state, municipality, latitude, longitude string) (domain.Address, error) {
	return domain.NewAddress(country, state, municipality, latitude, longitude)
}

// Random generate an instance of domain.Address by random values
func (CreateAddressRequestMother) Random() (domain.Address, error) {
	return createAddressRequestMother("Mexico", "HIDALGO", "Tula de Allende Hidalgo", "6.5568768", "3.3488896")
}

// CreateOrderRequestMother returns instances of domain.Order
type CreateOrderRequestMother struct{}

// createOrderRequestMother create a new instance of domain.Order
func createOrderRequestMother(id string, createAt string, status string, price domain.Price, address domain.Address,
	requestedTime string, isProduct string, isSubscription string, typeSubscription string, userId string, foodDishesId []string) (domain.Order, error) {

	return domain.NewOrder(id, createAt, status, price, address, requestedTime, isProduct, isSubscription,
		typeSubscription, userId, foodDishesId)
}

// RandomOrderOne create a new instance of domain.Order
func (CreateOrderRequestMother) RandomOrderOne() (domain.Order, error) {
	address, err := CreateAddressRequestMother{}.Random()
	if err != nil {
		return domain.Order{}, err
	}

	price, err := CreatePriceRequestMother{}.Random()
	if err != nil {
		return domain.Order{}, err
	}

	return createOrderRequestMother("1e737f50-07f1-4d1b-9c3a-62f4d38523a7", "2022-11-21 19:51:39", "WAITING", price, address,
		"2022-11-21 19:51:39", "true", "false", "YEAR", "c2f91217-de8b-46fa-9168-132fe9285d87",
		[]string{"d3527262-f415-41c8-9aee-38812af4e484", "d3527262-f415-41c8-9aee-38812af4e484"})
}

// RandomOrderTwo create a new instance of domain.Order
func (CreateOrderRequestMother) RandomOrderTwo() (domain.Order, error) {
	address, err := CreateAddressRequestMother{}.Random()
	if err != nil {
		return domain.Order{}, err
	}

	price, err := CreatePriceRequestMother{}.Random()
	if err != nil {
		return domain.Order{}, err
	}

	return createOrderRequestMother("1e737f50-07f1-4d1b-9c3a-62f4d38559a9", "2022-11-21 19:51:39", "WAITING", price, address,
		"2022-11-21 19:51:39", "true", "false", "YEAR", "c2f91217-de8b-46fa-9168-132fe9285d87",
		[]string{"d3527262-f415-41c8-9aee-38812af4e484", "d3527262-f415-41c8-9aee-38812af4e484"})
}
