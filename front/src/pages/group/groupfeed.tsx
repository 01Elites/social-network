import { Tabs } from '@kobalte/core/tabs';
import { createSignal, Index, JSXElement, Show } from 'solid-js';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { cn } from "~/lib/utils";
import FeedPosts from '~/components/Feed/FeedPosts';
import { GroupRequests } from "~/pages/group/creatorsrequest";
import { Card } from '~/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { User } from '~/types/User';
import { Button } from '~/components/ui/button';
import NewEventCell from './neweventcell';
import GroupContacts from "./groupcontacts";
import GroupChatPage from '~/components/GroupChat';

type GroupPostFeedProps = {
  groupID: string;
  groupTitle: string | undefined;
  creator: boolean;
  requesters: requester[] | undefined;
  explore: User[] | undefined;
  members: User[] | undefined;
};
export type GroupChatState = {
  isOpen: boolean; // Whether the chat window is open
  chatWith: string; // The recipient's username
};
export type requester = {
  user: User;
  creation_date: string;
};
export default function GroupFeed(props: GroupPostFeedProps): JSXElement {
  var [buttonData, setButtonData] = createSignal<{ [key: string]: string }>({});
  const [groupChatState, setGroupChatState] = createSignal<GroupChatState>({
    isOpen: true,
    chatWith: props.groupID
  });

  function sendRequestApi(username: string) {
    if (buttonData() === null) {
      return
    }
    fetchWithAuth(config.API_URL + "/invitation", {
      method: 'POST',
      body: JSON.stringify({
        receiver: username,
        group_id: Number(props.groupID)
      })
    }).then(async (res) => {
      if (res.status === 200) {
        // setButtonData([username, "Invite Pending"])
        setButtonData((prev) => ({ ...prev, [username]: "Invite Pending" }));
        return;
      } else {
        console.log('Error making request');
        return
      }
    })

    // Open the group chat after sending the request
    setGroupChatState({
      isOpen: true,
      chatWith: props.groupID
    });
  }

  return (
    <Tabs aria-label='Main navigation' class='tabs'>
      <Tabs.List class='tabs__list'>
        <div class='tabs__list__indicator' >
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
          <Tabs.Trigger class='tabs__trigger' value='members'>
            Members
          </Tabs.Trigger>
          <Show when={props.creator}>
            <Tabs.Trigger class='tabs__trigger' value='requests'>
              Requests
            </Tabs.Trigger>
          </Show>
        </div>
        <Tabs.Indicator class='tabs__indicator' />
      </Tabs.List>

      <Tabs.Content
        class='tabs__content h-[80vh] overflow-y-scroll'
        value='posts'
      >
        <div class={cn('flex flex-col gap-4 p-2')}>
          <div class={cn('flex flex-col gap-4 p-2')}>
            <FeedPosts path={`/group/${props.groupID}/posts`} />
          </div>
        </div>
      </Tabs.Content>


      <Tabs.Content class="tabs__content overflow-y-scroll h-[80vh]" value="chat">
      <GroupChatPage class='grow place-content-end overflow-hidden' chatState={groupChatState()} setChatState={setGroupChatState} />
      </Tabs.Content>

      <Tabs.Content class="tabs__content overflow-y-scroll h-[80vh]" value="events">
        <NewEventCell groupTitle={props.groupTitle} groupID={props.groupID} />
      </Tabs.Content>

      <Tabs.Content class="tabs__content overflow-y-scroll flex flex-wrap " value="invite">
        <Index each={props?.explore ?? []}>
          {(explore, i) => <>
            <Card class='m-2 flex w-44 flex-col items-center space-y-4 p-3'>
              <a
                href={`/profile/${explore().user_name}`}
                class='flex flex-col items-center text-base font-bold text-blue-500'
              >
                <Avatar class='w-[5rem] h-[5rem] mb-2'>
                  <AvatarFallback>
                    <Show when={explore().avatar} fallback={
                      explore().first_name.charAt(0).toUpperCase()
                    }><img
                        alt='avatar'
                        class='size-full rounded-md rounded-b-none object-cover'
                        loading='lazy'
                        src={`${config.API_URL}/image/${explore().avatar}`}
                      /></Show></AvatarFallback>
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
        <Tabs.Content class="tabs__content overflow-y-scroll h-[80vh]" value="requests">
          <Show when={props.requesters?.length === 0}>
            <h1 class='text-center font-bold text-muted-foreground'>
              Hmmm, we don't seem to have any requests
            </h1>
          </Show>
          <GroupRequests requesters={props.requesters} groupID={props.groupID} />
        </Tabs.Content>
      </Show>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="members">
        <GroupContacts members={props.members} class=' flex flex-wrap' />
      </Tabs.Content>
    </Tabs>
  );
}
