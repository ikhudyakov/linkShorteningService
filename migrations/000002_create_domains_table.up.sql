CREATE TABLE domains
(
    id bigint NOT NULL,
    domain character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT domains_pkey PRIMARY KEY (id)
);

