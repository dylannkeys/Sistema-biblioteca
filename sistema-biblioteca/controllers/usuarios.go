// controllers/usuarios.go
package controllers

import (
	"encoding/json"
	"net/http"
	"sistema-biblioteca/database"
	"sistema-biblioteca/models"
)

// ObtenerUsuarios maneja la solicitud GET para obtener los usuarios
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	// Establecer la conexión a la base de datos
	db, err := database.Conectar()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Obtener los usuarios de la base de datos
	usuarios, err := models.ObtenerUsuarios(db)
	if err != nil {
		http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}

	// Devolver la lista de usuarios como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}
