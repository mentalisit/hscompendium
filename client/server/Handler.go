package server

import (
	"compendium/Compendium/generate"
	"compendium/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) CheckIdentityHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization")
	code := c.GetHeader("authorization")
	fmt.Println("code", code)

	// Проверка наличия кода в запросе и его длины
	if code == "" || len(code) != 14 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}

	// Здесь должна быть ваша логика для генерации идентификации на основе полученного кода
	identity := generate.CheckCode(code)

	// Проверка на наличие токена в полученной идентификации
	if identity.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Outdated or invalid code"})
		return
	}

	// Отправка идентификации в формате JSON
	c.JSON(http.StatusOK, identity)
	cc = identity

	// Запуск асинхронной операции вставки идентификации в базу данных
	go s.db.InsertIdentity(identity)
}

var cc models.Identity

func (s *Server) CheckConnectHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization, content-type")
	//c.Header("Access-Control-Allow-Headers", "Authorization, no-cors, method, cache, headers")
	code := c.GetHeader("authorization")
	fmt.Println("authorization", code)
	c.JSON(http.StatusOK, cc)
}
func (s *Server) CheckSyncTechHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization, content-type")
	//c.Header("Access-Control-Allow-Headers", "Authorization, no-cors, method, cache, headers")
	code := c.GetHeader("authorization")
	fmt.Println("authorization", code)
	c.JSON(http.StatusOK, cc)
}

func (s *Server) SyncHandler(c *gin.Context) {
	// Получение параметра "token" из запроса
	token := c.Query("token")

	// Проверка наличия токена в запросе
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	// Получение параметра "mode" из запроса
	mode := c.Query("mode")

	fmt.Println(mode)
	// Проверка корректности режима
	if mode != "get" && mode != "set" && mode != "sync" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid sync mode %s", mode)})
		return
	}

	// Подготовка данных для отправки
	var currentTech map[int]models.TechLevel
	if mode == "get" {
		sd := s.db.ReadSyncData(token)
		if len(sd.TechLevels) > 0 {
			currentTech = sd.TechLevels
		}
	}

	data := map[string]interface{}{
		"ver":        1,
		"techLevels": currentTech,
	}

	// Преобразование данных в JSON
	payload, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	// Отправка ответа клиенту
	c.JSON(http.StatusOK, string(payload))
}
func (s *Server) GetTokenIdentity(token string) *models.Identity {
	for _, identity := range s.db.ReadIdentity() {
		if identity.Token == token {
			return &identity
		}
	}
	return nil
}

func (s *Server) CheckCorpDataHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Authorization")
	token := c.GetHeader("authorization")
	fmt.Println("token", token)
	i := s.db.ReadIdentityByToken(token)
	if i.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}
	var corp models.CorpData
	corp.Members = append(corp.Members, models.CorpMember{
		Name:         "mentalisit",
		UserId:       "1111111",
		ClientUserId: "222222",
	})
	corp.Roles = append(corp.Roles, models.CorpRole{
		Id:   "333",
		Name: "erevan",
	})
	c.JSON(http.StatusOK, corp)
	//if code == "" || len(code) != 14 {
	//
	//}
	//
	//// Здесь должна быть ваша логика для генерации идентификации на основе полученного кода
	//identity := generate.CheckCode(code)
	//
	//// Проверка на наличие токена в полученной идентификации
	//if identity.Token == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Outdated or invalid code"})
	//	return
	//}
	//
	//// Отправка идентификации в формате JSON
	//c.JSON(http.StatusOK, identity)
}
