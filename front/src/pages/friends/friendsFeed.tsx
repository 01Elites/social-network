import { Tabs } from '@kobalte/core/tabs';
import 'solid-devtools';
import { For, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Card } from '~/components/ui/card';
import Friends from '~/types/friends';

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
                class='flex flex-col items-center text-base font-bold text-blue-500'
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
                class='flex flex-col items-center text-base font-bold text-blue-500'
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
            <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${request.requester}`}
                class='flex flex-col items-center text-base font-bold text-blue-500'
              >
                <Avatar class='mb-3 h-20 w-20'>
                  <AvatarImage src={request.user_info.avatar} />
                  <AvatarFallback>
                    {request.user_info.first_name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                {request.user_info.first_name} {request.user_info.last_name}
              </a>
              <div class='flex w-full space-x-1'>
                <Button class='flex-1' variant='default'>
                  Accept
                </Button>
                <Button class='flex-1' variant='default'>
                  Reject
                </Button>
              </div>
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='m-6 flex flex-wrap' value='explore'>
        <For each={friends?.explore ?? []}>
          {(explore) => (
            <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${explore.user_name}`}
                class='flex flex-col items-center text-base font-bold text-blue-500'
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
