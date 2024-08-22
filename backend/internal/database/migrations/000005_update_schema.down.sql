-- Down Migration: Remove the 'event_date' column from 'public.event'
ALTER TABLE public.event
DROP COLUMN event_date;

