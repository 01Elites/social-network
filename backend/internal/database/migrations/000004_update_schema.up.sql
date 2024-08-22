-- Up Migration: Add the 'about' column to 'public.profile'
ALTER TABLE public.profile
ADD COLUMN about TEXT;
