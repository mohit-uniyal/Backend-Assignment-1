ALTER TABLE reservations
    ALTER COLUMN reservation_id DROP DEFAULT,
    ALTER COLUMN reservation_id SET DATA TYPE UUID USING gen_random_uuid(),
    ALTER COLUMN reservation_id SET DEFAULT gen_random_uuid();