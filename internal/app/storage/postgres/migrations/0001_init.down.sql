BEGIN TRANSACTION;

DROP INDEX  IF EXISTS "items_unique_id";
DROP INDEX  IF EXISTS "payments_unique_id";
DROP INDEX  IF EXISTS "delivery_unique_id";
DROP INDEX IF EXISTS "items_unique_track_number";

DROP TYPE IF EXISTS "currency_enum";
DROP TYPE IF EXISTS "provider_enum";
DROP TYPE IF EXISTS "entry_enum";

DROP TABLE IF EXISTS delivery CASCADE;
DROP TABLE IF EXISTS items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS payments CASCADE;

COMMIT;