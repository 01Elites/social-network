import { JSXElement, useContext } from 'solid-js';
import { cn } from '~/lib/utils';
import { User } from '../../types/User';
import { Card } from '../../components/ui/card';
import { Avatar, AvatarFallback } from '../../components/ui/avatar';
import config from '../../config';
import { Show, For } from 'solid-js';
import FollowRequest from '../profile/followRequest';
interface HomeContactsProps {
  members: User[] | undefined;
  class: string;
}

export default function GroupContacts(props: HomeContactsProps): JSXElement {
  return (
    <section class={cn('flex flex-wrap', props.class)}>
      <For each={props.members ?? []}>
        {(member) => (
          <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
            <a
              href={`/profile/${member.user_name}`}
              class='flex flex-col items-center text-base font-bold hover:underline'
            >
              <Avatar class='w-[5rem] h-[5rem] mb-2'>
                <AvatarFallback>
                  <Show when={member.avatar} fallback={
                    member.first_name.charAt(0).toUpperCase()
                  }><img
                      alt='avatar'
                      class='size-full rounded-md rounded-b-none object-cover'
                      loading='lazy'
                      src={`${config.API_URL}/image/${member.avatar}`}
                    /></Show></AvatarFallback>
              </Avatar>
              <div class='flex flex-wrap items-center justify-center space-x-1'>
                <div>{member.first_name}</div>
                <div>{member.last_name}</div>
              </div>
            </a>
            <FollowRequest username={member.user_name} status={member.follow_status} privacy={member.profile_privacy} profilePage={false}/>
          </Card>
        )}
      </For>
    </section>
  );
}