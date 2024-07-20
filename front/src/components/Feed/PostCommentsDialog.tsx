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
import { Comment } from '~/types/Comment';
import { Post } from '~/types/Post';
import { UserDetailsHook } from '~/types/User';
import Repeat from '../core/repeat';
import { Button } from '../ui/button';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '../ui/dialog';
import { Separator } from '../ui/separator';
import { TextField, TextFieldTextArea } from '../ui/text-field';
import { showToast } from '../ui/toast';
import FeedPostCellSkeleton from './FeedPostCellSkeleton';
import PostCommentCell from './PostCommentCell';

const [post, setPost] = createSignal<Post>();
let newCommentCallback: () => void;

export function showComments(post: Post, _newCommentCallback: () => void) {
  setPost(post);
  newCommentCallback = _newCommentCallback;
}

export function PostCommentsDialog(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const [open, setOpen] = createSignal(false);
  const [comment, setComment] = createSignal('');
  const [uploadedImage, setUploadedImage] = createSignal<File | null>(null);

  const [commentPosting, setCommentPosting] = createSignal(false);
  const [postComments, setPostComments] = createSignal<Comment[]>();

  function close() {
    setOpen(false);
    setPost(undefined);
    setComment('');
    setUploadedImage(null);
    newCommentCallback = () => {};
  }

  function fetchPostComments() {
    fetchWithAuth(`${config.API_URL}/post/${post()?.post_id}/comments`)
      .then(async (res) => {
        if (!res.ok) {
          const body = await res.json();
          throw new Error(body.reason || 'Failed to fetch comments');
        }
        const comments = await res.json();
        setPostComments(comments.comments);
      })
      .catch((err) => {
        showToast({
          title: 'Failed to fetch comments',
          description: err.message,
          variant: 'error',
        });
      });
  }

  createEffect(() => {
    setOpen(post() !== undefined);
    // don't show comments skeleton if there are no comments
    if (post()) {
      if (post()?.comments_count === 0) {
        setPostComments([]);
      }

      fetchPostComments();
    }
  });

  function handleImageUpload(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      setUploadedImage(target.files[0]);
    }
  }

  async function postComment(e: SubmitEvent) {
    setCommentPosting(true);
    e.preventDefault();

    const payload = {
      body: comment(),
      image: '',
    };

    if (uploadedImage()) {
      try {
        const base64 = await uploadedImage()?.toBase64();
        payload.image = base64 as string;
      } catch (error) {
        console.error('Error converting image to base64:', error);
      }
    }

    fetchWithAuth(`${config.API_URL}/post/${post()?.post_id}/comments`, {
      method: 'POST',
      body: JSON.stringify(payload),
    })
      .then(async (res) => {
        setCommentPosting(false);
        setComment('');
        setUploadedImage(null);

        if (!res.ok) {
          const body = await res.json();
          throw new Error(body.reason || 'Failed to post comment');
        }
        newCommentCallback();
        fetchPostComments();
      })
      .catch((err) => {
        showToast({
          title: 'Failed to post comment',
          description: err.message,
          variant: 'error',
        });
      });
  }

  return (
    <Dialog open={open()} onOpenChange={close}>
      <Show when={post()}>
        <DialogContent class='h-full max-h-[70%] overflow-hidden'>
          <DialogHeader>
            <DialogTitle>
              Comments{' '}
              <span class='text-primary/80'>
                ({postComments()?.length || post()!.comments_count})
              </span>
            </DialogTitle>
          </DialogHeader>
          <div class='flex max-h-full flex-col gap-2 overflow-hidden'>
            <div class='flex flex-1 flex-col gap-4 overflow-y-scroll'>
              <Show when={postComments() === undefined}>
                <Repeat count={10}>
                  <FeedPostCellSkeleton />
                </Repeat>
              </Show>

              <Show
                when={post()!.comments_count > 0}
                fallback={
                  <h1 class='text-center text-primary/60'>No comments yet</h1>
                }
              >
                <For each={postComments()}>
                  {(comment) => <PostCommentCell comment={comment} />}
                </For>
              </Show>
            </div>

            <DialogFooter class='flex !flex-col gap-4'>
              <Separator />
              <form class='!m-1 flex !flex-col gap-4' onSubmit={postComment}>
                <TextField
                  onChange={setComment}
                  value={comment()}
                  disabled={!userDetails() || commentPosting()}
                  name='comment'
                >
                  <TextFieldTextArea
                    class='resize-none'
                    placeholder="That's so boring"
                  ></TextFieldTextArea>
                </TextField>

                <div class='flex w-full flex-col justify-between gap-4 sm:flex-row'>
                  <input
                    class='hidden'
                    type='file'
                    id='commentImageUpload'
                    accept='image/*'
                    onChange={handleImageUpload}
                  />
                  <Button
                    class='w-full self-end sm:w-fit'
                    variant='secondary'
                    onClick={() =>
                      document.getElementById('commentImageUpload')?.click()
                    }
                    disabled={!userDetails() || commentPosting()}
                  >
                    upload image
                  </Button>

                  <Button
                    type='submit'
                    class='w-full self-end sm:w-fit'
                    disabled={
                      !userDetails() || comment().length < 1 || commentPosting()
                    }
                  >
                    Comment
                  </Button>
                </div>
              </form>
            </DialogFooter>
          </div>
        </DialogContent>
      </Show>
    </Dialog>
  );
}
