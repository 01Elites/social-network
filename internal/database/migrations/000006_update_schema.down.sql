-- Remove ON DELETE CASCADE
ALTER TABLE public.post_user
DROP CONSTRAINT IF EXISTS post_user_post_id_fkey,
ADD CONSTRAINT post_user_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.post(post_id);

-- Recreate the interaction_type enum type and add the column back
CREATE TYPE public.interaction_type AS ENUM ('like', 'dislike');

ALTER TABLE public.post_interaction
ADD COLUMN interaction_type public.interaction_type NOT NULL;