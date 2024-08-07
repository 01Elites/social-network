import { Tabs } from '@kobalte/core/tabs';
import 'solid-devtools';
import { For, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Card } from '~/components/ui/card';
import moment from 'moment';
import { FaSolidCheck } from 'solid-icons/fa';
import { IoClose } from 'solid-icons/io';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Show } from 'solid-js';
import FollowRequest from '../profile/followRequest';
import { Notifications } from '~/types/notifications';
import {createSignal} from 'solid-js'
import { EventsFeed } from '../events/eventsfeed';
import { handleInvite } from '../group/request';
import { A } from '@solidjs/router';

export default function NotificationsFeed(): JSXElement {
  const [notifications, setnotification] = createSignal<Notifications>();

  return (<>
        <Show when={notifications()?.payload.type == "FOLLOW_REQUEST"}>
          <div id={notifications()?.payload.metadata.requester.user_name}>
            <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
                <a
                  href={`/profile/${notifications()?.payload.metdata.requester.user_name}`}
                  class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
                >
                  <Avatar class='w-[5rem] h-[5rem] mb-2'>
                    <AvatarFallback>
                      <Show when={notifications()?.payload.metadata.requester.avatar} fallback={
                        notifications()?.payload.metdata.requester.first_name.charAt(0).toUpperCase()
                      }><img
                          alt='avatar'
                          class='size-full rounded-md rounded-b-none object-cover'
                          loading='lazy'
                          src={`${config.API_URL}/image/${notifications()?.payload.metadata.requester.avatar}`}
                        /></Show></AvatarFallback>
                  </Avatar>
                  <div class='flex flex-wrap items-center justify-center space-x-1'>
                    <div>{notifications()?.payload.metadata.requester.first_name}</div>
                    <div>{notifications()?.payload.metadata.requester.last_name}</div>
                  </div>
                </a>
                <time
                  class='text-xs font-light text-muted-foreground'
                  dateTime={moment(notifications()?.payload.metadata.creation_date).calendar()}
                  title={moment(notifications()?.payload.metadata.creation_date).calendar()}
                >
                  {moment(notifications()?.payload.metadata.creation_date).fromNow()}</time>
                <div class='flex flex-row gap-2'>
                  <Button
                    variant='ghost'
                    class='flex-1 gap-2'
                    onClick={() => { handleFollowRequest("accepted", notifications()?.payload.metadata.requester.user_name); }}
                  >
                    <FaSolidCheck
                      class='size-4'
                      color='green'
                    />
                  </Button>
                  <Button
                    variant='ghost'
                    class='flex-1 gap-2'
                    color="red"
                    onClick={() => { handleFollowRequest("rejected", notifications()?.payload.metdata.requester.user_name) }}
                  >
                    <IoClose class='size-4' color='red' />
                  </Button>
                </div>
              </Card>
              </div>
              </Show>
              <Show when={notifications()?.payload.type == "GROUP_INVITATION"}>
              <div id={notifications()?.payload.metdata.requester.invited_by.user.user_name}>
      <Card class='flex flex-col items-center space-y-4 p-3'>
      <p class="flex-col justify-center items-center">
            {<A
        href={"/profile/" + notifications()?.payload.metdata.requester.invited_by.user.user_name} class='flex flex-col justify-center items-center'>
    {notifications()?.payload.metdata.requester.invited_by.user.first_name}  {notifications()?.payload.metdata.requester.invited_by.user.last_name}</A> }
    invited you to join {notifications()?.payload.metadata.group.title}<br></br>
    <time
    class='text-xs font-light text-muted-foreground'
    dateTime={moment(notifications()?.payload.metdata.requester.invited_by.creation_date).calendar()}
    title={moment(notifications()?.payload.metdata.requester.invited_by.creation_date).calendar()}
  >
    {moment(notifications()?.payload.metdata.requester.invited_by.creation_date).fromNow()}</time>
    </p>
    <div class='flex flex-row gap-2'>
      <Button
      variant='ghost'
      class='flex-1 gap-2'
      onClick={() => {handleInvite("accepted", notifications()?.payload.metdata.group.id, notifications()?.payload.metdata.requester.invited_by.user.user_name);}}
    >
      <FaSolidCheck
       class='size-4'
       color='green'
      />
    </Button>
    <Button
      variant='ghost'
      class='flex-1 gap-2'
      color="red"
      onClick={() => {handleInvite("rejected", notifications()?.payload.metdata.group.id, notifications()?.payload.metdata.requester.invited_by.user.user_name)}}
    >
    <IoClose class='size-4' color='red'/>
    </Button>
    </div>
  </Card></div>      </Show>
              </>)
      {/* <Tabs.Content class='m-6 flex flex-wrap gap-4' value='invites'> */}
        {/* <For each={notifications?.GroupInvite ?? []}>
          {(invite) => (
            <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${invite.invited_by.user.user_name}`}
                class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
              >
                <Avatar class='w-[5rem] h-[5rem] mb-2'>
                  <AvatarFallback>
                    <Show when={invite.invited_by.user.avatar} fallback={
                      invite.invited_by.user.first_name.charAt(0).toUpperCase()
                    }><img
                        alt='avatar'
                        class='size-full rounded-md rounded-b-none object-cover'
                        loading='lazy'
                        src={`${config.API_URL}/image/${invite.invited_by.user.avatar}`}
                      /></Show></AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{invite.invited_by.user.first_name}</div>
                  <div>{invite.invited_by.user.last_name}</div>
                </div>
              </a>
              <FollowRequest username={invite.invited_by.user.user_name} status={invite.invited_by.user.follow_status} privacy={invite.invited_by.user.profile_privacy} />
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='events'>
          <EventsFeed events={notifications?.Events} />
      </Tabs.Content> */}

      {/* <Tabs.Content class='m-6 flex flex-wrap' value='explore'>
        <For each={friends?.explore ?? []}>
          {(explore) => (
            <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${explore.user_name}`}
                class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
              >
                <Avatar class='w-[5rem] h-[5rem] mb-2'>
                  <AvatarFallback>
                    <Show when={explore.avatar} fallback={
                      explore.first_name.charAt(0).toUpperCase()
                    }><img
                        alt='avatar'
                        class='size-full rounded-md rounded-b-none object-cover'
                        loading='lazy'
                        src={`${config.API_URL}/image/${explore.avatar}`}
                      /></Show></AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{explore.first_name}</div>
                  <div>{explore.last_name}</div>
                </div>
              </a>
              <FollowRequest username={explore.user_name} status={explore.follow_status} privacy={explore.profile_privacy} />
            </Card>
          )}
        </For>
      </Tabs.Content> */}
}

function handleFollowRequest(response: string, follower: string) {
  fetchWithAuth(`${config.API_URL}/follow_response`, {
    method: 'POST',
    body: JSON.stringify({
      follower: follower,
      status: response,
    })
  })
    .then(async (res) => {
      if (!res.ok) {
        throw new Error(
          // reason ?? 'An error occurred while responding to request',
        );
      }
    })
    .catch((err) => {
      console.log('Error responding to request');
    });
  const elem = document.getElementById(follower);
  elem?.remove();
}
