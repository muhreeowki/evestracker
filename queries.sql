CREATE TABLE IF NOT EXISTS mother (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    firstname TEXT,
    lastname TEXT,
    birth_date TIMESTAMP,
    email TEXT,
    phone TEXT,
    address TEXT,
    partner_name TEXT,
    image_url TEXT,
    lmp TIMESTAMP,
    conception_date TIMESTAMP,
    sono_date TIMESTAMP,
    crl FLOAT,
    crl_date TIMESTAMP,
    edd TIMESTAMP,
    rh_factor TEXT,
    delivered BOOLEAN
);

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
