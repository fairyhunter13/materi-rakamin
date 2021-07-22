package designpattern

import (
	"fmt"
	"sync"
)

type restaurant struct {
	kitchen      *kitchen
	orderService *orderService
	delivery     *delivery
}

func newRestaurant() *restaurant {
	orderList := newOrderList()
	return &restaurant{
		orderService: newOrderService(orderList),
		kitchen:      newKitchen(),
		delivery:     newDelivery(orderList),
	}
}

func (r *restaurant) CustomerOrder(customerID int) (order order) {
	order = r.orderService.NewOrder(customerID)
	return
}

func (r *restaurant) HandleOrder(order order) {
	r.kitchen.Cook(order)
	r.delivery.Deliver(order)
}

type order struct {
	orderID uint64
	custID  int
}

type orderList struct {
	sync.RWMutex
	orders []order
}

func newOrderList() *orderList {
	return &orderList{}
}

func (ol *orderList) Add(order order) {
	ol.Lock()
	defer ol.Unlock()
	fmt.Printf("Adding order to the list, orderID: %d and customerID: %d.\n", order.orderID, order.custID)
	ol.orders = append(ol.orders, order)
}

func (ol *orderList) Remove(order order) {
	ol.Lock()
	defer ol.Unlock()
	fmt.Printf("Removing order from the list, orderID: %d and customerID: %d.\n", order.orderID, order.custID)
	for key := range ol.orders {
		if ol.orders[key].orderID != order.orderID || ol.orders[key].custID != order.custID {
			continue
		}
		ol.orders = append(ol.orders[:key], ol.orders[key+1:]...)
	}
}

type orderService struct {
	sync.RWMutex
	orderCounter uint64
	orderList    *orderList
}

func newOrderService(ol *orderList) *orderService {
	return &orderService{
		orderCounter: 1,
		orderList:    ol,
	}
}

func (os *orderService) NewOrder(customerID int) (res order) {
	os.Lock()
	defer os.Unlock()
	res = order{
		custID:  customerID,
		orderID: os.orderCounter,
	}
	os.orderList.Add(res)
	os.orderCounter++
	return
}

type kitchen struct{}

func newKitchen() *kitchen {
	return new(kitchen)
}

func (k *kitchen) Cook(order order) {
	fmt.Printf("Cook order #%d from the customer #%d.\n", order.orderID, order.custID)
}

type delivery struct {
	orderList *orderList
}

func newDelivery(ol *orderList) *delivery {
	return &delivery{
		orderList: ol,
	}
}

func (d *delivery) Deliver(order order) {
	d.orderList.Remove(order)
	fmt.Printf("Deliver order #%d for the customer #%d.\n", order.orderID, order.custID)
}

func Facade() {
	restaurant := newRestaurant()
	custID := 9
	newOrder := restaurant.CustomerOrder(custID)
	restaurant.HandleOrder(newOrder)
}
