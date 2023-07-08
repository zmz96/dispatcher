
CREATE TABLE IF NOT EXISTS public.drivers
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    phone text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default",
    plate text COLLATE pg_catalog."default",
    car text COLLATE pg_catalog."default",
    CONSTRAINT drivers_pkey PRIMARY KEY (id)
)

ALTER TABLE IF EXISTS public.drivers
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS public.passengers
(
    id text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default",
    phone text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default",
    CONSTRAINT passengers_pkey PRIMARY KEY (id)
)

ALTER TABLE IF EXISTS public.passengers
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS public.rides
(
    uuid uuid DEFAULT gen_random_uuid(),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    passenger_id text COLLATE pg_catalog."default",
    driver_id text COLLATE pg_catalog."default",
    lat numeric,
    lon numeric,
    addr text COLLATE pg_catalog."default",
    accepted boolean,
    CONSTRAINT fk_rides_driver FOREIGN KEY (driver_id)
        REFERENCES public.drivers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_rides_passenger FOREIGN KEY (passenger_id)
        REFERENCES public.passengers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

ALTER TABLE IF EXISTS public.rides
    OWNER to postgres;