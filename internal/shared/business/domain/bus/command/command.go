package command

import (
	"fmt"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/errors"

	"golang.org/x/net/context"
)

type (
	// Command represents the intention to execute an operation on our system that modifies its state
	//
	// Command is a DTO (Data Transfer Object) which represents the action to be performed
	Command interface {
		// CommandId method that implements all commands(Data Transfer Object)
		CommandId() string
	}

	// Bus will be in charge of searching among the registered Command Handlers
	// and executing the function of such Handler when it receives a Command as parameter in its Handle method
	Bus[V Command] interface {
		// Dispatch method that implements the CommandBus that looks for the registered Command with its Handler
		// and executes the function
		//
		// If the register no searching it is returns an error
		Dispatch(ctx context.Context, v V) error
	}

	// Handler will be in charge of performing the action we are looking for,
	// simply returning an error if the process is not correct
	//
	// Maps the DTO values to value objects in our domain and invokes the use case
	Handler[V Command] interface {
		// Handler represents the action that you want to perform by means of the Command
		Handler(context.Context, V) error
	}
)

// CommandBus is a map that implements the Bus interface and registers the Command with its Handler.
type CommandBus[V Command] map[string]Handler[V]

// RegisterHandler receives the Command and Handler and registers them
func (cb *CommandBus[V]) RegisterHandler(c Command, v Handler[V]) error {
	cmdId := c.CommandId()

	_, ok := (*cb)[cmdId]
	if ok {
		return errors.CommandAlreadyRegisteredError(fmt.Sprintf("Command Already Registered %v", v))
	}

	(*cb)[cmdId] = v

	return nil
}

// Dispatch receives the Command and calls the registered Handler
func (cb *CommandBus[V]) Dispatch(ctx context.Context, v V) (err error) {
	cmdID := v.CommandId()

	ch, ok := (*cb)[cmdID]
	if !ok {
		return errors.CommandNotRegisteredError(fmt.Sprintf("Command Not Registered %v", v))
	}

	return ch.Handler(ctx, v)
}
