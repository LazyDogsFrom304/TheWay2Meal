package service

import (
	"math/rand"
	"sync"
	"testing"
	"theway2meal/models"
)

func Test_PendingOrderGet(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	order2 := PendingOrderService.GetPendingOrder(2)
	if order2.OrderTime != models.Orders[2].OrderTime {
		t.Error("PendingOrderService Get test failed")
	}
}

func Test_DoneOrderGet(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	order0 := DoneOrderService.GetDoneOrder(0)
	if order0.OrderTime != models.Orders[0].OrderTime {
		t.Error("DoneOrderService Get test failed")
	}
}

func Test_OrderPending(t *testing.T) {
	clear()
	var wg sync.WaitGroup
	wg.Add(test_case)

	appendOrder := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			order := models.Orders[rand.Intn(len(models.Orders))].Detach().(models.Order)
			uid := PendingOrderService.GenerateUID()
			order.OrderID = uid
			PendingOrderService.Update(uid, order)
			wg.Done()
		}
	}

	N := 4
	for i := 0; i < N; i++ {
		go appendOrder(int(test_case) / N)
	}
	wg.Wait()

	if PendingOrderService.indexNext != test_case {
		t.Errorf("UID iota failed: indexNext is %d", PendingOrderService.indexNext)
	}
}

func Test_OrderDone(t *testing.T) {
	clear()

	generateOrder := func(dealNums int, reqc chan models.Order) {
		for i := 0; i < dealNums; i++ {
			order := models.Orders[rand.Intn(len(models.Orders))].Detach().(models.Order)
			uid := PendingOrderService.GenerateUID()
			// t.Log(uid)
			order.OrderID = uid

			reqc <- order
		}
		close(reqc)
	}

	appendOrder := func(reqc chan models.Order, donec chan uint32) {
		for order := range reqc {
			// pending

			PendingOrderService.Update(order.OrderID, order)
			donec <- order.OrderID
		}
		close(donec)
	}

	consumeOrder := func(donec chan uint32, flag chan bool) {
		for orderId := range donec {
			// consuming
			oldOrder, _ := PendingOrderService.Update(orderId, nil)
			// t.Log(oldOrder)
			DoneOrderService.Update(orderId, oldOrder)
		}
		flag <- true
	}
	reqC := make(chan models.Order, 100)
	doneC := make(chan uint32, 100)
	flag := make(chan bool)

	go generateOrder(test_case, reqC)
	go appendOrder(reqC, doneC)
	go consumeOrder(doneC, flag)
	<-flag

	t.Log(DoneOrderService.SelectAll())
	t.Log(PendingOrderService.SelectAll())
	if len(DoneOrderService.SelectAll()) != test_case ||
		len(PendingOrderService.SelectAll()) != 0 {
		t.Errorf("UID iota failed: Done is %d, Pending is %d",
			len(DoneOrderService.SelectAll()),
			len(PendingOrderService.SelectAll()))
	}
}
