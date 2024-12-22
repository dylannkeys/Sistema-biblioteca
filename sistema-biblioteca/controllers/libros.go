package controllers

import (
	"encoding/json"
	"net/http"
	"sistema-biblioteca/models"
)

// ObtenerLibros obtiene la lista de todos los libros
func ObtenerLibros(w http.ResponseWriter, r *http.Request) {
	libros, err := models.ObtenerLibros()
	if err != nil {
		http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}
