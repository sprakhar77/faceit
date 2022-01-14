CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE public.users (
    id SERIAL NOT NULL,
    first_name varchar(20) NOT NULL CHECK (first_name <> ''),
    last_name varchar(20) NOT NULL CHECK (last_name <> ''),
    nickname varchar(20),
    email varchar(40) UNIQUE NOT NULL CHECK (email <> ''),
    password varchar(100) NOT NULL CHECK (password <> ''),
    country varchar(100) NOT NULL CHECK (country <> ''),
    created_at timestamp  NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp  NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON public.users
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();