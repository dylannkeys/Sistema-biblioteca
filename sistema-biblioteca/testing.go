package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Cadena de conexión a la base de datos
	connStr := "sqlserver://localhost:1433?database=BibliotecaDB&trusted_connection=true"
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		fmt.Println("Error al conectar:", err)
		return
	}
	defer db.Close()

	// Intentar hacer ping a la base de datos
	err = db.Ping()
	if err != nil {
		fmt.Println("Error al hacer ping a la base de datos:", err)
	} else {
		fmt.Println("Conexión exitosa a la base de datos.")
	}

	// Realizar consulta para obtener libros
	rows, err := db.Query("SELECT LibroID, Titulo, Autor, Categoria, Año, Prestado, Precio FROM Libros")
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return
	}
	defer rows.Close()

	// Iterar sobre los resultados
	for rows.Next() {
		var libroID int
		var titulo, autor, categoria string
		var año int
		var prestado bool
		var precio sql.NullFloat64 // Cambiar a sql.NullFloat64

		err := rows.Scan(&libroID, &titulo, &autor, &categoria, &año, &prestado, &precio)
		if err != nil {
			fmt.Println("Error al leer los resultados:", err)
			return
		}

		// Imprimir los detalles del libro
		if precio.Valid {
			fmt.Printf("LibroID: %d, Titulo: %s, Autor: %s, Categoria: %s, Año: %d, Prestado: %v, Precio: %.2f\n",
				libroID, titulo, autor, categoria, año, prestado, precio.Float64)
		} else {
			fmt.Printf("LibroID: %d, Titulo: %s, Autor: %s, Categoria: %s, Año: %d, Prestado: %v, Precio: N/A\n",
				libroID, titulo, autor, categoria, año, prestado)
		}
	}

	// Verificar si hubo errores al iterar
	if err := rows.Err(); err != nil {
		fmt.Println("Error al iterar sobre los resultados:", err)
	}
}
