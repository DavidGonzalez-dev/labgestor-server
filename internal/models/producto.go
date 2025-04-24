package models

// Declare Foreign key tables related to product
type tipoProducto struct {
	ID         int `gorm:"primaryKey" json:"-"`
	NombreTipo string
}

type estadoProducto struct {
	ID           int `gorm:"primaryKey" json:"-"`
	NombreEstado string
}

// Declaracion del modelo de un Producto
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
	IDCliente        int             `json:"-"`
	Cliente          *Cliente        `gorm:"foreignKey: IDCliente" json:"cliente,omitempty"`
	IDFabricante     int             `json:"-"`
	Fabricante       *Fabricante     `gorm:"foreignKey: IDFabricante" json:"fabricante,omitempty"`
	IDTipo           int             `json:"-"`
	TipoProducto     *tipoProducto   `gorm:"foreignKey: IDTipo" json:"tipo,omitempty"`
	IDEstado         int             `json:"-"`
	EstadoProducto   *estadoProducto `gorm:"foreignKey: IDEstado" json:"estado,omitempty"`
}

func (producto Producto) ToMap() map[string]any {
	return map[string]any{
		"NumeroRegistro":   producto.NumeroRegistro,
		"Nombre":           producto.Nombre,
		"FechaFabricacion": producto.FechaFabricacion,
		"FechaVencimiento": producto.FechaFabricacion,
		"Descripcion":      producto.FechaFabricacion,
		"CompuestoActivo":  producto.CompuestoActivo,
		"Presentacion":     producto.Presentacion,
		"Cantidad":         producto.Cantidad,
		"NumeroLote":       producto.NumeroLote,
		"TamanoLote":       producto.TamanoLote,
	}
}

// Declaracion del modelo de entrada de un producto
type EntradaProducto struct {
	CodigoEntrada          int `gorm:"primaryKey"`
	PropositoAnalisis      string
	CondicionesAmbientales string
	FechaRecepcion         string
	FechaInicioAnalisis    string
	FechaFinalAnalisis     string
	IDUsuario              string    `json:"-"`
	Usuario                *Usuario  `gorm:"foreignKey: IDUsuario" json:"usuario,omitempty"`
	NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
	Producto               *Producto `gorm:"foreignKey:NumeroRegistroProducto" json:"producto,omitempty"`
}

func (entrada EntradaProducto) ToMap() map[string]any {
	return map[string]any{
		"CodigoEntrada":          entrada.CodigoEntrada,
		"PropositoAnalisis":      entrada.PropositoAnalisis,
		"CondicionesAmbientales": entrada.CondicionesAmbientales,
		"FechaRecepcion":         entrada.FechaRecepcion,
		"FechaInicioAnalsis":     entrada.FechaInicioAnalisis,
		"FechaFinalAnalisis":     entrada.FechaFinalAnalisis,
	}
}
