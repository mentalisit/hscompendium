package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) CheckIdentityOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Status(http.StatusOK)
}
func (s *Server) CheckConnectOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	requestedHeaders := c.GetHeader("Access-Control-Request-Headers")
	//fmt.Println(requestedHeaders)
	if requestedHeaders != "" {
		c.Header("Access-Control-Allow-Headers", requestedHeaders)
	}

	c.Status(http.StatusOK)
}
func (s *Server) CheckSyncTechOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	requestedHeaders := c.GetHeader("Access-Control-Request-Headers")
	if requestedHeaders != "" {
		c.Header("Access-Control-Allow-Headers", requestedHeaders)
	}

	c.Status(http.StatusOK)
}

func (s *Server) Check(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// Получение запрашиваемого метода из заголовка Access-Control-Request-Method
	requestedMethod := c.GetHeader("Access-Control-Request-Method")

	// Получение запрашиваемых заголовков из заголовка Access-Control-Request-Headers
	requestedHeaders := c.GetHeader("Access-Control-Request-Headers")

	// Вывод полученных опций в консоль для отладки
	fmt.Println("Requested method:", requestedMethod)
	fmt.Println("Requested headers:", requestedHeaders)
	if requestedHeaders != "" {
		c.Header("Access-Control-Allow-Headers", requestedHeaders)
	}
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	// Отправка пустого ответа с кодом состояния 200 (OK)
	c.Status(http.StatusOK)
}
