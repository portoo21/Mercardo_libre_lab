package models

import (
	"mercado-libre/inputs"
	"strconv"
	"time"
)

type ClientSensitive struct {
	CreditCardNum string
	CreditCardCcv string
}

type Client struct {
	ClientSensitive
	Id                        string
	FecAlta                   time.Time
	ClientName                string
	CodigoZip                 string
	CuentaNumero              int
	Direccion                 string
	GeoLatitud                float32
	GeoLongitud               float32
	ColorFavorito             string
	FotoDni                   string
	Ip                        string
	Auto                      string
	AutoModelo                string
	AutoTipo                  string
	AutoColor                 string
	CantidadComprasRealizadas int
	Avatar                    string
	FecBirthday               time.Time
}

func formatDate(input string) time.Time {
	date, err := time.Parse("02/01/2006", input)

	if err == nil {
		return date
	}

	date1, _ := time.Parse("2006-01-02T05:04:05.000Z", input)

	return date1
}

func (client *Client) FromInput(input inputs.ClientInput) {
	cuentaNumero, _ := strconv.ParseInt(input.CuentaNumero, 10, 0)
	geoLatitud, _ := strconv.ParseFloat(input.GeoLatitud, 32)
	geoLongitud, _ := strconv.ParseFloat(input.GeoLatitud, 32)

	client.ClientSensitive = ClientSensitive{
		CreditCardNum: input.CreditCardNum,
		CreditCardCcv: input.CreditCardCcv,
	}

	client.Id = input.Id
	client.FecAlta = formatDate(input.FecAlta)
	client.ClientName = input.ClientName
	client.CodigoZip = input.CodigoZip
	client.CuentaNumero = int(cuentaNumero)
	client.Direccion = input.Direccion
	client.GeoLatitud = float32(geoLatitud)
	client.GeoLongitud = float32(geoLongitud)
	client.ColorFavorito = input.ColorFavorito
	client.FotoDni = input.FotoDni
	client.Ip = input.Ip
	client.Auto = input.Auto
	client.AutoModelo = input.AutoModelo
	client.AutoTipo = input.AutoTipo
	client.AutoColor = input.AutoColor
	client.CantidadComprasRealizadas = input.CantidadComprasRealizadas
	client.Avatar = input.Avatar
	client.FecBirthday = formatDate(input.FecBirthday)
}
