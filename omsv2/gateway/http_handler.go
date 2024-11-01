package main

import (
	"errors"
	"github.com/showyquasar88/proj-combine/omsv2/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	pb "github.com/showyquasar88/proj-combine/omsv2/common/api"
)

type Handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{client}
}

func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items = make([]*pb.ItemsWithQuantity, 0)
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	// rpc 返回的错误带有一个错误码 所以可以更好的进行处理
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusInternalServerError, rStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, resp)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item must have an ID")
		}
		if item.Quantity <= 0 {
			return errors.New("item must have a positive Quantity")
		}
	}

	return nil
}
