package controller

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"theway2meal/models"
	"theway2meal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func orderPreviewHandler(c *gin.Context) {
	_t := template.Must(template.ParseFiles(HTMLPath + "Order.html"))

	_userName := c.MustGet(gin.AuthUserKey).(string)

	//get userID by name
	_userID := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		return _user.Name == _userName
	})[0].(models.User).UserID

	// try to get cookie
	_cookieMap := restoreCookies(c, []string{AUTHKEY})
	if len(_cookieMap) == 0 {
		fmt.Println("Authkey is not stored in cookie.")
		return
	}

	_mealId, e := strconv.Atoi(c.Params.ByName(MEALKEY))
	if e != nil {
		fmt.Printf("url error when precessing mealId.")
		return
	}
	_mealInfo := service.MealService.GetMeal(uint32(_mealId))
	_candiAccepter := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		_thisUserID, _ := strconv.Atoi(_cookieMap[AUTHKEY])
		return _user.UserID != uint32(_thisUserID)
	})

	_orderInfo := [...]interface{}{_userName, _mealInfo, _candiAccepter, ACCEPTERID}

	resetCookies(c, []string{
		REQUESTERNAME,
		MEALNAME,
		REQUESTERID,
		ORDERMEALID,
		MEALPRICE,
		MEALID,
	})
	setCookies(c, map[string]string{
		REQUESTERNAME: _userName,
		MEALNAME:      _mealInfo.Name,
		REQUESTERID:   strconv.Itoa(int(_userID)),
		ORDERMEALID:   c.Params.ByName(MEALKEY),
		MEALPRICE:     fmt.Sprintf("%.1f", _mealInfo.Price),
	})

	_t.Execute(c.Writer, _orderInfo)
}

func orderApplyHandler(c *gin.Context) {
	_t := template.Must(template.ParseFiles(HTMLPath + "Order_Pending.html"))

	_acceptInfo := strings.Split(c.PostForm(ACCEPTERID), ":")

	// extract order information
	_cookieMap := restoreCookies(c, []string{
		AUTHKEY,
		REQUESTERNAME,
		MEALNAME,
		REQUESTERID,
		ORDERMEALID,
		MEALPRICE,
	})

	_orderInfo := [...]interface{}{
		_cookieMap[REQUESTERNAME],
		_cookieMap[MEALNAME],
		_cookieMap[MEALPRICE],
		_acceptInfo[1],
		_cookieMap[REQUESTERID],
	}

	// Append Waiting order
	_selectID, _ := strconv.Atoi(_acceptInfo[0])
	_requestID, _ := strconv.Atoi(_cookieMap[REQUESTERID])
	_orderMealID, _ := strconv.Atoi(_cookieMap[ORDERMEALID])
	_orderUid := service.PendingOrderService.GenerateUID()
	_price, _ := strconv.ParseFloat(_cookieMap[MEALPRICE], 64)

	_orderPending := models.Order{
		OrderID:         _orderUid,
		OrderTime:       time.Now().Format(models.TimeFormat),
		RequesterName:   _cookieMap[REQUESTERNAME],
		AcceptorName:    _acceptInfo[1],
		OrderedMealName: _cookieMap[MEALNAME],
		Price:           _price,
		RequesterId:     uint32(_requestID),
		AcceptorId:      uint32(_selectID),
		OrderedMealId:   uint32(_orderMealID),
		IsReadyDelete:   false,
	}

	fmt.Println(_orderPending)
	service.PendingOrderService.Update(_orderUid, _orderPending)

	_t.Execute(c.Writer, _orderInfo)
}
