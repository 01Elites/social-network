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
import { handleRequest } from '../group/creatorsrequest';

export default function FriendsFeed(props: {
  targetFriends: () => Friends | undefined;
}): JSXElement {
  const friends = props.targetFriends();
  console.log(friends);

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
                class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
              >
                <Avatar class='mb-3 h-20 w-20'>
                  <AvatarImage src={follower.avatar} />
                  <AvatarFallback>
                    {follower.first_name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                {follower.first_name} {follower.last_name}
              </a>
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
                class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
              >
                <Avatar class='mb-3 h-20 w-20'>
                  <AvatarImage src={following.avatar} />
                  <AvatarFallback>
                    {following.first_name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                {following.first_name} {following.last_name}
              </a>
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
                  href={`/profile/${request.requester}`}
                  class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
                >
                  <Avatar class='mb-3 h-20 w-20'>
                    <AvatarImage src={request.user_info.avatar} />
                    <AvatarFallback>
                      {request.user_info.first_name.charAt(0).toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  {request.user_info.first_name} {request.user_info.last_name}
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
                class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
              >
                <Avatar class='mb-3 h-20 w-20'>
                  <AvatarImage src={explore.avatar} />
                  <AvatarFallback>
                    {explore.first_name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{explore.first_name}</div>
                  <div>{explore.last_name}</div>
                </div>
              </a>
            </Card>
          )}
        </For>
      </Tabs.Content>
    </Tabs>
  );
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
