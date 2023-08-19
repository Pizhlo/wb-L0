-- UP
-- delivery
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS delivery (
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    zip TEXT NOT NULL,
    city TEXT NOT NULL,
    address TEXT NOT NULL,
    region TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX "delivery_unique_id" ON delivery(id);

COMMIT;

-- payments
BEGIN TRANSACTION;

CREATE TYPE currency_enum AS ENUM ('USD', 'RUB');
CREATE TYPE provider_enum AS ENUM ('wbpay');

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL NOT NULL,
    "transaction" TEXT NOT NULL,
    request_id TEXT,
    currency currency_enum NOT NULL,
    provider provider_enum NOT NULL,
    amount INT NOT NULL,
    payment_date text NOT NULL,
    bank TEXT NOT NULL,
    delivery_cost INT NOT NULL,
    goods_total int NOT NULL,
    custom_fee int NOT NULL,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX "payments_unique_id" ON payments(id);

COMMIT;

-- items
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS items (
    id SERIAL NOT NULL,
    chrt_id INT NOT NULL,
    track_number TEXT NOT NULL,
    price int NOT NULL,
    rid TEXT NOT NULL,
    name TEXT NOT NULL,
    sale INT NOT NULL,
    "size" TEXT NOT NULL,
    total_price INT NOT NULL,
    nm_id INT NOT NULL,
    brand TEXT NOT NULL,
    status INT NOT NULL,
    PRIMARY KEY(id)
);

CREATE UNIQUE INDEX "items_unique_id" ON items(id);
CREATE UNIQUE INDEX "items_unique_track_number" ON items(track_number);

COMMIT;

-- orders
BEGIN TRANSACTION;

CREATE TYPE entry_enum AS ENUM ('WBIL');

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL NOT NULL,
    order_id UUID NOT NULL,
    track_number TEXT NULL,
    entry entry_enum NOT NULL,
    delivery_id serial NOT NULL,
    payment_id serial NOT NULL,
    locale TEXT NOT NULL,
    internal_signature TEXT,
    customer_id TEXT NOT NULL,
    delivery_service TEXT NOT NULL,
    shard_key TEXT NOT NULL,
    sm_id INT NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard TEXT NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (delivery_id) REFERENCES delivery (id),
    FOREIGN KEY (payment_id) REFERENCES payments (id),
    FOREIGN KEY (track_number) REFERENCES items (track_number)
);
COMMIT;