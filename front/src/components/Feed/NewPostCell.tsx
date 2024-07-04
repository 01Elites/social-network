import { JSXElement, createSignal, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

import photo from '~/assets/photo.svg';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';

export default function NewPostCell(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const [postText, setPostText] = createSignal('');

  return (
    <div class='flex gap-2 rounded border-[1px] p-2'>
      <Avatar>
        <AvatarImage src={userDetails()?.avatar_url} />
        <AvatarFallback>
          {userDetails()?.first_name.charAt(0).toUpperCase()}
        </AvatarFallback>
      </Avatar>

      <TextField class='grow' onChange={setPostText}>
        <TextFieldInput
          class='border-none'
          type='text'
          placeholder='Your father disappointed?'
        />
      </TextField>
      {postText().trim().length > 0 && (
        <Button class='animate-content-show3' title='Post Button'>
          Post
        </Button>
      )}
      <Button variant='ghost' title='Upload an image Button'>
        <img src={photo} alt='Upload an image icon' />
      </Button>
    </div>
  );
}
