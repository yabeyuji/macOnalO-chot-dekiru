package delivery

import (
	"context"
	"log"
	"net"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"backend/internal/2_interface_adapter/controller"
	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	orderType = "delivery"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "delivery")
}

// Delivery ...
type Delivery struct {
	Server
}

// Server ...
type Server struct {
	UnimplementedDeliveryServiceServer
	Controller controller.ToController
}

// NewDelivery ...
func NewDelivery(ctrl controller.ToController) *Delivery {
	d := &Delivery{
		Server: Server{
			Controller: ctrl,
		},
	}

	return d
}

// Start ....
func (dlvr *Delivery) Start() {
	log.Println("start GRPC ------------------------- ")

	lis, err := net.Listen("tcp", pkg.DeliveryAddress)
	if err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterDeliveryServiceServer(s, &dlvr.Server)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		myErr.Logging(err)
		log.Fatalf("failed to serve: %v", err)
	}
}

// DeliveryRPC ...
func (srvr *Server) DeliveryRPC(ctx context.Context, in *DeliveryRequest) (*DeliveryResponse, error) {

	// web_uiのデータ型をControllerに持ち込まないようにproductに変換
	product := &entity.Product{}
	err := copier.Copy(product, in.Order)
	if err != nil {
		myErr.Logging(err)
		return nil, err
	}
	order := &entity.Order{Product: *product}

	srvr.Controller.Reserve(ctx, order, orderType)                          // オーダー番号発行
	srvr.Controller.Order(&ctx, order)                                      // オーダー
	return &DeliveryResponse{OrderNumber: order.OrderInfo.OrderNumber}, nil // オーダー番号返却
}
