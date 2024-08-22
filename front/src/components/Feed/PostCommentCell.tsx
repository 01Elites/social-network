import { createEffect, JSXElement, Show, useContext } from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { Comment } from '~/types/Comment';
import { UserDetailsHook } from '~/hooks/userDetails';
import TextBreaker from '../core/textbreaker';
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';
import { fetchWithAuth } from '~/extensions/fetch';

type PostCommentCellProps = {
  comment: Comment;
};

export default function PostCommentCell(
  props: PostCommentCellProps,
): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  let imageRef!: HTMLImageElement;

  createEffect(() => {
    if (!props.comment.image) {
      return;
    }
    fetchWithAuth(`${config.API_URL}/image/${props.comment.image}`).then(
      async (res) => {
        if (!res.ok) {
          const body = await res.json();
          throw new Error(
            body.reason ?? 'An error occurred while fetching the image',
          );
        }
        const blob = await res.blob();
        const url = URL.createObjectURL(blob);
        imageRef.src = url;
      },
    );
  });

  return (
    <div class='space-y-4 rounded border-[0.5px] pb-4 shadow-lg'>
      <Show when={props.comment.image}>
        <AspectRatio ratio={16 / 9}>
          <img
            ref={imageRef}
            alt='comment'
            class='size-full rounded-md rounded-b-none object-cover'
            loading='lazy'
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
