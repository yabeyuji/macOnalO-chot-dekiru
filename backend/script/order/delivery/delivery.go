package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "delivery/pb"
)

func main() {

	dial, err := grpc.Dial("localhost:3456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	defer dial.Close()

	conn := pb.NewDeliveryServiceClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var request = &pb.DeliveryRequest{
		Order: &pb.Order{
			Combos:    []*pb.Combo{},
			SideMenus: []*pb.SideMenu{},
			Drinks:    []*pb.Drink{},
			Hamburgers: []*pb.Hamburger{
				{
					Top:     1,
					Cheese:  1,
					Lettuce: 1,
					Tomato:  1,
				},
			},
		},
	}

	res, err := conn.DeliveryRPC(ctx, request)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Printf("%+v\n", res.String())
	return
}
