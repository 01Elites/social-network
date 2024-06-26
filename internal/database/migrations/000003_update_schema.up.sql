-- Add nick_name column to profile table
ALTER TABLE public.profile ADD COLUMN nick_name VARCHAR(100);

-- Copy user_name data from user to profile table
UPDATE public.profile SET nick_name = usr.user_name
FROM public.user AS usr
WHERE public.profile.user_id = usr.user_id;

-- Remove the user_name column from user table
ALTER TABLE public.user DROP COLUMN user_name;