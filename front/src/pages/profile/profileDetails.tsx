import { JSXElement, Show, useContext } from 'solid-js';
import {
  ProfileEditDialog,
  showEditProfile,
} from '~/components/EditProfileDialog';
import { AspectRatio } from '~/components/ui/aspect-ratio';
import { Avatar, AvatarFallback } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import Follow_Icon from '~/components/ui/icons/follow_icon';
import Globe_Icon from '~/components/ui/icons/globe_icon';
import Message_Icon from '~/components/ui/icons/message_icon';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import User, { UserDetailsHook } from '~/types/User';
import FollowRequest from './followRequest';

export default function ProfileDetails(props: {
  targetUser: () => User;
}): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  console.log(props.targetUser());
  return (
    <div class='flex flex-col'> {/* Left div */}
      <div class='flex flex-col justify-center items-center'>
        <AspectRatio ratio={16 / 9}> {/* Profile picture */}
          <div class='absolute inset-0 bg-black bg-opacity-50 flex justify-center items-end rounded-lg'>
            <Avatar class='w-[5rem] h-[5rem] mb-2'>
              <AvatarFallback>{props.targetUser().avatar}</AvatarFallback>
            </Avatar>
          </div>
        </AspectRatio> {/* Profile picture */}
        <div class='flex flex-col items-center w-full'> {/* Username, followers, following */}
          <p class='text-2xl font-bold m-2'>{props.targetUser().first_name} {props.targetUser().last_name}</p>
          <div class='grid w-full grid-cols-2 text-sm m-2'>
            <p class='flex justify-center'>Followers {props.targetUser().follower_count}</p>
            <p class='flex justify-center'>Followring {props.targetUser().following_count}</p>
          </div>
        </div> {/* Username, followers, following */}
        <div class='m-4'> {/* Bio */}
          <p>{props.targetUser().about}</p>
        </div> {/* Bio */}
        <Show
          when={userDetails()?.user_name === props.targetUser().user_name}
          fallback={
            <div class='m-4 flex w-full flex-row justify-between gap-2'>
              {/* <Button class='flex grow' variant='default'>
                <Follow_Icon />
                {props.targetUser().follow_status === 'following'
                  ? 'Unfollow'
                  : 'Follow'}
              </Button> */}
              <div class='flex gap-2'>
                {/* Follow button */}
                <FollowRequest username={props.targetUser().user_name} status={props.targetUser().follow_status}/>
                <Button variant='default'>
                  <Globe_Icon class='w-5 justify-center' />
                </Button>
                <Button variant='default'>
                  <Message_Icon class='w-5 justify-center' />
                </Button>
              </div>
              {/* Follow button */}
            </div>
          }
        >
          <Button
            class='m-4 w-full'
            variant='default'
            onClick={showEditProfile}
          >
            Edit Profile
          </Button>
          <ProfileEditDialog />
        </Show>
      </div>
    </div>
  );
}
