-- Mother Table Query
CREATE TABLE IF NOT EXISTS mother (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    firstname TEXT NOT NULL,
    lastname TEXT,
    birth_date TIMESTAMP WITH TIME ZONE,
    email TEXT,
    phone TEXT,
    address TEXT,
    partner_name TEXT,
    image_url TEXT,
    lmp TIMESTAMP WITH TIME ZONE,
    conception_date TIMESTAMP WITH TIME ZONE,
    sono_date TIMESTAMP WITH TIME ZONE,
    crl FLOAT,
    crl_date TIMESTAMP WITH TIME ZONE,
    edd TIMESTAMP WITH TIME ZONE,
    rh_factor TEXT,
    delivered BOOLEAN,
    delivery_date TIMESTAMP WITH TIME ZONE,
    midwife_id INTEGER
    CONSTRIANT fk_midwife_id FOREIGN KEY(midwife_id) REFERENCES midwife(id)
);

-- Midwife Table Query
CREATE TABLE IF NOT EXISTS midwife (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    firstname TEXT,
    lastname TEXT,
    birth_date TIMESTAMP,
    email TEXT,
    pass TEXT,
    image_url TEXT
);
