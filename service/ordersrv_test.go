package service

import (
	"math/rand"
	"os"
	"sync"
	"testing"
	"theway2meal/models"
)

func Test_PendingOrderGet(t *testing.T) {
	clear()
	defer os.Remove(dbPath)
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, true, true)
	order2, e := PendingOrderService.GetPendingOrder(2)
	if e != nil || order2.OrderTime != models.Orders[2].OrderTime {
		t.Error("PendingOrderService Get test failed")
	}
}

func Test_DoneOrderGet(t *testing.T) {
	clear()
	defer os.Remove(dbPath)
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, true, true)
	order0, e := DoneOrderService.GetDoneOrder(0)
	if e != nil || order0.OrderTime != models.Orders[0].OrderTime {
		t.Error("DoneOrderService Get test failed")
	}
}

func Test_OrderPending(t *testing.T) {
	clear()
	defer os.Remove(dbPath)
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
		go appendOrder(test_case / N)
	}
	wg.Wait()

	if PendingOrderService.indexNext != test_case {
		t.Errorf("UID iota failed: indexNext is %d", PendingOrderService.indexNext)
	}
}

func Test_OrderDone(t *testing.T) {
	clear()
	defer os.Remove(dbPath)

	generateOrder := func(dealNums int, reqc chan models.Order) {
		for i := 0; i < dealNums; i++ {
			order := models.Orders[rand.Intn(len(models.Orders))].Detach().(models.Order)
			uid := PendingOrderService.GenerateUID()
			order.OrderID = uid

			reqc <- order
		}
		close(reqc)
	}

	appendOrder := func(reqc chan models.Order, donec chan int) {
		for order := range reqc {
			// pending
			PendingOrderService.Update(order.OrderID, order)
			donec <- order.OrderID

		}
		close(donec)
	}

	consumeOrder := func(donec chan int, flag chan bool) {
		for orderId := range donec {
			// consuming
			oldOrder, _ := PendingOrderService.Update(orderId, nil)
			DoneOrderService.Update(orderId, oldOrder)
		}
		flag <- true
	}
	reqC := make(chan models.Order, 100)
	doneC := make(chan int, 100)
	flag := make(chan bool)

	go generateOrder(test_case, reqC)
	go appendOrder(reqC, doneC)
	go consumeOrder(doneC, flag)
	<-flag

	if len(DoneOrderService.SelectAll()) != test_case ||
		len(PendingOrderService.SelectAll()) != 0 {
		t.Errorf("UID iota failed: Done is %d, Pending is %d",
			len(DoneOrderService.SelectAll()),
			len(PendingOrderService.SelectAll()))
	}
}
