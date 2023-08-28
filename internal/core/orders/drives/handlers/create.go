package handlers

import (
	"net/http"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/services"
	"github.com/erik-sostenes/notifications-api/internal/core/orders/drives/handlers/dto"
	"github.com/erik-sostenes/notifications-api/pkg/bus/command"
	"github.com/erik-sostenes/notifications-api/pkg/server/response"
)

func HttpHandlerOrderCreator(cmd command.Bus[services.CreateOrderCommand]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.OrderRequest

		if ok, err := response.Bind(w, r, &request); err != nil || !ok {
			return
		}

		command := services.CreateOrderCommand{
			Id:       request.Id,
			CreateAt: request.CreateAt,
			Status:   request.Status,
			Price: services.Price{
				Amount:   request.PriceRequest.Amount,
				Currency: request.PriceRequest.Currency,
			},
			Address: services.Address{
				Country:      request.AddressRequest.Country,
				State:        request.AddressRequest.State,
				Municipality: request.AddressRequest.Municipality,
				Latitude:     request.AddressRequest.Latitude,
				Longitude:    request.AddressRequest.Longitude,
			},
			RequestedTime:    request.RequestedTime,
			IsProduct:        request.IsProduct,
			IsSubscription:   request.IsSubscription,
			TypeSubscription: request.TypeSubscription,
			UserId:           request.UserId,
		}

		err := cmd.Dispatch(r.Context(), command)
		if err != nil {
			_ = response.ErrorHandler(w, err)
			return
		}

		_ = response.JSON(w, http.StatusCreated, nil)
	}
}
