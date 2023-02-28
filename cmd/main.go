package main

import "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend/dependency"

func main() {
	if err := dependency.NewInjector(); err != nil {
		panic(err.Error())
	}
}
