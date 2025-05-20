# LabGestor Server API Documentation

Este documento describe las APIs implementadas en el servidor LabGestor. La API está construida usando Go y el framework Echo.

## Índice
1. [Autenticación](#autenticación)
2. [Usuarios](#usuarios)
3. [Clientes](#clientes)
4. [Fabricantes](#fabricantes)
5. [Productos](#productos)

## Autenticación

### Login
- **URL**: `/login`
- **Método**: `POST`
- **Request Body**:
```json
{
    "documento": "string",
    "password": "string"
}
```
- **Respuestas**:
  - `200 OK`: Login exitoso
  - `400 Bad Request`: Error en los datos de entrada
  - `401 Unauthorized`: Credenciales inválidas

### Logout
- **URL**: `/logout`
- **Método**: `POST`
- **Autenticación**: Requerida
- **Respuestas**:
  - `200 OK`: Sesión cerrada exitosamente
  - `401 Unauthorized`: Token inválido o expirado

### Validar Token
- **URL**: `/validar-token`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Token válido
  - `401 Unauthorized`: Token inválido o expirado

### Cambiar Contraseña
- **URL**: `/contrasena/actualizar`
- **Método**: `PATCH`
- **Request Body**:
```json
{
    "contrasenaActual": "string",
    "nuevaContrasena": "string"
}
```
- **Respuestas**:
  - `200 OK`: Contraseña actualizada exitosamente
  - `400 Bad Request`: Error en los datos de entrada
  - `401 Unauthorized`: Contraseña actual incorrecta

## Usuarios

### Registrar Usuario
- **URL**: `/usuarios/registrar`
- **Método**: `POST`
- **Request Body**:
```json
{
    "documento": "string",
    "nombres": "string",
    "apellidos": "string",
    "correo": "string",
    "rol": "number"
}
```
- **Respuestas**:
  - `201 Created`: Usuario creado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `409 Conflict`: El usuario ya existe

### Actualizar Usuario
- **URL**: `/usuarios/actualizar`
- **Método**: `PUT`
- **Request Body**:
```json
{
    "id": "string",
    "nombres": "string",
    "apellidos": "string",
    "correo": "string",
    "estado": "boolean",
    "rolId": "number"
}
```
- **Respuestas**:
  - `200 OK`: Usuario actualizado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `404 Not Found`: Usuario no encontrado

### Obtener Perfil
- **URL**: `/usuarios/:id`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Retorna información del usuario
  - `404 Not Found`: Usuario no encontrado

### Obtener Usuarios
- **URL**: `/usuarios`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Lista de usuarios
  - `500 Internal Server Error`: Error al obtener los usuarios

### Deshabilitar Usuario
- **URL**: `/usuarios/:id`
- **Método**: `DELETE`
- **Respuestas**:
  - `200 OK`: Usuario deshabilitado exitosamente
  - `404 Not Found`: Usuario no encontrado

## Clientes

### Crear Cliente
- **URL**: `/clientes/registrar`
- **Método**: `POST`
- **Request Body**:
```json
{
    "nombre": "string",
    "direccion": "string"
}
```
- **Respuestas**:
  - `201 Created`: Cliente creado exitosamente
  - `400 Bad Request`: Error en el formato de los datos

### Actualizar Cliente
- **URL**: `/clientes/actualizar`
- **Método**: `PUT`
- **Request Body**:
```json
{
    "id": "number",
    "nombre": "string",
    "direccion": "string"
}
```
- **Respuestas**:
  - `200 OK`: Cliente actualizado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `404 Not Found`: Cliente no encontrado

### Obtener Cliente
- **URL**: `/clientes/:id`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Retorna información del cliente
  - `404 Not Found`: Cliente no encontrado

### Obtener Clientes
- **URL**: `/clientes`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Lista de clientes
  - `404 Not Found`: Error al obtener la información

### Eliminar Cliente
- **URL**: `/clientes/:id`
- **Método**: `DELETE`
- **Respuestas**:
  - `200 OK`: Cliente eliminado exitosamente
  - `400 Bad Request`: ID inválido
  - `404 Not Found`: Cliente no encontrado

## Fabricantes

### Crear Fabricante
- **URL**: `/fabricantes/registrar`
- **Método**: `POST`
- **Request Body**:
```json
{
    "nombre": "string",
    "direccion": "string"
}
```
- **Respuestas**:
  - `201 Created`: Fabricante creado exitosamente
  - `400 Bad Request`: Error en el formato de los datos

### Actualizar Fabricante
- **URL**: `/fabricantes/actualizar`
- **Método**: `PUT`
- **Request Body**:
```json
{
    "id": "number",
    "nombre": "string",
    "direccion": "string"
}
```
- **Respuestas**:
  - `200 OK`: Fabricante actualizado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `404 Not Found`: Fabricante no encontrado

### Obtener Fabricante
- **URL**: `/fabricantes/:id`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Retorna información del fabricante
  - `404 Not Found`: Fabricante no encontrado

### Obtener Fabricantes
- **URL**: `/fabricantes`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Lista de fabricantes
  - `404 Not Found`: Error al obtener la información

### Eliminar Fabricante
- **URL**: `/fabricantes/:id`
- **Método**: `DELETE`
- **Respuestas**:
  - `200 OK`: Fabricante eliminado exitosamente
  - `400 Bad Request`: ID inválido
  - `404 Not Found`: Fabricante no encontrado

## Productos

### Obtener Producto
- **URL**: `/productos/:id`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Retorna información del producto
  - `404 Not Found`: Producto no encontrado

### Obtener Registros de Entrada
- **URL**: `/registroEntradaProductos/`
- **Método**: `GET`
- **Respuestas**:
  - `302 Found`: Lista de registros de entrada
  - `500 Internal Server Error`: Error al obtener los registros

### Crear Producto
- **URL**: `/productos/crear`
- **Método**: `POST`
- **Request Body**:
```json
{
    "nombre": "string",
    "fechaFabricacion": "string",
    "fechaVencimiento": "string",
    "descripcion": "string",
    "compuestoActivo": "string",
    "presentacion": "string",
    "cantidad": "number",
    "numeroLote": "string",
    "tamanoLote": "string",
    "idCliente": "number",
    "idFabricante": "number",
    "idTipo": "number"
}
```
- **Respuestas**:
  - `201 Created`: Producto creado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `422 Unprocessable Entity`: Error en la validación de datos

### Actualizar Producto
- **URL**: `/productos/actualizar`
- **Método**: `PUT`
- **Request Body**:
```json
{
    "numeroRegistro": "string",
    "nombre": "string",
    "fechaFabricacion": "string",
    "fechaVencimiento": "string",
    "descripcion": "string",
    "compuestoActivo": "string",
    "presentacion": "string",
    "cantidad": "number",
    "numeroLote": "string",
    "tamanoLote": "string",
    "idCliente": "number",
    "idFabricante": "number",
    "idTipo": "number"
}
```
- **Respuestas**:
  - `200 OK`: Producto actualizado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `404 Not Found`: Producto no encontrado
  - `422 Unprocessable Entity`: Error en la validación de datos

### Actualizar Registro de Entrada
- **URL**: `/registroEntradaProductos/actualizar`
- **Método**: `PUT`
- **Request Body**:
```json
{
    "numeroRegistroProducto": "string",
    "propositoAnalisis": "string",
    "condicionesAmbientales": "string",
    "fechaRecepcion": "string",
    "fechaInicioAnalisis": "string",
    "fechaFinalAnalisis": "string"
}
```
- **Respuestas**:
  - `200 OK`: Registro actualizado exitosamente
  - `400 Bad Request`: Error en el formato de los datos
  - `404 Not Found`: Registro no encontrado
  - `422 Unprocessable Entity`: Error en la validación de datos

### Eliminar Producto
- **URL**: `/productos/:id`
- **Método**: `DELETE`
- **Respuestas**:
  - `200 OK`: Producto eliminado exitosamente
  - `404 Not Found`: Producto no encontrado

## Formato de Respuesta Global

Todas las respuestas siguen un formato estandarizado:

```json
{
    "message": "string",
    "error": "string (opcional)",
    "data": "any (opcional)"
}
```

## Códigos de Estado HTTP

- `200 OK`: Operación exitosa
- `201 Created`: Recurso creado exitosamente
- `302 Found`: Recurso encontrado
- `400 Bad Request`: Error en la solicitud
- `401 Unauthorized`: No autorizado
- `404 Not Found`: Recurso no encontrado
- `422 Unprocessable Entity`: Error de validación
- `500 Internal Server Error`: Error interno del servidor
