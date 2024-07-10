import { JSXElement, Show } from 'solid-js';
import config from '~/config';
import { Post } from '~/types/Post';

interface FeedPostCellProps {
  post: Post;
}

export default function FeedPostCell(props: FeedPostCellProps): JSXElement {
  console.log(props.post);

  return (
    <div class='rounded-lg p-4 shadow-lg'>
      <p>{props.post.content}</p>
      <Show when={props.post.image}>
        <img src={`${config.API_URL}/image/${props.post.image}`} />
      </Show>
    </div>
  );
}
