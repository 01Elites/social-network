import { Tabs } from '@kobalte/core/tabs';
import { createEffect, createSignal, Index, JSXElement, Show } from 'solid-js';
import { showToast } from '~/components/ui/toast';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { cn } from "~/lib/utils";
import { Post } from "~/types/Post";
import FeedPosts from '~/components/Feed/FeedPosts';
import { GroupRequests } from "~/pages/group/creatorsrequest";
import { Card } from '~/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { For } from 'solid-js';
import { User } from '~/types/User';
import { Button } from '~/components/ui/button';

type GroupPostFeedProps = {
  groupID: string;
  creator: boolean;
  requesters: requester[] | undefined;
  explore: User[] | undefined;
};

export type requester = {
  user: User;
  creation_date: string;
};
export default function GroupFeed(props: GroupPostFeedProps): JSXElement {
  var [buttonData, setButtonData] = createSignal(["", ""]);
  function sendRequestApi(username: string) {
    if (buttonData() === null) {
      return
    }
    fetchWithAuth(config.API_URL + "/invitation", {
      method: 'POST',
      body: JSON.stringify({
        receiver: username,
        group_id: props.groupID
      })
    }).then(async (res) => {
      if (res.status === 200) {
        // setButtonData([username, "Invite Pending"])
        setButtonData((prev) => ({ ...prev, [username]: "Invite Pending" }));
        return;
      } else {
        console.log(res.body);
        console.log('Error making request');
        return
      }
    })
  }
  return (
    <Tabs aria-label='Main navigation' class='tabs'>
      <Tabs.List class='tabs__list'>
        <Tabs.Trigger class='tabs__trigger' value='posts'>
          Posts
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='chat'>
          Chat
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='events'>
          Events
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='invite'>
          Invite
        </Tabs.Trigger>
        <Show when={props.creator}>
          <Tabs.Trigger class='tabs__trigger' value='requests'>
            Requests
          </Tabs.Trigger>
        </Show>
        <Tabs.Indicator class='tabs__indicator' />
      </Tabs.List>
      <Tabs.Content
        class='tabs__content h-[80vh] overflow-scroll'
        value='posts'
      >
        <div class={cn('flex flex-col gap-4 p-2')}>
          <div class={cn('flex flex-col gap-4 p-2')}>
            <FeedPosts path={`/group/${props.groupID}/posts`} />
          </div>
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="chat">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="events">still NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content m-6 flex flex-wrap gap-4 " value="invite">
        <Index each={props?.explore ?? []}>
          {(explore, i) => <>
            <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${explore().user_name}`}
                class='flex flex-col items-center text-base font-bold text-blue-500'
              >
                <Avatar class='mb-3 h-20 w-20'>
                  <AvatarImage src={explore().avatar} />
                  <AvatarFallback>
                    {explore().first_name.charAt(0).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div class='flex flex-wrap items-center justify-center space-x-1'>
                  <div>{explore().first_name}</div>
                  <div>{explore().last_name}</div>
                </div>
              </a>
              <div class="flex flex-row gap-2">
                <Show when={buttonData()[explore().user_name] !== "Invite Pending"}
                  fallback={buttonData()[explore().user_name]}>
                  <Button
                    class="flex grow"
                    variant="default"
                    onClick={() => sendRequestApi(explore().user_name)}
                  >
                    {buttonData()[explore().user_name] || "Invite"}
                  </Button>
                </Show>
              </div>
            </Card>
          </>}
        </Index>
      </Tabs.Content>
      <Show when={props.creator}>
        <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="requests">
          <Show when={props.requesters?.length === 0}>
            <h1 class='text-center font-bold text-muted-foreground'>
              Hmmm, we don't seem to have any requests
            </h1>
          </Show>
          <GroupRequests requesters={props.requesters} groupID={props.groupID} />
        </Tabs.Content>
      </Show>
    </Tabs>
  );
}
