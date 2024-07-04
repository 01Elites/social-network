import { JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

import photo from '~/assets/photo.svg';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';

export default function NewPostCell(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <div class='flex gap-2 rounded border-[1px] p-2'>
      {/* <NewPostPreview /> */}

      <Avatar>
        <AvatarImage src={userDetails()?.avatar_url} />
        <AvatarFallback>
          {userDetails()?.first_name.charAt(0).toUpperCase()}
        </AvatarFallback>
      </Avatar>
      <Button variant='ghost' class='w-full justify-start'>
        Your father disappointed?
      </Button>
      <Button
        variant='ghost'
        title='Upload an image Button'
        class='hidden xs:inline-flex'
      >
        <img src={photo} alt='Upload an image icon' />
      </Button>
    </div>
  );
}
