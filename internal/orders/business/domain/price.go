package domain

import "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"

// Price represents an entity related to the entity Order
type Price struct {
	PriceAmount   PriceAmount
	PriceCurrency PriceCurrency
}

func NewPrice(amount, currency string) (Price, error) {
	priceAmount, err := NewPriceAmount(amount)
	if err != nil {
		return Price{}, err
	}

	priceCurrency := PriceCurrency(currency)
	if err != nil {
		return Price{}, err
	}

	return Price{
		PriceAmount:   priceAmount,
		PriceCurrency: priceCurrency,
	}, err
}

type PriceAmount struct {
	value float64
}

func NewPriceAmount(value string) (PriceAmount, error) {
	v, err := domain.Float(value).Validate()

	return PriceAmount{v}, err
}

func (a PriceAmount) Value() float64 {
	return a.value
}

type PriceCurrency string

const (
	MX PriceCurrency = "MX"
)
