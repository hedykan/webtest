package main

import (
	"strconv"

	"strings"

	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, everyone!\n可以试试在地址栏输入:\nhttp://localhost:1323/users/名字\nhttp://localhost:1323/sort/逗号数字字符串")
	})
	
	e.GET("/users/:id", getUser)
	e.GET("/sort/:str", getSort)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
  return c.String(http.StatusOK, "hello, "+id)
}

func getSort(c echo.Context) error {
	str_arr := strings.Split(c.Param("str"), ",")
	sr(str_arr[:])
	str := strings.Join(str_arr, ",")
	return c.String(http.StatusOK, str)
}

func sr(arr []string) {
	arrLen := len(arr)
	num_arr := make([]int, arrLen, arrLen)

	for i := 0; i < arrLen; i++ {
		num_arr[i], _ = strconv.Atoi(arr[i])
	}

	sort(num_arr[:], 0, (len(num_arr) - 1))

	for i := 0; i < arrLen; i++ {
		arr[i] = strconv.Itoa(num_arr[i])
	}
}
