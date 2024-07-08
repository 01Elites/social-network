import { createEffect, createSignal, JSXElement, useContext } from 'solid-js';
import { Button } from '~/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';
import { Separator } from '~/components/ui/separator';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import User, { UserDetailsHook } from '~/types/User';
import { AspectRatio } from '../ui/aspect-ratio';
import { TextField, TextFieldTextArea } from '../ui/text-field';
import PostAuthorCell from './PostAuthorCell';

interface NewPostPreviewProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

export default function NewPostPreview(props: NewPostPreviewProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [uploadedImage, setUploadedImage] = createSignal<File | null>(null);

  // Reset uploaded image when dialog is closed
  createEffect(() => {
    if (props.open) {
      setUploadedImage(null);
    }
  });

  function handleImageUpload(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      setUploadedImage(target.files[0]);
    }
  }

  return (
    <Dialog open={props.open} onOpenChange={props.setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Apparently he is Disappointed</DialogTitle>
          <DialogDescription>
            It's okay <b>clown</b> 🤡 don't cry yet. Wait till you read the
            comments
          </DialogDescription>
        </DialogHeader>

        <AspectRatio ratio={16 / 9} class='rounded bg-muted'>
          {uploadedImage() ? (
            <>
              <Button
                class='absolute right-2 top-2 h-6 rounded-full px-2 py-2 text-xs'
                variant='secondary'
                onClick={() => setUploadedImage(null)}
              >
                X
              </Button>
              <img
                class='size-full rounded-md object-cover'
                src={URL.createObjectURL(uploadedImage() as File)}
                alt='selected image'
              />
            </>
          ) : (
            <>
              <input
                class='hidden'
                type='file'
                id='postImageUpload'
                accept='image/*'
                onChange={handleImageUpload}
              />
              <Button
                variant={'secondary'}
                class='h-full w-full flex-col'
                onClick={() =>
                  document.getElementById('postImageUpload')?.click()
                }
              >
                Upload an image
                <p class='font-light text-muted-foreground'>
                  make sure your image is 16:9 ratio
                </p>
              </Button>
            </>
          )}
        </AspectRatio>

        <PostAuthorCell author={userDetails() as User} date={new Date()} />

        <TextField>
          <TextFieldTextArea
            placeholder='What do you want to say?'
            class='resize-none'
          />
        </TextField>

        <Separator />
        <DialogFooter class='!justify-between gap-4'>
          <Button variant={'secondary'}>Post Privacy</Button>
          <Button>Post</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
