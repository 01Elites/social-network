import { JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import { UserDetailsHook } from '~/types/User';
import NewPostCell from './NewPostCell';
import PostsCards from './PostsCards';

interface FeedProps {
  class?: string;
}

export default function Feed(props: FeedProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Feed</h1>
      {userDetails() && <NewPostCell />}
      <PostsCards />
    </section>
  );
}
