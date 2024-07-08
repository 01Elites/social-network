-- Up Migration: Add the 'about' column to 'public.profile'
ALTER TABLE public.event
ADD COLUMN event_date TIMESTAMP NOT NULL;

ALTER TABLE public.event
ALTER COLUMN description SET NOT NULL;
