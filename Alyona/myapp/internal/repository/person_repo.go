package repository

import (
	"awesomeProject/Alyona/myapp/internal/models"
)

// PersonRepository представляет интерфейс для работы с данными людей
type PersonRepository interface {
	GetAll() ([]models.Person, error)
	GetByID(id int) (models.Person, error)
	Create(person models.Person) (models.Person, error)
}

// InMemoryPersonRepository — это репозиторий, хранящий данные в памяти
type InMemoryPersonRepository struct {
	//mu     sync.RWMutex
	people []models.Person
	nextID int
}

// NewInMemoryPersonRepository Новый репозиторий
func NewInMemoryPersonRepository() *InMemoryPersonRepository {
	return &InMemoryPersonRepository{
		people: []models.Person{},
		nextID: 1,
	}
}

// GetAll Получить всех людей
func (r *InMemoryPersonRepository) GetAll() ([]models.Person, error) {
	//r.mu.RLock()
	//defer r.mu.RUnlock()
	return r.people, nil
}

// GetByID Получить человека по ID
func (r *InMemoryPersonRepository) GetByID(id int) (models.Person, error) {
	//r.mu.RLock()
	//defer r.mu.RUnlock()
	for _, p := range r.people {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Person{}, nil
}

// Create Создать нового человека
func (r *InMemoryPersonRepository) Create(person models.Person) (models.Person, error) {
	//r.mu.Lock()
	//defer r.mu.Unlock()
	person.ID = r.nextID
	r.nextID++
	r.people = append(r.people, person)
	return person, nil
}
