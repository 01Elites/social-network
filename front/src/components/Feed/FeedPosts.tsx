import { createSignal, JSXElement, Show, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { Post } from '~/types/Post';
import { UserDetailsHook } from '~/types/User';
import { Skeleton } from '../ui/skeleton';

interface FeedPostsProps {}

export default function FeedPosts(props: FeedPostsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [posts, setPosts] = createSignal<Post[]>();

  return (
    <div class='flex flex-col gap-4'>
      <Show
        when={posts()}
        fallback={
          <>
            {Array.from({ length: 10 }).map((_, i) => (
              <div class='space-y-4'>
                <div class='flex items-center space-x-4'>
                  <Skeleton height={40} circle animate={false} />
                  <div class='w-full space-y-2'>
                    <Skeleton height={16} radius={10} class='max-w-40' />
                    <Skeleton height={16} radius={10} class='max-w-32' />
                  </div>
                </div>
                <Skeleton height={150} radius={10} />
              </div>
            ))}
          </>
        }
      >
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
        <h1>asdasd</h1>
      </Show>
    </div>
  );
}
