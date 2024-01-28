CREATE SCHEMA stocks AUTHORIZATION stocks_app;

GRANT ALL ON SCHEMA stocks TO stocks_app;

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA stocks
GRANT ALL ON TABLES TO stocks_app;

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA stocks
GRANT ALL ON SEQUENCES TO stocks_app;

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA stocks
GRANT EXECUTE ON FUNCTIONS TO stocks_app;

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA stocks
GRANT USAGE ON TYPES TO stocks_app;


CREATE TABLE stocks.tickers
(
    id integer NOT NULL,
    ticker character(4) NOT NULL,
    watch boolean NOT NULL,
    bought boolean NOT NULL,
    CONSTRAINT id_pk PRIMARY KEY (id),
    CONSTRAINT ticker UNIQUE (ticker)
        WITH (FILLFACTOR=95)
);

ALTER TABLE IF EXISTS stocks.tickers
    OWNER to stocks_app;

GRANT ALL ON TABLE stocks.tickers TO stocks_app;

SET search_path TO stocks;

CREATE SEQUENCE stocks.tickers_id
    INCREMENT 1
    START 1;

ALTER SEQUENCE stocks.tickers_id
    OWNER TO stocks_app;

GRANT ALL ON SEQUENCE stocks.tickers_id TO stocks_app;

-- vincula a sequence com a coluna da tabela
ALTER SEQUENCE IF EXISTS stocks.tickers_id
    OWNED BY tickers.id;

-- cadastrar ações
set search_path to stocks;
INSERT INTO stocks.tickers(
	id, ticker, watch, bought)
    -- id (sequence), ticker, watch (t,f), bought(t,f)
	VALUES (nextval('tickers_id'), 'PETR4', true, false);


-- adição campos de alvo de compra e venda
ALTER TABLE IF EXISTS stocks.tickers
    ADD COLUMN target_b real;

ALTER TABLE IF EXISTS stocks.tickers
    ADD COLUMN target_s real;
    