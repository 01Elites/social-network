import { JSXElement, Show } from 'solid-js';
import config from '~/config';
import { Post } from '~/types/Post';
import TextBreaker from '../core/textbreaker';
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';

interface FeedPostCellProps {
  post: Post;
}

export default function FeedPostCell(props: FeedPostCellProps): JSXElement {
  console.log(props.post);

  return (
    <div class='space-y-4 overflow-hidden rounded-md border-[0.5px] pb-4 shadow-lg'>
      <Show when={props.post.image}>
        <AspectRatio ratio={16 / 9}>
          <img
            class='size-full rounded-md rounded-b-none object-cover'
            loading='lazy'
            src={`${config.API_URL}/image/${props.post.image}`}
          />
        </AspectRatio>
      </Show>
      <div class={props.post.image ? 'space-y-4 px-4' : 'space-y-4 px-4 pt-4'}>
        <PostAuthorCell
          author={props.post.poster}
          date={new Date(props.post.creation_date)}
        />
        <p class='break-words'>
          <TextBreaker text={props.post.content} />
        </p>
      </div>
    </div>
  );
}
