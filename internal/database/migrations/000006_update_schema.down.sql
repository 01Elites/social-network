-- Remove ON DELETE CASCADE
ALTER TABLE public.post_user
DROP CONSTRAINT IF EXISTS post_user_post_id_fkey,
ADD CONSTRAINT post_user_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.post(post_id);