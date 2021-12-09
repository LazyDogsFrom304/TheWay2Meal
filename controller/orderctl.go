package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"theway2meal/models"
	"theway2meal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func orderPreviewHandler(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err.Error())
		}
	}()

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
		log.Println("Authkey is not stored in cookie.")
		return
	}

	_mealId, err := strconv.Atoi(c.Params.ByName(MEALKEY))
	if err != nil {
		return
	}

	_mealInfo, err := service.MealService.GetMeal(_mealId)
	if err != nil {
		return
	}

	_candiAccepter := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		_thisUserID, _ := strconv.Atoi(_cookieMap[AUTHKEY])
		return _user.UserID != _thisUserID
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
		REQUESTERID:   strconv.Itoa(_userID),
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
		RequesterId:     _requestID,
		AcceptorId:      _selectID,
		OrderedMealId:   _orderMealID,
		IsReadyDelete:   false,
	}

	log.Println("Apply Order INFO: ", _orderPending)
	_, err := service.PendingOrderService.Update(_orderUid, _orderPending)
	if err != nil {
		log.Println(err)
		return
	}

	_t.Execute(c.Writer, _orderInfo)
}

func checkOrderStatus(c *gin.Context) {
	_id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		log.Println(e)
		return
	}
	_order, e := service.PendingOrderService.GetPendingOrder(_id)
	if e != nil {
		log.Println(e)
		return
	}

	ans := strconv.FormatBool(_order.IsReadyDelete)
	c.String(http.StatusOK, ans)

}
