import { JSXElement } from "solid-js";
import type {Post} from '~/types/Post';
import { showToast } from '~/components/ui/toast';
import { createEffect, createSignal } from 'solid-js';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Tabs } from "@kobalte/core/tabs";
import { Skeleton } from '~/components/ui/skeleton';
import { cn } from "~/lib/utils";
import Repeat from '~/components/core/repeat';
import FeedPosts from '~/components/Feed/FeedPosts';

type GroupPostFeedProps = {
  groupID: string;
}

export default function GroupFeed(props: GroupPostFeedProps): JSXElement {
  console.log(props)
  const [groupPosts, setGroupPosts] = createSignal<Post[]>();
  createEffect(() => {
    fetchWithAuth(config.API_URL + '/group/' + props.groupID + '/posts').then(async (res) => {
      const body = await res.json();
      if (res.status === 404) {
        console.log('User not found');
        return;
      }
      if (res.ok) {
        setGroupPosts(body);
        return;
      }
      throw new Error(
        body.reason ? body.reason : 'An error occurred while fetching posts',
      );
    })
    .catch((err) => {
      showToast({
        title: 'Error fetching posts',
        description: err.message,
        variant: 'error',
      });
    });
  })
  return (
      <Tabs aria-label="Main navigation" class="tabs">
      <Tabs.List class="tabs__list">
        <Tabs.Trigger class="tabs__trigger" value="posts">Posts</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill">Chat</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill2">Events</Tabs.Trigger>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="posts">
        <div class={cn('flex flex-col gap-4 p-2')}>
        <div class={cn('flex flex-col gap-4 p-2')}>
          <FeedPosts path={`/group/${props.groupID}/posts`} />
        </div>
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill2">still NOTHING!!!</Tabs.Content>
    </Tabs>);
}