package models

type Producto struct {
	Numero_Registro   string `gorm:"foreignKey"`
	Nombre            string
	Fecha_fabricacion string
	Fecha_vencimiento string
	Descripcion       string
	Compuesto_activo  string
	Presentacion      string
	Cantidad          string
	Numero_lote       string
	Tamano_lote       string
	Id_cliente        int
	Id_fabricante     int
	Id_tipo           int
	Id_estado         int
}
