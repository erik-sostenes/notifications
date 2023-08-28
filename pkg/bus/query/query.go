package query

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/notifications-api/pkg/domain/wrongs"
)

type (
	// Query represents the intention to request any data from our system without altering the state of our universe
	//
	// Query is a DTO(Data Transfer Object) able of representing the request you want to query
	Query interface {
		// QueryId method that implements all queries(Data Transfer Object)
		// returns a string representing the type of query to be performed
		QueryId() string
	}

	// Bus will be in charge of searching among the registered Query Handlers
	// and executing the function of such Handler when it receives a Query as parameter in its Handle method
	Bus[K Query, V any] interface {
		// Ask method that implements the QueryBus that looks for the registered Query with its Handler
		// and executes the function
		//
		// If the register no searching it is returns an error
		Ask(context.Context, K) (V, error)
	}

	// Handler is in charge of creating the Value Object with the primitive value
	// in order to validate that the query meets the requirements of our domain
	Handler[K Query, V any] interface {
		// Handler method will send the value of our domain, which will send it to the service layer
		Handler(context.Context, K) (V, error)
	}
)

// QueryBus is a map that implements the Bus interface and registers the Query with its Handler
type QueryBus[K Query, V any] map[string]Handler[K, V]

// Record receives the Query and Handler and registers them
func (qb *QueryBus[K, V]) Record(c Query, h Handler[K, V]) error {
	qryID := c.QueryId()

	if _, ok := (*qb)[qryID]; ok {
		return wrongs.QueryAlreadyRegisteredError(fmt.Sprintf("Query Already Registered %v", h))
	}

	(*qb)[qryID] = h

	return nil
}

// Ask receives the Query and calls the registered Handler
func (qb *QueryBus[K, V]) Ask(ctx context.Context, k K) (v V, err error) {
	qryID := k.QueryId()

	qh, ok := (*qb)[qryID]
	if !ok {
		return v, wrongs.QueryNotRegisteredError(fmt.Sprintf("Query Not Registered %v", v))
	}

	return qh.Handler(ctx, k)
}
