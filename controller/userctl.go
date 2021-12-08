package controller

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func getAccounts(authPath string) gin.Accounts {
	_f, e := os.Open(authPath)
	if e != nil {
		fmt.Printf("Fail to find authFile %s\n", authPath)
		return nil
	}
	defer _f.Close()

	_br := bufio.NewReader(_f)
	_accounts := make(gin.Accounts, 3)
	for {
		_accPair, _, e := _br.ReadLine()
		if e == io.EOF {
			break
		}
		_accPairDict := strings.Split(string(_accPair), " ")
		if len(_accPairDict) != 2 {
			fmt.Printf("Format error meets when loading auths.\n")
			return nil
		}
		_accounts[_accPairDict[0]] = _accPairDict[1]
	}
	return _accounts
}

func userPremissionInterceotor(c *gin.Context) {
	_userID, _ := strconv.Atoi(c.Params.ByName("userid"))
	_userName := c.MustGet(gin.AuthUserKey).(string)
	_userVisited := service.UserService.GetUser(uint32(_userID))

	if _userVisited.Name != _userName {
		fmt.Printf("User %d has no permission visit %s's page.\n", _userID, _userName)
		return
	}
	c.Set("userInfo", _userVisited)
	c.Next()
}

func userHandler(c *gin.Context) {
	_t := template.Must(template.ParseFiles(HTMLPath + "User.html"))

	_userInfo, ok := c.Get("userInfo")
	if !ok {
		fmt.Println("Cann't extract userInfo from context.")
		return
	}
	_user := _userInfo.(*models.User)
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

	fmt.Println(_requestOrders...)

	_orderInfo := [...]interface{}{_userInfo, _requestOrders, _acceptOrders, _doneOrders}

	_t.Execute(c.Writer, _orderInfo)

}

func userActionsHandler(c *gin.Context) {
	fmt.Println("Action Handler!")
	fmt.Println(c.PostForm("cancelorderid"))

	if _orderid := c.PostForm("cancelorderid"); _orderid != "" {
		fmt.Println("cancel!")
		_orderid, _ := strconv.Atoi(_orderid)
		service.PendingOrderService.Update(uint32(_orderid), nil)
	} else if _orderid := c.PostForm("finishorderid"); _orderid != "" {
		fmt.Println("Done!")
		_orderid, _ := strconv.Atoi(_orderid)

		_oldOrderInfo, _ := service.PendingOrderService.Update(uint32(_orderid), nil)
		_olderOrder := _oldOrderInfo.(models.Order)
		service.UserService.Update(_olderOrder.AcceptorId, _olderOrder.Price)
		service.UserService.Update(_olderOrder.RequesterId, -_olderOrder.Price)
		service.DoneOrderService.Update(uint32(_orderid), _olderOrder)
	}

	c.Next()
}
