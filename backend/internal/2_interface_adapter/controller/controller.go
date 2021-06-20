package controller

import (
	"context"
	"fmt"
	"time"

	"backend/internal/2_interface_adapter/gateway"
	"backend/internal/2_interface_adapter/presenter"
	"backend/internal/3_application_business_rule/usecase"
	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "controller")
}

type (
	// controller ...
	controller struct {
		UseCase     usecase.ToUseCase
		OrderNumber int
	}

	// orderChannel ...
	orderChannel struct {
		ctx   *context.Context
		order *entity.Order
	}

	// ToController ...
	ToController interface {
		Start()
		Reserve(ctx context.Context, order *entity.Order, orderType string)
		Order(ctx *context.Context, order *entity.Order)
	}
)

// orderChannel ...
var odrChnnl = make(chan orderChannel)

// NewController ...
func NewController(
	toRefrigerator gateway.ToRefrigerator,
	toFreezer gateway.ToFreezer,
	toShelf gateway.ToShelf,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	toEntity := entity.NewEntity()
	toGateway := gateway.NewGateway(
		toRefrigerator,
		toFreezer,
		toShelf,
	)
	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)
	uscs := usecase.NewUseCase(
		toEntity,
		toGateway,
		toPresenter,
	)
	ct := &controller{
		UseCase: uscs,
	}

	return ct
}

func (ctrl *controller) Start() {
	go ctrl.bulkOrder()
	go ctrl.UseCase.Start()
}

// Reserve ...
func (ctrl *controller) Reserve(ctx context.Context, order *entity.Order, orderType string) {
	ctrl.OrderNumber++ // オーダー番号発行
	if ctrl.OrderNumber >= 1000 {
		ctrl.OrderNumber = 1 // オーダー番号を3桁以内にする
	}

	order.OrderInfo.OrderNumber = fmt.Sprintf("%03d", ctrl.OrderNumber) // オーダー番号を3桁で表示
	order.OrderInfo.OrderType = orderType                               // オーダーの種類(mobile/pc/delivery/register)
	order.OrderInfo.OrderTime = time.Now()                              // オーダー受け付け時間

	ctrl.UseCase.Reserve(ctx, &order.OrderInfo) // オーダー情報更新
}

// Order ...
func (ctrl *controller) Order(ctx *context.Context, order *entity.Order) {
	oc := &orderChannel{
		ctx:   ctx,
		order: order,
	}

	// オーダー番号をweb_uiに即時返却する必要があるため、
	// 後続処理をチャネル経由で処理
	odrChnnl <- *oc
}

func (ctrl *controller) bulkOrder() {
	// 無限ループでチャネルを待ち受け
	for {
		oc := <-odrChnnl // Orderメソッドのチャネルの受け取り
		err := ctrl.UseCase.Order(oc.ctx, oc.order)
		if err != nil {
			myErr.Logging(err)
		}
	}
}
