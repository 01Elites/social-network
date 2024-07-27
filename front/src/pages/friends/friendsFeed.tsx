import { Tabs } from '@kobalte/core/tabs';
import 'solid-devtools';
import { For, JSXElement } from 'solid-js';
import { Button } from '~/components/ui/button';
import Follow_Icon from '~/components/ui/icons/follow_icon';
import Friends from '~/types/friends';
import { Card } from '~/components/ui/card';

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
      <Tabs.Content class='tabs__content h-[80vh]' value='followers'>
        <For each={friends?.followers ?? []}>
          {(follower) => (
            <Card>
              <a href={`/profile/${follower}`} class='text-blue-500 underline'>
                {follower}
              </a>
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class='tabs__content h-[80vh]' value='following'>
        <For each={friends?.following ?? []}>
          {(following) => (
            <Card>
              <a href={`/profile/${following}`} class='text-blue-500 underline'>
                {following}
              </a>
            </Card>
          )}
        </For>
      </Tabs.Content>
      <Tabs.Content class='tabs__content h-[80vh]' value='friend_requests'>
        <For each={friends?.friend_requests ?? []}>
          {(request) => (
            <Card>
              <a
                href={`/profile/${request.requester}`}
                class='text-blue-500 underline'
              >
                {request.requester}
              </a>
              <Button class='flex grow' variant='default'>
                Accept
              </Button>
              <Button class='flex grow' variant='default'>
                Reject
              </Button>
            </Card>
          )}
        </For>
      </Tabs.Content>
      <Tabs.Content class='tabs__content h-[80vh]' value='explore'>
        <For each={friends?.explore ?? []}>
          {(explore) => (
            <Card>
              <a href={`/profile/${explore}`} class='text-blue-500 underline'>
                {explore}
              </a>
              <Button class='flex grow' variant='default'>
                <Follow_Icon />
                Follow
              </Button>
            </Card>
          )}
        </For>
      </Tabs.Content>
    </Tabs>
  );
}
