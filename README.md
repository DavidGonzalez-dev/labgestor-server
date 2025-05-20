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
  - `200 OK`: Login exitoso. Retorna información del usuario y establece cookie de autenticación.
  - `404 Not Found`: Error al leer el request
  - `401 Unauthorized`: Credenciales inválidas o usuario deshabilitado
  - `500 Internal Server Error`: Error al generar el token

### Logout
- **URL**: `/logout`
- **Método**: `POST`
- **Autenticación**: Requerida
- **Respuestas**:
  - `200 OK`: Sesión cerrada exitosamente. La cookie de autenticación es invalidada.

### Validar Token
- **URL**: `/validar-token`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Token válido. Retorna estado del token y rol del usuario.
  - `401 Unauthorized`: Token inválido o expirado
  - `500 Internal Server Error`: Error al decodificar el token

### Cambiar Contraseña
- **URL**: `/contrasena/actualizar`
- **Método**: `PATCH`
- **Request Body**:
```json
{
    "id": "string",
    "contrasena": "string"
}
```
- **Respuestas**:
  - `200 OK`: Contraseña actualizada exitosamente
  - `404 Not Found`: Error al leer el request o usuario no encontrado
  - `400 Bad Request`: Error al encriptar la contraseña

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
**Validaciones**:
- `documento`: Solo puede contener números
- `nombres`: Solo puede contener letras y espacios
- `apellidos`: Solo puede contener letras y espacios
- `correo`: Debe ser un email válido (ejemplo@dominio.com)

- **Respuestas**:
  - `201 Created`: Usuario registrado exitosamente
  - `404 Not Found`: Error al leer el request
  - `400 Bad Request`: Error de validación o rol inválido o usuario ya existe
  - `500 Internal Server Error`: Error al crear el usuario

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
**Validaciones**:
- `nombres`: Solo puede contener letras y espacios
- `apellidos`: Solo puede contener letras y espacios
- `correo`: Debe ser un email válido

- **Respuestas**:
  - `200 OK`: Usuario actualizado exitosamente
  - `400 Bad Request`: Error al leer el request o error de validación
  - `404 Not Found`: Usuario no encontrado
  - `500 Internal Server Error`: Error al actualizar el usuario

### Obtener Perfil
- **URL**: `/usuarios/:id`
- **Método**: `GET`
- **Respuestas**:
  - `200 OK`: Retorna información del usuario
  - `400 Bad Request`: ID inválida
  - `404 Not Found`: Error al obtener el usuario

### Obtener Usuarios
- **URL**: `/usuarios`
- **Método**: `GET`
- **Autenticación**: Requerida
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
  - `302 Found`: Retorna información completa del producto incluyendo detalles del registro de entrada
  - `404 Not Found`: Producto no encontrado
  - `400 Bad Request`: Error al leer el cuerpo del request

### Obtener Registros de Entrada
- **URL**: `/registroEntradaProductos/`
- **Método**: `GET`
- **Descripción**: Devuelve un array con los registros de entrada de los productos sin detalles
- **Respuestas**:
  - `200 OK`: Lista de registros de entrada sin detalles del producto
  - `404 Not Found`: Error al obtener los registros de entrada

### Crear Producto
- **URL**: `/productos/crear`
- **Método**: `POST`
- **Descripción**: Crea un producto en la base de datos con su respectivo registro de entrada al área
- **Request Body**:
```json
{
    "producto": {
        "numeroRegistro": "string",
        "nombre": "string",
        "fechaFabricacion": "string",
        "fechaVencimiento": "string",
        "descripcion": "string",
        "compuestoActivo": "string",
        "presentacion": "string",
        "cantidad": "string",
        "numeroLote": "string",
        "tamanoLote": "string",
        "idCliente": "number",
        "idFabricante": "number",
        "idTipo": "number"
    },
    "detallesEntrada": {
        "propositoAnalisis": "string",
        "condicionesAmbientales": "string",
        "fechaRecepcion": "string",
        "fechaInicioAnalisis": "string",
        "fechaFinalAnalisis": "string",
        "idUsuario": "string"
    }
}
```
**Validaciones**:
- Las fechas son opcionales y pueden ser cadenas vacías
- El estado del producto se establece automáticamente como 1
- Se validan los campos usando las reglas definidas en `ProductoRules` y `RegistroEntradaRules`

- **Respuestas**:
  - `201 Created`: Producto creado exitosamente
  - `400 Bad Request`: Error al leer el cuerpo del request o error en el formato de los datos
  - `422 Unprocessable Entity`: Error en la validación de los campos
  - `500 Internal Server Error`: Error al crear el producto

### Actualizar Producto
- **URL**: `/productos/actualizar`
- **Método**: `PUT`
- **Descripción**: Actualiza la información de un producto existente
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
    "cantidad": "string",
    "numeroLote": "string",
    "tamanoLote": "string",
    "idCliente": "number",
    "idFabricante": "number",
    "idTipo": "number"
}
```
**Validaciones**:
- Las fechas son opcionales y pueden ser cadenas vacías
- Se validan los campos usando las reglas definidas en `ProductoRules`

- **Respuestas**:
  - `200 OK`: Producto actualizado exitosamente
  - `400 Bad Request`: Error al leer el cuerpo del request
  - `404 Not Found`: Producto no encontrado
  - `422 Unprocessable Entity`: Error en la validación de campos
  - `500 Internal Server Error`: Error al actualizar el producto

### Actualizar Registro de Entrada
- **URL**: `/registroEntradaProductos/actualizar`
- **Método**: `PUT`
- **Descripción**: Actualiza el registro de entrada de un producto existente
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
**Validaciones**:
- Las fechas son opcionales y pueden ser cadenas vacías
- Se validan los campos usando las reglas definidas en `RegistroEntradaRules`
- No se valida el campo `IDUsuario` en las actualizaciones

- **Respuestas**:
  - `200 OK`: Registro actualizado exitosamente
  - `400 Bad Request`: Error al leer el cuerpo del request
  - `404 Not Found`: Registro de entrada no encontrado
  - `422 Unprocessable Entity`: Error en la validación de campos
  - `500 Internal Server Error`: Error al actualizar el registro

### Eliminar Producto
- **URL**: `/productos/:id`
- **Método**: `DELETE`
- **Descripción**: Elimina un producto y todos sus registros relacionados
- **Respuestas**:
  - `200 OK`: Producto eliminado exitosamente
  - `400 Bad Request`: Error al leer el ID del producto
  - `404 Not Found`: Producto no encontrado
  - `500 Internal Server Error`: Error al eliminar el producto

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
