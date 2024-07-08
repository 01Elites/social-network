-- Up Migration: Add the 'event_date' column to 'public.event'
ALTER TABLE public.event
ADD COLUMN event_date TIMESTAMP NOT NULL;
