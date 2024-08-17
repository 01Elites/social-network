import {
  createEffect,
  createSignal,
  For,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { cn } from '~/lib/utils';
import { Post } from '~/types/Post';
import { UserDetailsHook } from '~/hooks/userDetails';
import Repeat from '../core/repeat';
import { showToast } from '../ui/toast';
import FeedPostCell from './FeedPostCell';
import FeedPostCellSkeleton from './FeedPostCellSkeleton';
import { PostCommentsDialog } from './PostCommentsDialog';

interface FeedPostsProps {
  class?: string;
  path: string;
}

export default function FeedPosts(props: FeedPostsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [posts, setPosts] = createSignal<Post[]>();

  function updatePost(updatedPost: Post) {
    const updatedPosts = posts()?.map((post) =>
      post.post_id === updatedPost.post_id ? updatedPost : post,
    );
    setPosts(updatedPosts);
  }
  createEffect(() => {
    if (!userDetails()) {
      setPosts(null as any);
      return;
    };

    fetchWithAuth(config.API_URL + props.path)
      .then(async (res) => {
        const body = await res.json();
        if (res.status === 404) {
          setPosts([]);
          return;
        }
        if (res.ok) {
          setPosts(body);
          return;
        }
        throw new Error(
          body.reason ? body.reason : 'An error occurred while fetching posts',
        );
      })
      .catch((err) => {
        showToast({
          title: 'Error fetching posts',
          description: err.message,
          variant: 'error',
        });
      });
  });

  return (
    <div class={cn('flex flex-col gap-4', props.class)}>
      <PostCommentsDialog />
      <Show
        when={posts()}
        fallback={
          <Repeat count={10}>
            <FeedPostCellSkeleton />
          </Repeat>
        }
      >
        <Show when={posts()?.length === 0}>
          <h1 class='text-center font-bold text-muted-foreground'>
            Hmmm, we don't seem to have any posts :(
          </h1>
          <p class='text-center text-muted-foreground'>
            Maybe you could post some
          </p>
        </Show>
        <For each={posts()}>
          {(post) => <FeedPostCell post={post} updatePost={updatePost} />}
        </For>
      </Show>
    </div>
  );
}
