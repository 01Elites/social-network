-- Add ON DELETE CASCADE to post_user table
ALTER TABLE public.post_user
DROP CONSTRAINT IF EXISTS post_user_post_id_fkey,
ADD CONSTRAINT post_user_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.post(post_id) ON DELETE CASCADE;

-- Add ON DELETE CASCADE to comment table for post_id
ALTER TABLE public.comment
DROP CONSTRAINT IF EXISTS comment_post_id_fkey,
ADD CONSTRAINT comment_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.post(post_id) ON DELETE CASCADE;
