import { JSXElement, Show, useContext } from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { Comment } from '~/types/Comment';
import { UserDetailsHook } from '~/types/User';
import TextBreaker from '../core/textbreaker';
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';

type PostCommentCellProps = {
  comment: Comment;
};

export default function PostCommentCell(
  props: PostCommentCellProps,
): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <div class='space-y-4 rounded border-[0.5px] pb-4 shadow-lg'>
      <Show when={props.comment.image}>
        <AspectRatio ratio={16 / 9}>
          <img
            class='size-full rounded-md rounded-b-none object-cover'
            loading='lazy'
            src={`${config.API_URL}/image/${props.comment.image}`}
          />
        </AspectRatio>
      </Show>
      <div
        class={props.comment.image ? 'space-y-4 px-2' : 'space-y-4 px-2 pt-4'}
      >
        <div class='flex items-center justify-between'>
          <PostAuthorCell
            author={props.comment.commenter}
            date={new Date(props.comment.creation_date)}
          />
          {/* <DropdownMenu>
            <DropdownMenuTrigger class={buttonVariants({ variant: 'ghost' })}>
              <IconEllipsis
                class='size-4'
                fill='hsl(var(--muted-foreground))'
              />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem onClick={console.log}>
                <span class='w-full'>Report</span>
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem onClick={console.log}>
                <span class='w-full text-error-foreground'>Delete</span>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu> */}
        </div>
        <p class='text- break-words text-primary/70'>
          <TextBreaker text={props.comment.body} />
        </p>
      </div>
    </div>
  );
}
