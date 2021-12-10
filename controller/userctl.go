package controller

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func getAccounts(authPath string) gin.Accounts {
	var err error
	defer func() {
		if err != nil && err != io.EOF {
			log.Println(err.Error())
		}
	}()

	_f, err := os.Open(authPath)
	if err != nil {
		return nil
	}
	defer _f.Close()

	_br := bufio.NewReader(_f)
	_accounts := make(gin.Accounts, 3)
	for {
		_accPair, _, err := _br.ReadLine()
		if err == io.EOF {
			break
		}
		_accPairDict := strings.Split(string(_accPair), " ")
		if len(_accPairDict) != 2 {
			log.Panic("format error meets when loading auths")
		}
		_accounts[_accPairDict[0]] = _accPairDict[1]
	}
	return _accounts
}

func userPremissionInterceotor(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

	_userID, err := strconv.Atoi(c.Params.ByName("userid"))
	if err != nil {
		return
	}

	_userName := c.MustGet(gin.AuthUserKey).(string)
	_userVisited, err := service.UserService.GetUser(_userID)
	if err != nil {
		return
	}

	if _userVisited.Name != _userName {
		err = fmt.Errorf("user %s/id:%d has no permission visit %s's page", _userVisited.Name, _userID, _userName)
		return
	}
	c.Set("userInfo", _userVisited)
	c.Next()
}

func userHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "User.html", gin.H{})

}

func userOrderPresentor(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

	_userInfo, ok := c.Get("userInfo")
	if !ok {
		err = fmt.Errorf("can't extract userInfo from context")
		return
	}

	_user, ok := _userInfo.(models.User)
	if !ok {
		err = fmt.Errorf("can't convert userInfo to model.User")
		return
	}

	_requestOrders := service.PendingOrderService.Select(10, func(i interface{}) bool {
		_order := i.(models.Order)
		return _user.UserID == _order.RequesterId
	})

	_acceptOrders := service.PendingOrderService.Select(10, func(i interface{}) bool {
		_order := i.(models.Order)
		return _user.UserID == _order.AcceptorId
	})

	_doneOrders := service.DoneOrderService.Select(10, func(i interface{}) bool {
		_order := i.(models.Order)
		return _user.UserID == _order.RequesterId || _user.UserID == _order.AcceptorId
	})

	_orderInfo := [...]interface{}{_userInfo, _requestOrders, _acceptOrders, _doneOrders}
	c.IndentedJSON(http.StatusOK, _orderInfo)
}

func userActionsHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

	log.Println("User Action Handler!")

	// TODO: switch?
	if _orderid := c.PostForm("cancelorderid"); _orderid != "" {
		log.Println("Action Name: Cancel")

		_orderid, err := strconv.Atoi(_orderid)
		if err != nil {
			return
		}

		_olderOrder, err := service.PendingOrderService.GetPendingOrder(_orderid)
		if err != nil {
			return
		}

		_olderOrder.IsReadyDelete = true
		service.PendingOrderService.Update(_orderid, _olderOrder)
	} else if _orderid := c.PostForm("finishorderid"); _orderid != "" {
		log.Println("Action Name: Done!")

		_orderid, err := strconv.Atoi(_orderid)
		if err != nil {
			return
		}

		_oldOrderInfo, err := service.PendingOrderService.Update(_orderid, nil)
		if err != nil {
			return
		}

		_olderOrder := _oldOrderInfo.(models.Order)
		// TODO: recover from a middle way operation
		_, err = service.UserService.Update(_olderOrder.AcceptorId, _olderOrder.Price)
		if err != nil {
			return
		}

		_, err = service.UserService.Update(_olderOrder.RequesterId, -_olderOrder.Price)
		if err != nil {
			return
		}

		_, err = service.DoneOrderService.Update(_orderid, _olderOrder)
		if err != nil {
			return
		}
	} else if _orderid := c.PostForm("confirmcancel"); _orderid != "" {
		log.Println("Action Name: Confirm Canceling")

		_orderid, err := strconv.Atoi(_orderid)
		if err != nil {
			return
		}

		_olderOrder, err := service.PendingOrderService.GetPendingOrder(_orderid)
		if err != nil {
			return
		}

		if _olderOrder.IsReadyDelete {
			_, err = service.PendingOrderService.Update(_orderid, nil)
			if err != nil {
				return
			}
		}
	}

	c.Next()
}
