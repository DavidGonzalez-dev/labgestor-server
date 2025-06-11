package repository

import (
	"errors"
	"labgestor-server/internal/models"

	"gorm.io/gorm"
)

// Interfaz que define los metodos que se emplean en la tabla de los usuarios en la base datos
type UsuarioRepository interface {
	ObtenerUsuarioID(id string) (*models.Usuario, error)
	CrearUsuario(usuario *models.Usuario) error
	ActualizarUsuario(usuario *models.Usuario) error
	ObtenerUsuarios() (*[]models.Usuario, error)
	ObtenerUsuarioCorreo(correo string) (*models.Usuario, error)
}

// Structura que implementa la interfaz anteriormente definida
type usuarioRepository struct {
	DB *gorm.DB
}

// Funcion que devuelve una instancia de la estructura usuarioRepository
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{DB: db}
}

// ? ---------------------------
// ? Metodos de la estructura
// ? ---------------------------
// Este metodo permite obtener los detalles del registro de un usuario
func (repo *usuarioRepository) ObtenerUsuarioID(id string) (*models.Usuario, error) {

	// Instanciamos el modelo
	var usuario models.Usuario
	// Cargamos el usuario y las tablas anidadas a la tabla usuarios
	if err := repo.DB.Preload("Rol").First(&usuario, id).Error; err != nil {
		// En caso de error se retorna este error y no se retorna informacion
		return nil, err
	}
	// Si todo salio bien se retorna el modelo con la informacion del registro del usuario
	return &usuario, nil
}

// Este metodo permite crear un usuario en la base de datos
func (repo *usuarioRepository) CrearUsuario(usuario *models.Usuario) error {

	// Se crea el registro en base al modelo pasado como parametro
	if err := repo.DB.Create(&usuario).Error; err != nil {
		// Si hubo algun error se retorna
		return err
	}
	return nil
}

// Este metodo permite Actualizar un usuario en la base de datos
func (repo *usuarioRepository) ActualizarUsuario(usuario *models.Usuario) error {

	// Se hace un update especifico de cada campo para que no cause error al momento de actualizar las foreign keys
	if err := repo.DB.Model(&models.Usuario{}).Where("id = ?", usuario.ID).Updates(map[string]any{
		"nombres":   usuario.Nombres,
		"apellidos": usuario.Apellidos,
		"correo":    usuario.Correo,
		"contrasena": usuario.Contrasena,
		"firma":     usuario.Firma,
		"estado":    usuario.Estado,
		"rol_id":    usuario.RolID,
	}).Error; err != nil {
		// En caso de un error se retorna
		return err
	}

	return nil
}

// Este metodo permite obtener un slice con la informacion de todos los usuarios
func (repo *usuarioRepository) ObtenerUsuarios() (*[]models.Usuario, error) {
	// Se declara un slice para guaradar la informacion de los usuarios
	var usuarios []models.Usuario
	if err := repo.DB.Preload("Rol").Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return &usuarios, nil

}

// Este metodo nos permite obtener un usario dependiendo de su correo
func (repo *usuarioRepository) ObtenerUsuarioCorreo(correo string) (*models.Usuario, error) {
	
	var usuario models.Usuario

	if err := repo.DB.Where("correo=?", correo).First(&usuario).Error; err != nil{

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("este correo no esta registrado en nuestro sistema")
		}
		return nil, err
	} 

	return &usuario, nil
}
