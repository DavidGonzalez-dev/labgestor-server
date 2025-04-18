package models

// Declare Foreign key tables related to product
type tipoProducto struct {
	ID         int `gorm:"primaryKey" json:"-"`
	nombreTipo string
}

type estadoProducto struct {
	ID           int `gorm:"primaryKey" json:"-"`
	nombreEstado string
}

type Producto struct {
	NumeroRegistro   string `gorm:"primaryKey"`
	Nombre           string
	FechaFabricacion string
	FechaVencimiento string
	Descripcion      string
	CompuestoActivo  string
	Presentacion     string
	Cantidad         string
	NumeroLote       string
	TamanoLote       string
	ClienteID        int            `json:"-"`
	FabricanteID     int            `json:"-"`
	TipoID           int            `json:"-"`
	EstadoID         int            `json:"-"`
	Cliente          Cliente        `gorm:"foreignKey: ClienteID"`
	Fabricante       Fabricante     `gorm:"foreignKey: FabricanteID"`
	Tipo             tipoProducto   `gorm:"foreignKey: TipoID"`
	Estado           estadoProducto `gorm:"foreignKey: EstadoID"`
}
