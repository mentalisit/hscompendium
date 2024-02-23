package server

import (
	"compendium/Compendium/generate"
	"compendium/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) CheckIdentityHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметра "code" из запроса
	code := r.URL.Query().Get("code")

	// Проверка наличия кода в запросе
	if code == "" || len(code) != 14 {
		http.Error(w, "Invalid code", http.StatusBadRequest)
		return
	}

	identity := generate.CheckCode(code)

	if identity.Token == "" {
		http.Error(w, "Outdated or invalid code", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(identity)
	go s.db.InsertIdentity(identity)
}

func (s *Server) SyncHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметров из запроса
	token := r.FormValue("token") //getByToken
	identity := s.GetTokenIdentity(token)
	if identity == nil {
		http.Error(w, fmt.Sprintf("Invalid token "), http.StatusBadRequest)
		return
	}
	mode := r.FormValue("mode")
	fmt.Println(mode)
	// Примечание: r.FormValue получает значение параметра из URL-запроса или тела запроса (если это POST-запрос)

	// Проверка корректности режима
	if mode != "get" && mode != "set" && mode != "sync" {
		http.Error(w, fmt.Sprintf("Invalid sync mode %s", mode), http.StatusBadRequest)
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
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	//нужно еще реализовать методы set,sync

	// Отправка запроса к другому API (заглушка для примера)
	// Замените этот блок на фактический код для обращения к другому API
	// Примечание: В этом примере просто выводим данные для примера
	fmt.Println(string(payload))

	// Отправка ответа клиенту
	fmt.Fprintf(w, "Sync data: %s", string(payload))
}
func (s *Server) GetTokenIdentity(token string) *models.Identity {
	for _, identity := range s.db.ReadIdentity() {
		if identity.Token == token {
			return &identity
		}
	}
	return nil
}
