CREATE TABLE clients 
(
    id	VARCHAR(50) PRIMARY KEY,
    fec_alta	TIMESTAMP WITHOUT TIME ZONE,
    client_name	VARCHAR(100),
    codigo_zip	VARCHAR(100),
    credit_card_num	TEXT,
    credit_card_ccv	TEXT,
    cuenta_numero	NUMERIC(12, 0),
    direccion	VARCHAR(100),
    geo_latitud	numeric(5,3),
    geo_longitud numeric(5,3),
    color_favorito	VARCHAR(100),
    foto_dni	VARCHAR(100),
    ip	VARCHAR(16),
    "auto"	VARCHAR(100),
    auto_modelo	VARCHAR(100),
    auto_tipo	VARCHAR(100),
    auto_color	VARCHAR(100),
    cantidad_compras_realizadas	INT,
    avatar	VARCHAR(100),
    fec_birthday	TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE audit_sensitive_data
(
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50),
    client_id VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT fk_client
        FOREIGN KEY(client_id) 
	    REFERENCES clients(id)
);

CREATE VIEW view_clients AS
  SELECT id,
    fec_alta,
    client_name,
    codigo_zip,
    cuenta_numero,
    direccion,
    geo_latitud,
    geo_longitud,
    color_favorito,
    foto_dni,
    ip,
    "auto",
    auto_modelo,
    auto_tipo,
    auto_color,
    cantidad_compras_realizadas,
    avatar,
    fec_birthday
  FROM clients;

CREATE USER app_user WITH PASSWORD 'app-password';

GRANT CONNECT ON DATABASE meli TO app_user;
GRANT INSERT ON clients TO app_user;
GRANT SELECT ON view_clients TO app_user;

CREATE USER app_sensitive_user WITH PASSWORD 'app-password-1';
GRANT CONNECT ON DATABASE meli TO app_sensitive_user;
GRANT SELECT ON clients TO app_sensitive_user;