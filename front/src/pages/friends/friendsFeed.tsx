import { Tabs } from '@kobalte/core/tabs';
import 'solid-devtools';
import { For, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Card } from '~/components/ui/card';
import Friends from '~/types/friends';
import moment from 'moment';
import { FaSolidCheck } from 'solid-icons/fa';
import { IoClose } from 'solid-icons/io';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Show } from 'solid-js';
import FollowRequest from '../profile/followRequest';
import { useContext } from 'solid-js';
import NotificationsContext from '~/contexts/NotificationsContext';
import { createEffect, createSignal } from 'solid-js';


export default function FriendsFeed(props: {
  targetFriends: () => Friends | undefined;
}): JSXElement {
  const friends = props.targetFriends();
  console.log(friends);
  const [notificationId, setNotificationId] = createSignal<string>('');
  const notifications = useContext(NotificationsContext);
  createEffect(() => {
    if (notificationId() !== '') {
      notifications?.markRead(notificationId(), true);
    }
  })

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
      let data = await res.json();
      setNotificationId(data)
    })
    .catch((err) => {
      console.log('Error responding to request');
    });
  const elem = document.getElementById(follower);
  elem?.remove();
}
  return (
    <Tabs aria-label='Main navigation' class='tabs'>
      <Tabs.List class='tabs__list'>
        <Tabs.Trigger class='tabs__trigger' value='followers'>
          Followers ({friends?.followers?.length || 0})
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='following'>
          Following ({friends?.following?.length || 0})
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='friend_requests'>
          Requests ({friends?.friend_requests?.length || 0})
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='explore'>
          Explore ({friends?.explore?.length || 0})
        </Tabs.Trigger>
        <Tabs.Indicator class='tabs__indicator' />
      </Tabs.List>

      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='followers'>
        <For each={friends?.followers ?? []}>
          {(follower) => (
            <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${follower.user_name}`}
                class='flex flex-col items-center text-base font-bold hover:underline'
              >
                <Avatar class='w-[5rem] h-[5rem] mb-2'>
                  <AvatarFallback>
                    <Show when={follower.avatar} fallback={
                      follower.first_name.charAt(0).toUpperCase()
                    }><img
                        alt='avatar'
                        class='size-full rounded-md rounded-b-none object-cover'
                        loading='lazy'
                        src={`${config.API_URL}/image/${follower.avatar}`}
                      /></Show></AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{follower.first_name}</div>
                  <div>{follower.last_name}</div>
                </div>
              </a>
              <FollowRequest username={follower.user_name} status={follower.follow_status} privacy={follower.profile_privacy} profilePage={false} />

            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='following'>
        <For each={friends?.following ?? []}>
          {(following) => (
            <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${following.user_name}`}
                class='flex flex-col items-center text-base font-bold hover:underline'
              >
                <Avatar class='w-[5rem] h-[5rem] mb-2'>
                  <AvatarFallback>
                    <Show when={following.avatar} fallback={
                      following.first_name.charAt(0).toUpperCase()
                    }><img
                        alt='avatar'
                        class='size-full rounded-md rounded-b-none object-cover'
                        loading='lazy'
                        src={`${config.API_URL}/image/${following.avatar}`}
                      /></Show></AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{following.first_name}</div>
                  <div>{following.last_name}</div>
                </div>
              </a>
              <FollowRequest username={following.user_name} status={following.follow_status} privacy={following.profile_privacy} profilePage={false}/>
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='friend_requests'>
        <For each={friends?.friend_requests ?? []}>
          {(request) => (
            <div id={request.requester} class='flex w-full space-x-1'>
              <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
                <a
                  href={`/profile/${request.user_info.user_name}`}
                  class='flex flex-col items-center text-base font-bold hover:underline'
                >
                  <Avatar class='w-[5rem] h-[5rem] mb-2'>
                    <AvatarFallback>
                      <Show when={request.user_info.avatar} fallback={
                        request.user_info.first_name.charAt(0).toUpperCase()
                      }><img
                          alt='avatar'
                          class='size-full rounded-md rounded-b-none object-cover'
                          loading='lazy'
                          src={`${config.API_URL}/image/${request.user_info.avatar}`}
                        /></Show></AvatarFallback>
                  </Avatar>
                  <div class='flex flex-wrap items-center justify-center space-x-1'>
                    <div>{request.user_info.first_name}</div>
                    <div>{request.user_info.last_name}</div>
                  </div>
                </a>
                <time
                  class='text-xs font-light text-muted-foreground'
                  dateTime={moment(request.creation_date).calendar()}
                  title={moment(request.creation_date).calendar()}
                >
                  {moment(request.creation_date).fromNow()}</time>
                <div class='flex flex-row gap-2'>
                  <Button
                    variant='ghost'
                    class='flex-1 gap-2'
                    onClick={() => { handleFollowRequest("accepted", request.user_info.user_name); }}
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
                    onClick={() => { handleFollowRequest("rejected", request.user_info.user_name) }}
                  >
                    <IoClose class='size-4' color='red' />
                  </Button>
                </div>
              </Card>
            </div>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='m-6 flex flex-wrap' value='explore'>
        <For each={friends?.explore ?? []}>
          {(explore) => (
            <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${explore.user_name}`}
                class='flex flex-col items-center text-base font-bold hover:underline'
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
              <FollowRequest username={explore.user_name} status={explore.follow_status} privacy={explore.profile_privacy} profilePage={false} />
            </Card>
          )}
        </For>
      </Tabs.Content>
    </Tabs>
  );
}

