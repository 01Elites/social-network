import { Tabs } from '@kobalte/core/tabs';
import { useParams } from '@solidjs/router';
import { createEffect, createSignal, For, JSXElement, Show } from 'solid-js';
import FeedPosts from '~/components/Feed/FeedPosts';
import NewPostCell from '~/components/Feed/NewPostCell';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Card } from '~/components/ui/card';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import { cn } from '~/lib/utils';
import Friends from '~/types/friends';
import User from '~/types/User';
import FollowRequest from './followRequest';
import './style.css';
import { Post } from '~/types/Post';

type ProfileParams = {
  username: string;
};

export default function ProfileFeed(props: {
  targetUser: () => User;
}): JSXElement {
  const params: ProfileParams = useParams();
  const [targetFriends, setTargetFriends] = createSignal<Friends | undefined>();
  const [posts, setPosts] = createSignal<Post[]>();
  createEffect(() => {
    // Fetch user Friends
    fetchWithAuth(config.API_URL + '/friends/' + params.username).then(async (res) => {
      const body = await res.json();
      if (res.ok) {
        setTargetFriends(body);
        return;
      } else {
        console.log('Error fetching friends');
        return;
      }
    });
  });
  return (
    <Tabs aria-label='Main navigation' class='tabs'>
      <Tabs.List class='tabs__list'>
        <Tabs.Trigger class='tabs__trigger' value='posts'>
          Posts
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='followers'>
          Followers
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='following'>
          Following
        </Tabs.Trigger>
        <Tabs.Indicator class='tabs__indicator' />
      </Tabs.List>
      <Tabs.Content
        class='tabs__content h-[80vh] overflow-scroll'
        value='posts'
      >
        <div class={cn('flex flex-col gap-4 p-2')}>
          <Show
            when={
              props.targetUser()?.follow_status === 'following' ||
              props.targetUser()?.profile_privacy === 'public' ||
              props.targetUser()?.email !== undefined
            }
            fallback={
              <div class='flex h-full flex-col items-center justify-center'>
                <p class='text-2xl font-bold'>This user's profile is private</p>
                <p class='text-lg'>Follow this user to see their posts</p>
              </div>
            }
          >
            <Show when={props.targetUser().email != undefined}>
              <NewPostCell setPosts={setPosts} />
            </Show>

            <FeedPosts path={`/profile/${params.username}/posts`} posts={posts} setPosts={setPosts} />
          </Show>
        </div>
      </Tabs.Content>
      <Tabs.Content
        class='tabs__content flex flex-wrap gap-4'
        value='followers'
      >
        <Show
          when={
            props.targetUser()?.follow_status === 'following' ||
            props.targetUser()?.profile_privacy === 'public' ||
            props.targetUser()?.email !== undefined
          }
          fallback={
            <div class='flex h-full flex-col items-center justify-center'>
              <p class='text-2xl font-bold'>This user's profile is private</p>
              <p class='text-lg'>Follow this user to see their followers</p>
            </div>
          }
        >
          <For each={targetFriends()?.followers ?? []}>
            {(follower) => (
              <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
                <a
                  href={`/profile/${follower.user_name}`}
                  class='flex flex-col items-center text-base font-bold hover:underline'
                >
                  <Avatar class='mb-3 h-20 w-20'>
                    <AvatarImage src={follower.avatar} />
                    <AvatarFallback>
                      {follower.first_name.charAt(0).toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  {follower.first_name} {follower.last_name}
                </a>
                <FollowRequest
                  username={follower.user_name}
                  status={follower.follow_status}
                  privacy={follower.profile_privacy}
                  profilePage={false}
                />
              </Card>
            )}
          </For>
        </Show>
      </Tabs.Content>
      <Tabs.Content
        class='tabs__content flex flex-wrap gap-4'
        value='following'
      >
        <Show
          when={
            props.targetUser()?.follow_status === 'following' ||
            props.targetUser()?.profile_privacy === 'public' ||
            props.targetUser()?.email !== undefined
          }
          fallback={
            <div class='flex h-full flex-col items-center justify-center'>
              <p class='text-2xl font-bold'>This user's profile is private</p>
              <p class='text-lg'>Follow this user to see their following</p>
            </div>
          }
        >
          <For each={targetFriends()?.following ?? []}>
            {(following) => (
              <Card class='flex w-44 flex-col items-center space-y-4 p-3'>
                <a
                  href={`/profile/${following.user_name}`}
                  class='flex flex-col items-center text-base font-bold  hover:underline'
                >
                  <Avatar class='mb-3 h-20 w-20'>
                    <AvatarImage src={following.avatar} />
                    <AvatarFallback>
                      {following.first_name.charAt(0).toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  {following.first_name} {following.last_name}
                </a>
                <FollowRequest
                  username={following.user_name}
                  status={following.follow_status}
                  privacy={following.profile_privacy}
                  profilePage={false}
                />
              </Card>
            )}
          </For>
        </Show>
      </Tabs.Content>
    </Tabs>
  );
}
