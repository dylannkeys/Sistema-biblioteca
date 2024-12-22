# Biblioteca Virtual - Sistema de Gestión

## Objetivo del Programa
El sistema de gestión de biblioteca virtual es una aplicación diseñada para administrar usuarios y libros en una biblioteca digital. Este programa permite realizar operaciones básicas como agregar, actualizar, eliminar y consultar información sobre libros y usuarios. Está desarrollado utilizando el lenguaje de programación Go y una base de datos SQL Server, asegurando un rendimiento eficiente y una gestión segura de los datos.

## Principales Funcionalidades

### Gestión de Libros
- **Agregar un libro**: Permite registrar nuevos libros en la base de datos especificando título, autor, categoría, año, estado de préstamo y precio.
- **Actualizar información del libro**: Modifica los datos existentes de un libro mediante su ID.
- **Eliminar un libro**: Elimina un libro específico del sistema utilizando su ID.
- **Consultar libros**: Muestra una lista de todos los libros disponibles en la base de datos.

### Gestión de Usuarios
- **Agregar un usuario**: Registra nuevos usuarios proporcionando nombre, email y teléfono.
- **Eliminar un usuario**: Elimina a un usuario existente del sistema utilizando su ID.

### Manejo de Errores
- Mensajes claros para operaciones exitosas y fallidas.
- Validación de entradas y manejo adecuado de errores SQL para garantizar integridad de los datos.

### Interfaz
- Implementado como un servidor HTTP que proporciona una API REST para interactuar con el sistema.
- Respuestas en formato JSON para facilitar la integración con clientes externos (como Postman o interfaces web).

## Estructura del Código
El código está organizado en varias funciones:
- **Conexión a la base de datos**: Establece la conexión con SQL Server.
- **Rutas HTTP**: Define las rutas y métodos disponibles para la API.
- **Controladores**: Implementa la lógica para cada operación (crear, eliminar, consultar, etc.).

## Tecnologías Utilizadas
- Lenguaje de programación: Go (Golang)
- Base de datos: SQL Server
- Herramientas: Postman para pruebas, Visual Studio Code como entorno de desarrollo

## Datos del Autor
- **Nombre**: [Dylan Nixon Moreno Chávez]
- **Fecha de Creación**: 22 de diciembre de 2024

## Conclusión
Este proyecto demuestra la aplicación de principios de programación orientada a objetos y desarrollo funcional en la construcción de sistemas de gestión. El uso de Go y SQL Server garantiza una solución robusta y escalable.
