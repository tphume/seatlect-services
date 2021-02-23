package ordermb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"github.com/tphume/seatlect-services/internal/genproto/orderpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	repo Repo

	orderpb.UnimplementedOrderServiceServer
}

func (s *Server) ListOrder(ctx context.Context, req *orderpb.ListOrderRequest) (*orderpb.ListOrderResponse, error) {
	id := ctx.Value("id").(string)
	if len(id) <= 0 {
		return nil, status.Error(codes.Unauthenticated, "ID is not valid")
	}

	orders, err := s.repo.ListOrderByCustomer(ctx, id, req.Limit, req.Page)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	res := typedbToCommonpb(orders)

	return &orderpb.ListOrderResponse{Orders: res}, nil
}

type Repo interface {
	ListOrderByCustomer(ctx context.Context, CustomerId string, limit int32, page int32) ([]typedb.Order, error)
}

// Helper function
func typedbToCommonpb(orders []typedb.Order) []*commonpb.Order {
	res := make([]*commonpb.Order, len(orders))
	for i, o := range orders {
		res[i] = &commonpb.Order{
			XId:           o.Id.Hex(),
			ReservationId: o.ReservationId.Hex(),
			Business:      o.BusinessId.Hex(),
			Start:         o.Start.String(),
			End:           o.End.String(),
			Seats:         seatsToCommonpb(o.Seats),
			Status:        o.Status,
			Image:         o.Image,
			ExtraSpace:    int32(o.ExtraSpace),
			Name:          o.Name,
		}
	}

	return res
}

func seatsToCommonpb(seats []typedb.Seat) []*commonpb.OrderSeat {
	res := make([]*commonpb.OrderSeat, len(seats))
	for i, s := range seats {
		res[i] = &commonpb.OrderSeat{
			Name:  s.Name,
			Space: int32(s.Space),
		}
	}

	return res
}
