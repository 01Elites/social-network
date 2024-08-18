import { createSignal, JSXElement, Setter, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/hooks/userDetails';
import type { Group } from "~/types/group";

import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import NewGroupPostPreview from './NewGroupPostPreview';
import { Post } from '~/types/Post';
interface NewPostCellProps {
  class?: string;
  setPosts: Setter<Post[] | undefined>;
  groupID: string;
}
export default function NewGroupPostCell(props: NewPostCellProps ): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const groupID = props.groupID;

  const [postPreviewOpen, setPostPreviewOpen] = createSignal(false);

  return (
    <div class='flex gap-2 rounded border-[1px] p-2'>
      <NewGroupPostPreview setOpen={setPostPreviewOpen} open={postPreviewOpen()} setPosts={props.setPosts} groupID={groupID} />

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
        What's on your mind, {userDetails()?.first_name}?
      </Button>
    </div>
  );
}
