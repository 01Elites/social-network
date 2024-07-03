import { Image } from '@kobalte/core/image';
import { JSXElement, createSignal, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

import photo from '~/assets/photo.svg';
import { Button } from '../ui/button';

export default function NewPostCell(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const [postText, setPostText] = createSignal('');

  return (
    <div class='border-[1px] rounded p-2 flex gap-2'>
      <Image class='bg-muted size-10 flex rounded-full justify-center items-center'>
        <Image.Img
          src={userDetails()?.avatar_url}
          class='rounded-full'
          alt={userDetails()?.first_name + ' Image'}
        />
        <Image.Fallback>E</Image.Fallback>
      </Image>
      <TextField class='grow' onChange={setPostText}>
        <TextFieldInput
          class='border-none'
          type='text'
          placeholder='Your father disappointed?'
        />
      </TextField>
      {postText().trim().length > 10 && (
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
