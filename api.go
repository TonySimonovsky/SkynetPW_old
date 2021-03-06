package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"fmt"
)

// api

// отправляет json-ответ с массивом текущих ботов
func listBot(c *echo.Context) error {
	//	TODO сформировать список ботов и отправить клиенту
	// временные данные
	list, err := MB.GetListBots()
	if err != nil {
		return c.JSON(http.StatusConflict, nil)
	}
	return c.JSON(http.StatusOK, list)
}

// createBot получает от веб-клиента данные и передает сформированную структуру менеджеру ботов
// для создания нового бота
func createBot(c *echo.Context) error {
	infbot := make(map[string]string)
	infbot["name"] = c.Query("name")
	infbot["server"] = c.Query("server")
	infbot["login"] = c.Query("login")
	infbot["password"] = c .Query("password")
	id, err := MB.AddBot(infbot)
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, id)
}

// sendActionToBot отправить команду боту
func sendActionToBot(c *echo.Context) error {
	param := make(map[string]interface{})
	id := c.Param("id")			// распарсить и получить id
	action := c.Param("action")	// и команду

	switch action {
	case "update":
		ProcessID, _ := strconv.Atoi(c.Query("ProcessID"))
		param["ProcessID"] = ProcessID
	case "connect":
		infbot := make(map[string]string)
		infbot["name"] = "bot1"
		id, _ = MB.AddBot(infbot)
	case "disconnect":
//		TODO отключить бота от игрового клиента
	}

	err := MB.SendActionToBot(id, action, param)	// отправить команду боту
	if err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, "ok\n")
}

// deleteBot получает команду от веб-клиента (удалить бота) и перенаправляет ее менеджеру ботов
func deleteBot(c *echo.Context) error {
	id := c.Param("id")
	action := "delete"
	MB.SendActionToBot(id, action, nil)

	return c.String(http.StatusOK, "ok\n")
}
