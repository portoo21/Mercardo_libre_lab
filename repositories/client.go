package repositories

import (
	"log"
	"mercado-libre/models"
	"mercado-libre/utils"

	"github.com/jackc/pgx"
)

type ClientRepository struct {
	regularCon   *pgx.Conn
	sensitiveCon *pgx.Conn
}

func InitClientRepository() (*ClientRepository, error) {
	regularCon, err := connectDb(getDbUser(REGULAR))
	if err != nil {
		return nil, err
	}

	sensitiveCon, err := connectDb(getDbUser(SENSITIVE))
	if err != nil {
		return nil, err
	}

	repository := ClientRepository{
		regularCon:   regularCon,
		sensitiveCon: sensitiveCon,
	}

	return &repository, nil
}

func (r *ClientRepository) BatchUpdate(clients []models.Client) (int, error) {
	var rowsToken [][]interface{}
	for _, client := range clients {
		encryptedCreditCardNum, err := utils.GetAESEncrypted(client.CreditCardNum)
		if err != nil {
			continue
		}
		encryptedCreditCardCcv, err := utils.GetAESEncrypted(client.CreditCardCcv)
		if err != nil {
			continue
		}

		rowsToken = append(rowsToken, []interface{}{
			client.Id, client.FecAlta, client.ClientName, client.CodigoZip, encryptedCreditCardNum,
			encryptedCreditCardCcv, client.CuentaNumero, client.Direccion, client.GeoLatitud,
			client.GeoLongitud, client.ColorFavorito, client.FotoDni, client.Ip, client.Auto,
			client.AutoModelo, client.AutoTipo, client.AutoColor, client.CantidadComprasRealizadas,
			client.Avatar, client.FecBirthday})
	}

	return r.regularCon.CopyFrom(
		pgx.Identifier{"clients"},
		[]string{
			"id", "fec_alta", "client_name", "codigo_zip", "credit_card_num", "credit_card_ccv",
			"cuenta_numero", "direccion", "geo_latitud", "geo_longitud", "color_favorito",
			"foto_dni", "ip", "auto", "auto_modelo", "auto_tipo", "auto_color",
			"cantidad_compras_realizadas", "avatar", "fec_birthday",
		},
		pgx.CopyFromRows(rowsToken),
	)
}

func (r *ClientRepository) GetClients() []models.Client {
	rows, err := r.regularCon.Query("SELECT * FROM view_clients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var client models.Client
		err := rows.Scan(
			&client.Id, &client.FecAlta, &client.ClientName, &client.CodigoZip, &client.CuentaNumero,
			&client.Direccion, &client.GeoLatitud, &client.GeoLongitud,
			&client.ColorFavorito, &client.FotoDni, &client.Ip, &client.Auto, &client.AutoModelo,
			&client.AutoTipo, &client.AutoColor, &client.CantidadComprasRealizadas, &client.Avatar,
			&client.FecBirthday,
		)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, client)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return clients
}

func (r *ClientRepository) GetClientSensitiveData(id string) (models.ClientSensitive, error) {
	row := r.sensitiveCon.QueryRow("SELECT credit_card_num, credit_card_ccv FROM clients WHERE id = $1", id)

	var client models.ClientSensitive
	err := row.Scan(&client.CreditCardNum, &client.CreditCardCcv)
	if err != nil {
		return models.ClientSensitive{}, err
	}

	creditCardNumDescrypted, err := utils.GetAESDecrypted(client.CreditCardNum)

	if err != nil {
		return models.ClientSensitive{}, err
	}

	creditCardCcvDescrypted, err := utils.GetAESDecrypted(client.CreditCardCcv)

	if err != nil {
		return models.ClientSensitive{}, err
	}

	client.CreditCardNum = string(creditCardNumDescrypted)
	client.CreditCardCcv = string(creditCardCcvDescrypted)

	return client, nil
}

func (r *ClientRepository) Close() {
	r.regularCon.Close()
	r.sensitiveCon.Close()
}
