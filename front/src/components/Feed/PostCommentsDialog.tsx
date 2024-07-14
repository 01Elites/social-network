import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { Post } from '~/types/Post';
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

const [post, setPost] = createSignal<Post>();

export function showComments(post: Post) {
  setPost(post);
}

export function PostCommentsDialog(): JSXElement {
  const [open, setOpen] = createSignal(false);
  const [comment, setComment] = createSignal('');

  function close() {
    setOpen(false);
    setPost(undefined);
  }

  createEffect(() => {
    setOpen(post() !== undefined);
  });

  return (
    <Dialog open={open()} onOpenChange={close}>
      <Show when={post()}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>
              Comments{' '}
              <span class='text-primary/80'>({post()!.comments_count})</span>
            </DialogTitle>
          </DialogHeader>
          <DialogFooter class='flex !flex-col gap-4'>
            <Separator />

            <TextField onChange={setComment}>
              <TextFieldTextArea
                class='resize-none'
                placeholder="That's so boring"
              ></TextFieldTextArea>
            </TextField>

            <Button
              class='w-full self-end sm:w-fit'
              disabled={comment().length < 1}
            >
              Comment
            </Button>
          </DialogFooter>
        </DialogContent>
      </Show>
    </Dialog>
  );
}
