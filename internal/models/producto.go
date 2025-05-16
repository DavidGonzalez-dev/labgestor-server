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
	NumeroRegistro   string          `gorm:"primaryKey"`
	Nombre           string          `json:"nombre"`
	FechaFabricacion string          `json:"fechaFabricacion"`
	FechaVencimiento string          `json:"fechaVencimiento"`
	Descripcion      string          `json:"descripcion"`
	CompuestoActivo  string          `json:"compuestoActivo"`
	Presentacion     string          `json:"presentacion"`
	Cantidad         string          `json:"cantidad"`
	NumeroLote       string          `json:"numeroLote"`
	TamanoLote       string          `json:"tamanoLote"`
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
		"FechaVencimiento": producto.FechaVencimiento,
		"Descripcion":      producto.Descripcion,
		"CompuestoActivo":  producto.CompuestoActivo,
		"Presentacion":     producto.Presentacion,
		"Cantidad":         producto.Cantidad,
		"NumeroLote":       producto.NumeroLote,
		"TamanoLote":       producto.TamanoLote,
	}
}

// Declaracion del modelo de entrada de un producto
type RegistroEntradaProducto struct {
	CodigoEntrada          int `gorm:"primaryKey"`
	PropositoAnalisis      string
	CondicionesAmbientales string
	FechaRecepcion         string
	FechaInicioAnalisis    string
	FechaFinalAnalisis     string
	NumeroRegistroProducto string    `json:"numeroRegistroProducto"`
	Producto               *Producto `gorm:"foreignKey:NumeroRegistroProducto" json:"producto,omitempty"`
	IDUsuario              string    `json:"-"`
	Usuario                *Usuario  `gorm:"foreignKey: IDUsuario" json:"usuario,omitempty"`
}

func (entrada RegistroEntradaProducto) ToMap() map[string]any {
	return map[string]any{
		"CodigoEntrada":          entrada.CodigoEntrada,
		"PropositoAnalisis":      entrada.PropositoAnalisis,
		"CondicionesAmbientales": entrada.CondicionesAmbientales,
		"FechaRecepcion":         entrada.FechaRecepcion,
		"FechaInicioAnalisis":    entrada.FechaInicioAnalisis,
		"FechaFinalAnalisis":     entrada.FechaFinalAnalisis,
		"IDUsuario":              entrada.IDUsuario,
		"NumeroRegistroProducto": entrada.NumeroRegistroProducto,
	}
}
