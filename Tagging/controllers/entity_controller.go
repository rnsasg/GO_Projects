package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/rnsasg/Tagging/models"
	"github.com/rnsasg/Tagging/services"
	"net/http"
)

type EntityController struct {
	entityService services.EntityService
}

func NewEntityController(entiryService services.EntityService) *EntityController {
	fmt.Println("NewEntityController")
	return &EntityController{entiryService}
}

func (c *EntityController) CreateEntity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateEntity")
	var entity models.Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.entityService.CreateEntity(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *EntityController) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteEntity")
	params := mux.Vars(r)
	id, _ := gocql.ParseUUID(params["id"])

	if err := c.entityService.DeleteEntity(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *EntityController) GetEntity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetEntity")
	params := mux.Vars(r)
	id, _ := gocql.ParseUUID(params["id"])
	entity, err := c.entityService.GetEntityByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func (c *EntityController) GetAllEntities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllEntities")
	entities, err := c.entityService.GetAllEntities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entities)
}

func (c *EntityController) UpdateEntity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEntity")
	params := mux.Vars(r)
	id, _ := gocql.ParseUUID(params["id"])
	var entity models.Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	entity.ID = id
	if err := c.entityService.UpdateEntity(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
