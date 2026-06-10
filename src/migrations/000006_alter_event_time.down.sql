ALTER TABLE events 
  ALTER COLUMN "time" TYPE TIME WITHOUT TIME ZONE 
  USING ("time" AT TIME ZONE 'Asia/Kolkata')::time;