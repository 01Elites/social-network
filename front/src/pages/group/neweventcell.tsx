import { createSignal, JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import NewPostPreview from '../../components/Feed/NewPostPreview';
import CreateEvent from './createevent';

interface eventParams {
  groupTitle: string | undefined;
  groupID: string;
}

export default function NewEventCell(props: eventParams): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const [postPreviewOpen, setPostPreviewOpen] = createSignal(false);

  return (
    <div class='flex gap-2 rounded border-[1px] p-2'>
      <CreateEvent setOpen={setPostPreviewOpen} open={postPreviewOpen()} groupTitle={props.groupTitle} groupID={props.groupID}/>
      <Avatar>
        <AvatarImage src={userDetails()?.avatar} />
        <AvatarFallback>
          {userDetails()?.first_name.charAt(0).toUpperCase()}
        </AvatarFallback>
      </Avatar>
      <Button
        variant='ghost'
        class='w-full justify-start text-muted-foreground'
        onClick={() => setPostPreviewOpen(true)}
      >
        Create New Event
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
