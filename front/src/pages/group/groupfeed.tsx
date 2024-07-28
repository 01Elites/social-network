import { JSXElement } from "solid-js";
import type {Post} from '~/types/Post';
import { showToast } from '~/components/ui/toast';
import { createEffect, createSignal } from 'solid-js';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Tabs } from "@kobalte/core/tabs";
import { cn } from "~/lib/utils";
import FeedPosts from '~/components/Feed/FeedPosts';
import { Show } from "solid-js";
import User  from "~/types/User";
import { GroupRequests } from "~/pages/group/creatorsrequest";

type GroupPostFeedProps = {
  groupID: string;
  creator: boolean;
  requesters: requester[] | undefined;
}

export type requester = {
  user: User;
  creation_date: string;
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
        <Tabs.Trigger class="tabs__trigger" value="chat">Chat</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="events">Events</Tabs.Trigger>
        <Show when={props.creator}>
        <Tabs.Trigger class="tabs__trigger" value="requests">Requests</Tabs.Trigger>
        </Show>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="posts">
        <div class={cn('flex flex-col gap-4 p-2')}>
        <div class={cn('flex flex-col gap-4 p-2')}>
          <FeedPosts path={`/group/${props.groupID}/posts`} />
        </div>
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="chat">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="events">still NOTHING!!!</Tabs.Content>
      <Show when={props.creator}>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="requests">
        <Show when={props.requesters?.length === 0}>
          <h1 class='text-center font-bold text-muted-foreground'>
            Hmmm, we don't seem to have any requests
          </h1>
        </Show>
        <GroupRequests requesters={props.requesters} groupID={props.groupID}/>
      </Tabs.Content>
      </Show>
    </Tabs>);
}
