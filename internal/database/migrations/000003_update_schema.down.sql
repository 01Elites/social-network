-- Add user_name column back to user table
ALTER TABLE public.user ADD COLUMN user_name VARCHAR(100) UNIQUE;

-- Make sure nick_name is unique before transfer
-- We need a temporary column to assign unique row numbers
ALTER TABLE public.profile ADD COLUMN temp_unique_nick_name VARCHAR(100);

-- Use a Common Table Expression (CTE) to calculate unique nick_name
WITH RankedProfiles AS (
    SELECT user_id,
           nick_name,
           nick_name || '_' || ROW_NUMBER() OVER (PARTITION BY nick_name ORDER BY user_id) AS unique_nick_name
    FROM public.profile
)
-- Update the profile table with the unique nick_name from the CTE
UPDATE public.profile
SET temp_unique_nick_name = rp.unique_nick_name
FROM RankedProfiles rp
WHERE public.profile.user_id = rp.user_id;


-- Transfer the updated nick_name from profile to user
UPDATE public.user AS usr
SET user_name = prf.temp_unique_nick_name
FROM public.profile AS prf
WHERE usr.user_id = prf.user_id;

-- Drop the temporary column from profile table
ALTER TABLE public.profile DROP COLUMN temp_unique_nick_name;

-- Drop the nick_name column from profile table
ALTER TABLE public.profile DROP COLUMN nick_name;

-- Optionally, you might want to set the user_name as unique again if that is required
-- ALTER TABLE public.user ADD CONSTRAINT user_name_unique UNIQUE (user_name);
