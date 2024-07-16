package routers

import (
	"github.com/gorilla/mux"
	"github.com/rnsasg/Tagging/controllers"
)

func InitTaggingRoutes(entityController *controllers.EntityController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tags", entityController.GetAllEntities).Methods("GET")
	router.HandleFunc("/tags/{id}", entityController.GetEntity).Methods("GET")
	router.HandleFunc("/tags", entityController.CreateEntity).Methods("POST")
	router.HandleFunc("/tags/{id}", entityController.DeleteEntity).Methods("DELETE")
	router.HandleFunc("/tags/{id}", entityController.UpdateEntity).Methods("PUT")

	return router
}
