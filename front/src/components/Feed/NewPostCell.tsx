import { createSignal, JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/hooks/userDetails';

import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import NewPostPreview from './NewPostPreview';
import config from '~/config';

export default function NewPostCell(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const [postPreviewOpen, setPostPreviewOpen] = createSignal(false);
  return (
    <div class='flex gap-2 rounded border-[1px] p-2'>
      <NewPostPreview setOpen={setPostPreviewOpen} open={postPreviewOpen()} />
      <Avatar>
        <AvatarImage src={`${config.API_URL}/image/${userDetails()?.avatar}`} />
        <AvatarFallback>
          {userDetails()?.first_name.charAt(0).toUpperCase()}
        </AvatarFallback>
      </Avatar>
      <Button
        variant='ghost'
        class='w-full justify-start text-muted-foreground'
        onClick={() => setPostPreviewOpen(true)}
      >
        What's on your mind, {userDetails()?.first_name}?
      </Button>
      {/* <Button
        variant='ghost'
        title='Upload an image Button'
        class='hidden xs:inline-flex'
      >
        <img src={photo} alt='Upload an image icon' />
      </Button> */}
    </div>
  );
}
