package controller

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
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

func changePw(username, pw string, r io.Reader, w io.Writer) error {
	// use scanner to read line by line
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		_accPairDict := strings.Split(line, " ")
		if _accPairDict[0] == username {
			_accPairDict[1] = pw
		}
		line = _accPairDict[0] + " " + _accPairDict[1] + "\n"
		if _, err := io.WriteString(w, line); err != nil {
			return err
		}
	}
	return sc.Err()
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
	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

	_userInfo, ok := c.Get("userInfo")

	if !ok {
		err = fmt.Errorf("can't extract userInfo from context")
	}

	_user, ok := _userInfo.(models.User)

	if !ok {
		err = fmt.Errorf("can't convert userInfo to model.User")
		return
	}

	c.HTML(http.StatusOK, "User.html", gin.H{
		"userid": _user.UserID,
	})

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

	_doneOrders := service.DoneOrderService.Select(0, func(i interface{}) bool {
		_order := i.(models.Order)
		return _user.UserID == _order.RequesterId || _user.UserID == _order.AcceptorId
	})

	// List in order
	sortOrders := func(orders []interface{}) {
		sort.Slice(orders, func(i, j int) bool {
			return orders[i].(models.Order).OrderID > orders[j].(models.Order).OrderID
		})

	}

	sortOrders(_requestOrders)
	sortOrders(_acceptOrders)
	sortOrders(_doneOrders)

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

		_, err = service.MealService.Update(_olderOrder.OrderedMealId)
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
}

func userPasswordChanging(c *gin.Context) {
	c.HTML(http.StatusOK, "PwChange.html", gin.H{
		"userid": c.Param("userid"),
	})
}

func onUserPasswordChanged(c *gin.Context) {
	var err error
	defer func() {
		if err != nil && err != io.EOF {
			log.Println(err.Error())
		}
	}()

	_newpw := c.PostForm("newpw")
	_confirm := c.PostForm("confirmpw")

	if _newpw == _confirm {
		_userName := c.MustGet(gin.AuthUserKey).(string)

		_f, err := os.Open(AUTHPATH)
		if err != nil {
			return
		}
		defer _f.Close()

		// create temp file
		tmp, err := ioutil.TempFile("data", "replace-*")
		if err != nil {
			return
		}
		defer tmp.Close()

		if err = changePw(_userName, _newpw, _f, tmp); err != nil {
			return
		}

		if err = tmp.Close(); err != nil {
			return
		}

		if err = _f.Close(); err != nil {
			return
		}

		// overwrite the original file with the temp file
		if err = os.Rename(tmp.Name(), AUTHPATH); err != nil {
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/user/"+c.Param("userid"))

	} else {
		c.Redirect(http.StatusMovedPermanently, "/user/"+c.Param("userid")+"/pwchange")
	}
}

func resetOrders(c *gin.Context) {
	c.HTML(http.StatusOK, "Reset.html", gin.H{})
}

func onOrderReseted(c *gin.Context) {
	service.OrdersReset()
	c.Redirect(http.StatusMovedPermanently, "/menu")
}
