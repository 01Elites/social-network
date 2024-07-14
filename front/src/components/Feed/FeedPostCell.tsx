import { JSXElement, Show } from 'solid-js';
import config from '~/config';
import { Post } from '~/types/Post';
import TextBreaker from '../core/textbreaker';
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';
import { DropdownMenu, DropdownMenuTrigger } from '../ui/dropdown-menu';

interface FeedPostCellProps {
  post: Post;
}

export default function FeedPostCell(props: FeedPostCellProps): JSXElement {
  return (
    <div class='space-y-4 overflow-hidden rounded border-[0.5px] pb-4 shadow-lg'>
      <Show when={props.post.image}>
        <AspectRatio ratio={16 / 9}>
          <img
            class='size-full rounded-md rounded-b-none object-cover'
            loading='lazy'
            src={`${config.API_URL}/image/${props.post.image}`}
          />
        </AspectRatio>
      </Show>
      <div class={props.post.image ? 'space-y-4 px-2' : 'space-y-4 px-2 pt-4'}>
        <div class='flex items-center justify-between'>
          <PostAuthorCell
            author={props.post.poster}
            date={new Date(props.post.creation_date)}
          />
          <DropdownMenu>
            <DropdownMenuTrigger>{}</DropdownMenuTrigger>
          </DropdownMenu>
        </div>
        <p class='break-words'>
          <TextBreaker text={props.post.content} />
        </p>
      </div>
    </div>
  );
}
