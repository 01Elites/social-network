import {
  createEffect,
  createSignal,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import {
  ProfileEditDialog,
  showEditProfile,
} from '~/components/EditProfileDialog';
import { AspectRatio } from '~/components/ui/aspect-ratio';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import Message_Icon from '~/components/ui/icons/message_icon';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/hooks/userDetails';
import getRandomColor from '~/lib/randomColors';
import User from '~/types/User';
import FollowRequest from './followRequest';

export default function ProfileDetails(props: {
  targetUser: () => User;
}): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [imageURL, setImageURL] = createSignal<string | undefined>();

  // load image with auth headers
  createEffect(() => {
    if (!props.targetUser().avatar) {
      return;
    }
    fetchWithAuth(`${config.API_URL}/image/${props.targetUser().avatar}`).then(
      async (res) => {
        if (!res.ok) {
          const body = await res.json();
          throw new Error(
            body.reason ?? 'An error occurred while fetching the image',
          );
        }
        const blob = await res.blob();
        const url = URL.createObjectURL(blob);
        setImageURL(url);
      },
    );
  });

  return (
    <div class='flex flex-col'>
      {/* Left div */}
      <div class='flex flex-col items-center justify-center'>
        <AspectRatio ratio={16 / 9}>
          <div
            class={`absolute inset-0 flex items-end justify-center rounded-lg`}
            style={{
              'background-color': getRandomColor(),
            }}
          ></div>
        </AspectRatio>
        <Avatar class='bottom-10 z-10 size-[7rem] text-[2rem]'>
          <AvatarImage src={imageURL()} alt='avatar' />
          <AvatarFallback>
            {props.targetUser().first_name.charAt(0).toUpperCase()}
            {props.targetUser().last_name.charAt(0).toUpperCase()}
          </AvatarFallback>
        </Avatar>
        <div class='flex w-full flex-col items-center'>
          {/* Username, followers, following */}
          <p class='m-2 text-2xl font-bold'>
            {props.targetUser().first_name} {props.targetUser().last_name}{' '}
          </p>
          <div class='m-2 grid w-full grid-cols-2 text-sm'>
            <p class='flex justify-center'>
              Followers {props.targetUser().follower_count}
            </p>
            <p class='flex justify-center'>
              Following {props.targetUser().following_count}
            </p>
          </div>
          <p class='flex justify-center text-sm text-muted-foreground'>
            Account status: {props.targetUser().profile_privacy}
          </p>
          <Show
            when={userDetails()?.user_name === props.targetUser().user_name}
          >
            <p class='flex justify-center text-sm text-muted-foreground'>
              Email: {props.targetUser().email}
            </p>
            <p class='flex justify-center text-sm text-muted-foreground'>
              Username: {props.targetUser().user_name}
            </p>
            <p class='flex justify-center text-sm text-muted-foreground'>
              Nickname: {props.targetUser().nick_name}
            </p>
          </Show>
        </div>{' '}
        {/* Username, followers, following */}
        <div class='m-4'>
          {' '}
          {/* Bio */}
          <p>{props.targetUser().about}</p>
        </div>{' '}
        {/* Bio */}
        <Show
          when={userDetails()?.user_name === props.targetUser().user_name}
          fallback={
            <div class='m-4 flex w-full flex-row justify-center gap-2'>
              <div class='flex w-full gap-2'>
                {/* Follow button */}
                <FollowRequest
                  username={props.targetUser().user_name}
                  status={props.targetUser().follow_status}
                  privacy={props.targetUser().profile_privacy}
                  profilePage={true}
                />
                <Button disabled={true} variant='default'>
                  <Message_Icon darkBack={true} class='w-5 justify-center' />
                </Button>
              </div>
            </div>
          }
        >
          <div>
            <Button
              class='m-4 w-full'
              variant='default'
              onClick={showEditProfile}
            >
              Edit Profile
            </Button>
            <ProfileEditDialog />
          </div>
        </Show>
      </div>
    </div>
  );
}
