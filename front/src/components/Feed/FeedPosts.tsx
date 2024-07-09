import { JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

interface FeedPostsProps {}

export default function FeedPosts(props: FeedPostsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <div class='flex flex-col gap-4'>
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
    </div>
  );
}
