package command

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/notifications-api/pkg/domain/wrongs"
)

type (
	// Command represents the intention to execute an operation on our system that modifies its state
	//
	// Command is a DTO (Data Transfer Object) which represents the action to be performed
	Command interface {
		// CommandId method that implements all commands(Data Transfer Object)
		// returns a string representing the type of command to be performed
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
		// Handler represents the action that you want to perform by means of the Command, which will send it to the service layer
		Handler(ctx context.Context, v *V) error
	}
)

// CommandBus is a map that implements the Bus interface and registers the Command with its Handler
type CommandBus[V Command] map[string]Handler[V]

// Record receives the Command and Handler and registers them
func (cb *CommandBus[V]) Record(c Command, h Handler[V]) (err error) {
	cmdID := c.CommandId()

	_, ok := (*cb)[cmdID]
	if ok {
		err = wrongs.CommandAlreadyRegisteredError(fmt.Sprintf("Command Already Registered %v", h))
		return
	}

	(*cb)[cmdID] = h

	return nil
}

// Dispatch receives the Command and calls the registered Handler
func (cb *CommandBus[V]) Dispatch(ctx context.Context, v V) (err error) {
	cmdID := v.CommandId()

	ch, ok := (*cb)[cmdID]
	if !ok {
		err = wrongs.CommandNotRegisteredError(fmt.Sprintf("Command Not Registered %v", v))
		return
	}

	return ch.Handler(ctx, &v)
}
