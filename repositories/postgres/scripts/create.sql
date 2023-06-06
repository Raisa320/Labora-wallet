CREATE DATABASE labora_wallet
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS wallet
(
    id SERIAL PRIMARY KEY,
    person_id character varying(15) COLLATE pg_catalog."default" NOT NULL,
    date date NOT NULL,
    country character varying(4) COLLATE pg_catalog."default" NOT NULL,
    amount numeric NOT NULL DEFAULT 0,
    have_card boolean DEFAULT 'false'
);

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.wallet
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS log
(
    id SERIAL PRIMARY KEY,
    person_id character varying COLLATE pg_catalog."default",
    date date NOT NULL,
    status boolean NOT NULL,
    country character varying(4) COLLATE pg_catalog."default",
    check_id character varying COLLATE pg_catalog."default",
    type character varying(20) COLLATE pg_catalog."default",
    message character varying COLLATE pg_catalog."default"
);

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.log
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS transaction
(
    id SERIAL PRIMARY KEY,
    amount numeric NOT NULL,
    destiny_id integer NOT NULL,
    source_id integer NOT NULL,
    type integer NOT NULL,
    date date NOT NULL DEFAULT CURRENT_DATE,
    CONSTRAINT "Transaction_destiny_id_fkey" FOREIGN KEY (destiny_id)
        REFERENCES wallet (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Transaction_source_id_fkey" FOREIGN KEY (source_id)
        REFERENCES wallet (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.transaction
    OWNER to postgres;
