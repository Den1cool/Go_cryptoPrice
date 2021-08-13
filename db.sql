CREATE DATABASE db
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
    
CREATE TABLE public.pair
(
    id integer NOT NULL DEFAULT nextval('pair_id_seq'::regclass),
    fsyms text COLLATE pg_catalog."default" NOT NULL,
    tsyms text COLLATE pg_catalog."default" NOT NULL,
    raw_id integer NOT NULL,
    display_id integer NOT NULL,
    updatetime timestamp with time zone NOT NULL,
    CONSTRAINT pair_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.pair
    OWNER to postgres;

    CREATE TABLE public."RAW"
(
    raw_id integer NOT NULL DEFAULT nextval('"RAW_raw_id_seq"'::regclass),
    change24hour double precision NOT NULL,
    changepct24hour double precision NOT NULL,
    open24hour double precision NOT NULL,
    volume24hour double precision NOT NULL,
    volume24hourto double precision NOT NULL,
    low24hour double precision NOT NULL,
    high24hour double precision NOT NULL,
    price double precision NOT NULL,
    lastupdate double precision NOT NULL,
    supply double precision NOT NULL,
    mktcap double precision NOT NULL,
    CONSTRAINT "RAW_pkey" PRIMARY KEY (raw_id)
)

TABLESPACE pg_default;

ALTER TABLE public."RAW"
    OWNER to postgres;

    CREATE TABLE public."DISPLAY"
(
    display_id integer NOT NULL DEFAULT nextval('"DISPLAY_display_id_seq"'::regclass),
    change24hour text COLLATE pg_catalog."default" NOT NULL,
    changepct24hour text COLLATE pg_catalog."default" NOT NULL,
    open24hour text COLLATE pg_catalog."default" NOT NULL,
    volume24hour text COLLATE pg_catalog."default" NOT NULL,
    volume24hourto text COLLATE pg_catalog."default" NOT NULL,
    low24hour text COLLATE pg_catalog."default" NOT NULL,
    high24hour text COLLATE pg_catalog."default" NOT NULL,
    price text COLLATE pg_catalog."default" NOT NULL,
    lastupdate text COLLATE pg_catalog."default" NOT NULL,
    supply text COLLATE pg_catalog."default" NOT NULL,
    mktcap text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "DISPLAY_pkey" PRIMARY KEY (display_id)
)

TABLESPACE pg_default;

ALTER TABLE public."DISPLAY"
    OWNER to postgres;
    