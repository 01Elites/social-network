-- Add ON DELETE CASCADE to post_user table
ALTER TABLE public.post_user
DROP CONSTRAINT IF EXISTS post_user_post_id_fkey,
ADD CONSTRAINT post_user_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.post(post_id) ON DELETE CASCADE;

-- Drop the interaction_type column and enum type
ALTER TABLE public.post_interaction
DROP COLUMN interaction_type;

DROP TYPE public.interaction_type;