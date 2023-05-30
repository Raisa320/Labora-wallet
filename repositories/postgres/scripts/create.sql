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


CREATE TABLE IF NOT EXISTS public.transaction
(
    id integer NOT NULL DEFAULT 'nextval('transaction_id_seq'::regclass)',
    amount numeric NOT NULL,
    type character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "walletReceive_id" integer NOT NULL,
    "walletOrigin_id" integer NOT NULL,
    CONSTRAINT transaction_pkey PRIMARY KEY (id),
    CONSTRAINT "Transaction_walletOrigin_id_fkey" FOREIGN KEY ("walletOrigin_id")
        REFERENCES public.wallet (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Transaction_wallet_id_fkey" FOREIGN KEY ("walletReceive_id")
        REFERENCES public.wallet (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.transaction
    OWNER to postgres;
