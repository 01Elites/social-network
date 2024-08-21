-- Purpose: Update the provider_type
-- Rename the old enum type
ALTER TYPE public.provider_type RENAME TO provider_type_old;
-- Create the new enum type
CREATE TYPE public.provider_type AS ENUM ('google', 'github', 'manual', 'reboot');
-- Update the column using the old enum type
ALTER TABLE public.user
    ALTER COLUMN provider TYPE public.provider_type USING 
    CASE
        WHEN provider = 'password' THEN 'manual'::text
        ELSE provider::text
    END::public.provider_type;
-- Drop the old enum type
DROP TYPE provider_type_old;

-- Purpose: Update the profile table
-- Rename the old enum type
ALTER TYPE public.user_type RENAME TO profile_privacy;
-- Rename the column and change its default value
ALTER TABLE public.profile
    RENAME COLUMN "type" TO privacy;

-- Add the new constraints
ALTER TABLE public.user ADD CONSTRAINT unique_email_provider UNIQUE (email, provider);