package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"theway2meal/models"

	"github.com/gin-gonic/gin"
)

func setCookies(c *gin.Context, kvs map[string]string) {
	for k, v := range kvs {
		http.SetCookie(
			c.Writer,
			&http.Cookie{
				Name:  k,
				Value: url.QueryEscape(v),
				Path:  "/",
			})
	}
}

func restoreCookies(c *gin.Context, keys []string) map[string]string {
	restored := make(map[string]string)
	for _, k := range keys {

		_authCookie, e := c.Cookie(k)
		if e != nil {
			fmt.Printf("%s is not stored in cookie\n", k)
			continue
		}
		restored[k] = _authCookie
		fmt.Printf("cookie key %s restored %s\n", k, _authCookie)
	}
	return restored
}

func resetCookies(c *gin.Context, keys []string) {
	for _, k := range keys {
		http.SetCookie(
			c.Writer,
			&http.Cookie{
				Name:   k,
				Value:  "",
				MaxAge: -1})
	}
}

func floorSelector(floor int) func(interface{}) bool {
	return func(i interface{}) bool {
		_meal := i.(models.Meal)
		return _meal.Floor == floor
	}
}
