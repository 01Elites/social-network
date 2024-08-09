import { JSXElement, Show, useContext } from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { Post } from '~/types/Post';
import { UserDetailsHook } from '~/hooks/userDetails';
import TextBreaker from '../core/textbreaker';
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';
import { Button, buttonVariants } from '../ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '../ui/dropdown-menu';
import IconComments from '../ui/icons/IconComments';
import IconEllipsis from '../ui/icons/IconEllipsis';
import IconThumb from '../ui/icons/IconThumb';
import { showToast } from '../ui/toast';
import { showComments } from './PostCommentsDialog';

interface FeedPostCellProps {
  post: Post;
  updatePost: (updatedPost: Post) => void;
}

export default function FeedPostCell(props: FeedPostCellProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  function isLiked() {
    return props.post.likers_usernames?.includes(
      userDetails()?.user_name ?? '',
    );
  }

  function updateCommentsCount() {
    props.updatePost({
      ...props.post,
      comments_count: props.post.comments_count + 1,
    });
  }

  function handleLike() {
    fetchWithAuth(`${config.API_URL}/post/${props.post.post_id}/like`, {
      method: 'POST',
    })
      .then(async (res) => {
        if (!res.ok) {
          const body = await res.json();
          throw new Error(
            body.reason ?? 'An error occurred while (un)liking the post',
          );
        }
      })
      .catch((err) => {
        showToast({
          title: 'Error (un)liking post',
          description: err.message,
          variant: 'error',
        });
      });

    if (isLiked()) {
      props.updatePost({
        ...props.post,
        likers_usernames: props.post.likers_usernames?.remove(
          userDetails()?.user_name!,
        ),
      });
    } else {
      const currentUser = userDetails()?.user_name!;
      const newLikers = props.post.likers_usernames
        ? [...props.post.likers_usernames, currentUser]
        : [currentUser];

      props.updatePost({
        ...props.post,
        likers_usernames: newLikers,
      });
    }
  }

  return (
    <div class='space-y-4 overflow-hidden rounded border-[0.5px] pb-4 shadow-lg'>
      <Show when={props.post.image}>
        <AspectRatio ratio={16 / 9}>
          <img
            alt='post image'
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
          </DropdownMenu>
        </div>
        <p class='text- break-words text-primary/70'>
          <TextBreaker text={props.post.content} />
        </p>

        <div class='flex gap-2 xs:block'>
          <Button
            variant='ghost'
            class='flex-1 gap-2'
            disabled={!userDetails()}
            onClick={handleLike}
          >
            <IconThumb
              class='size-4'
              variant={isLiked() ? 'solid' : 'outline'}
            />
            {props.post.likers_usernames?.length ?? 0}
          </Button>

          <Button
            variant='ghost'
            class='flex-1 gap-2'
            onClick={() => showComments(props.post, updateCommentsCount)}
          >
            <IconComments class='size-4' />
            {props.post.comments_count}
          </Button>
        </div>
      </div>
    </div>
  );
}
