package service

import (
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

// func Test_OrderAppend(t *testing.T) {
// 	clear()
// 	db := GetDefaultDB()
// 	db_loadTestingData(db)
// 	var wg sync.WaitGroup
// 	wg.Add(test_case)

// 	appendOrder := func(dealNums int) {
// 		for i := 0; i < dealNums; i++ {
// 			meal := models.Meals[rand.Intn(len(models.Meals))]
// 			UserService.Update(uint32(rand.Intn(len(models.Users))),
// 				meal.Price)
// 			UserService.Update(uint32(rand.Intn(len(models.Users))),
// 				-meal.Price)
// 			wg.Done()
// 		}
// 	}

// 	N := 4
// 	for i := 0; i < N; i++ {
// 		go appendOrder(int(test_case) / N)
// 	}
// 	wg.Wait()

// 	var sum float64
// 	for _, user := range UserService.SelectAll() {
// 		sum += user.(models.User).Balance
// 	}

// 	if math.Abs(sum) > 1e4 {
// 		t.Errorf("Bill checks failed: sum is %f", sum)
// 	}

// }
