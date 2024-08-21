-- Down Migration: V2__remove_engineer_from_gender_type.sql

-- Step 1: Create a new enum type without the 'engineer' value
CREATE TYPE public.gender_type_new AS ENUM ('male', 'female');

-- Step 2: Update the 'gender' column in the 'profile' table to use the new enum type
ALTER TABLE public.profile
    ALTER COLUMN gender TYPE public.gender_type_new 
    USING gender::text::public.gender_type_new;

-- Step 3: Drop the old enum type
DROP TYPE public.gender_type;

-- Step 4: Rename the new enum type to the original name
ALTER TYPE public.gender_type_new RENAME TO gender_type;
