package inputs

type ClientInput struct {
	CreditCardNum             string `json:"credit_card_num" binding:"required"`
	CreditCardCcv             string `json:"credit_card_ccv" binding:"required"`
	Id                        string `json:"id" binding:"required"`
	FecAlta                   string `json:"fec_alta" binding:"required"`
	ClientName                string `json:"user_name" binding:"required"`
	CodigoZip                 string `json:"codigo_zip" binding:"required"`
	CuentaNumero              string `json:"cuenta_numero" binding:"required"`
	Direccion                 string `json:"direccion" binding:"required"`
	GeoLatitud                string `json:"geo_latitud" binding:"required"`
	GeoLongitud               string `json:"geo_longitud" binding:"required"`
	ColorFavorito             string `json:"color_favorito" binding:"required"`
	FotoDni                   string `json:"foto_dni" binding:"required"`
	Ip                        string `json:"ip" binding:"required"`
	Auto                      string `json:"auto" binding:"required"`
	AutoModelo                string `json:"auto_modelo" binding:"required"`
	AutoTipo                  string `json:"auto_tipo" binding:"required"`
	AutoColor                 string `json:"auto_color" binding:"required"`
	CantidadComprasRealizadas int    `json:"cantidad_compras_realizadas" binding:"required"`
	Avatar                    string `json:"avatar" binding:"required"`
	FecBirthday               string `json:"fec_birthday" binding:"required"`
}

type BatchClientInput struct {
	Value []ClientInput `json:"value" binding:"required"`
}
