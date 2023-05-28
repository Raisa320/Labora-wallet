CREATE DATABASE labora_wallet
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public.wallet
(
    id integer NOT NULL DEFAULT 'nextval('wallet_id_seq'::regclass)',
    person_id character varying(15) COLLATE pg_catalog."default" NOT NULL,
    date date NOT NULL,
    country character varying(4) COLLATE pg_catalog."default" NOT NULL,
    amount numeric NOT NULL DEFAULT 0,
    have_card boolean DEFAULT 'false',
    CONSTRAINT wallet_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.wallet
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public.log
(
    id integer NOT NULL DEFAULT 'nextval('log_id_seq'::regclass)',
    person_id character varying COLLATE pg_catalog."default" NOT NULL,
    date date NOT NULL,
    status boolean NOT NULL,
    country character varying(4) COLLATE pg_catalog."default" NOT NULL,
    check_id character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT log_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.log
    OWNER to postgres;