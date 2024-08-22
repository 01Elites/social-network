-- Down Migration: Remove the 'about' column from 'public.profile'
ALTER TABLE public.profile
DROP COLUMN about;
