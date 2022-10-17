CREATE TABLE public.links
(
    id bigint NOT NULL DEFAULT nextval('id_sequence'::regclass),
    link character varying(255) COLLATE pg_catalog."default" NOT NULL,
    shortlink character varying(10) COLLATE pg_catalog."default" NOT NULL,
    domain bigint NOT NULL DEFAULT 0,
    CONSTRAINT links_pkey PRIMARY KEY (id),
    CONSTRAINT id_domain FOREIGN KEY (domain)
        REFERENCES public.domains (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)