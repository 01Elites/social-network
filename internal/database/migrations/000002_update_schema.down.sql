-- Purpose: Down migration the provider_type.
-- Rename the old enum type
ALTER TYPE public.provider_type RENAME TO provider_type_new;
-- Create the original enum type
CREATE TYPE public.provider_type AS ENUM ('google', 'github', 'password', 'reboot');
-- Update the column using the new enum type
ALTER TABLE public.user
    ALTER COLUMN provider TYPE public.provider_type USING 
    CASE
        WHEN provider = 'manual' THEN 'password'::text
        ELSE provider::text
    END::public.provider_type;
-- Drop the new enum type
DROP TYPE provider_type_new;

-- Purpose: Down migration the profile table.
-- Rename the enum type back to the original
ALTER TYPE public.profile_privacy RENAME TO user_type;
-- Rename the column and change its default value back
ALTER TABLE public.profile
    RENAME COLUMN privacy TO "type";

-- Remove the new constraints
ALTER TABLE public.user DROP CONSTRAINT unique_email_provider;

