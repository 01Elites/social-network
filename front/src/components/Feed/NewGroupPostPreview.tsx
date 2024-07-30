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
import PostAuthorCell from '../PostAuthorCell';
import { AspectRatio } from '../ui/aspect-ratio';
import { TextField, TextFieldTextArea } from '../ui/text-field';
import NewPostPrivacy from './NewPostPrivacy';

import tailspin from '~/assets/svg-loaders/tail-spin.svg';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import { showToast } from '../ui/toast';
import { Tooltip, TooltipContent, TooltipTrigger } from '../ui/tooltip';

interface NewPostPreviewProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  groupID: int ;
}

type NewPostPrivacyOptions = 'public' | 'private' | 'almost_private';

export default function NewGroupPostPreview(props: NewPostPreviewProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [uploadedImage, setUploadedImage] = createSignal<File | null>(null);
  const [postText, setPostText] = createSignal<string>('');

  const [postPrivacyOpen, setPostPrivacyOpen] = createSignal(false);
  const [postPrivacy, setPostPrivacy] =
    createSignal<NewPostPrivacyOptions>('public');
  const [selectedUsers, setSelectedUsers] = createSignal<String[]>([]);

  const [formProcessing, setFormProcessing] = createSignal(false);

  async function makePost() {
    setFormProcessing(true);

    const payload = {
      title: '',
      body: postText(),
      image: '',
      privacy: postPrivacy(),
      usernames: selectedUsers(),
      group_id: props.groupID,
    };

    if (uploadedImage()) {
      try {
        const base64 = await uploadedImage()?.toBase64();
        payload.image = base64 as string;
      } catch (error) {
        console.error('Error converting image to base64:', error);
      }
    }

    fetchWithAuth(config.API_URL + '/post', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
      .then(async (response) => {
        setFormProcessing(false);
        if (!response.ok) {
          const errMsg = await response.json();
          showToast({
            title: 'Could not post',
            description: errMsg.reason
              ? errMsg.reason
              : 'An error occurred while posting your content',
            variant: 'error',
          });
        } else {
          props.setOpen(false);
        }
      })
      .catch((error) => {
        setFormProcessing(false);
        console.error('Error posting:', error);
        showToast({
          title: 'Could not post',
          description: 'An error occurred while posting your content',
          variant: 'error',
        });
      });
  }

  // Reset uploaded image when dialog is closed
  createEffect(() => {
    if (props.open) {
      setUploadedImage(null);
      setPostText('');
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
      <NewPostPrivacy
        onlyFollowersCallback={() => setPostPrivacy('private')}
        onlySelectedCallback={(selectedUsers) => {
          setPostPrivacy('almost_private');
          setSelectedUsers(selectedUsers);
        }}
        open={postPrivacyOpen()}
        setOpen={setPostPrivacyOpen}
      />
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create New Post</DialogTitle>
          <DialogDescription>
            Share your thoughts with the world. Make sure you are respectful
            and kind to others.
          </DialogDescription>
        </DialogHeader>

        <AspectRatio ratio={16 / 9} class='rounded bg-muted'>
          {uploadedImage() ? (
            <>
              <Button
                class='absolute right-2 top-2 h-6 rounded-full px-2 py-2 text-xs'
                variant='secondary'
                onClick={() => setUploadedImage(null)}
                disabled={formProcessing()}
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
                placeholder='Upload an image'
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

        <TextField onChange={setPostText} value={postText()}>
          <TextFieldTextArea
            placeholder='What do you want to say? you have 255 characters to express yourself.'
            class='resize-none'
            minLength={1}
            maxLength={400}
            disabled={formProcessing()}
          />
        </TextField>

        <Separator />
        <DialogFooter class='!justify-between gap-4'>
          <Tooltip>
            <TooltipTrigger
              as={Button<'button'>}
              variant='secondary'
              disabled={formProcessing()}
              onClick={() => setPostPrivacyOpen(true)}
            >
              {postPrivacy() === 'public' && 'Post Privacy'}
              {postPrivacy() === 'private' && 'Only My Followers'}
              {postPrivacy() === 'almost_private' && 'Only Selected Users'}
            </TooltipTrigger>
            <TooltipContent>
              <p>
                Who do you wnat to see your post? we show it to everyone be
                default.
              </p>
            </TooltipContent>
          </Tooltip>

          <Button
            class='gap-2'
            disabled={postText().length < 1 || formProcessing()}
            onClick={makePost}
          >
            {formProcessing() && <img src={tailspin} class='h-full' alt='processing' />}
            {formProcessing() ? 'Posting...' : 'Post'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
