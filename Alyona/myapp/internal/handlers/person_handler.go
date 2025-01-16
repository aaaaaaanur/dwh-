package handlers

import (
	"awesomeProject/Alyona/myapp/internal/models"
	"awesomeProject/Alyona/myapp/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PersonHandler представляет обработчики для работы с людьми
type PersonHandler struct {
	Repo repository.PersonRepository
}

// GetAll Получить всех людей
func (h *PersonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	people, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get people: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// GetByID Получить человека по ID
func (h *PersonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL (например, /people/1)
	var id int
	fmt.Sscanf(r.URL.Path, "/people/%d", &id)
	person, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get person: %v", err), http.StatusInternalServerError)
		return
	}
	if person.ID == 0 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// Create Создать нового человека
func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode JSON: %v", err), http.StatusBadRequest)
		return
	}
	newPerson, err := h.Repo.Create(person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create person: %v", err), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(newPerson)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
