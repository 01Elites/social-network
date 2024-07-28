import { Tabs } from '@kobalte/core/tabs';
import { A } from '@solidjs/router';
import moment from 'moment';
import { FaSolidCheck } from 'solid-icons/fa';
import { IoClose } from 'solid-icons/io';
import { createEffect, createSignal, Index, JSXElement, Show } from 'solid-js';
import FeedPosts from '~/components/Feed/FeedPosts';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { showToast } from '~/components/ui/toast';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import { cn } from '~/lib/utils';
import type { Post } from '~/types/Post';
import User from '~/types/User';

type GroupPostFeedProps = {
  groupID: string;
  creator: boolean;
  requesters: requester[] | undefined;
};

export type requester = {
  user: User;
  creation_date: string;
};
export default function GroupFeed(props: GroupPostFeedProps): JSXElement {
  console.log(props);
  const [requester, removeRequester] = createSignal<requester[]>();
  const [groupPosts, setGroupPosts] = createSignal<Post[]>();
  createEffect(() => {
    fetchWithAuth(config.API_URL + '/group/' + props.groupID + '/posts')
      .then(async (res) => {
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
  });
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
      <Tabs.Content class='tabs__content' value='nill'>
        NOTHING!!!
      </Tabs.Content>
      <Tabs.Content class='tabs__content' value='nill2'>
        still NOTHING!!!
      </Tabs.Content>
      <Show when={props.creator}>
        <Tabs.Content
          class='tabs__content h-[80vh] overflow-scroll'
          value='requests'
        >
          <Show when={props.requesters?.length === 0}>
            <h1 class='text-center font-bold text-muted-foreground'>
              Hmmm, we don't seem to have any requests
            </h1>
          </Show>
          <Index each={props.requesters}>
            {(requester, i) => (
              <>
                {' '}
                <div class='flex gap-2 xs:block'>
                  <div id={requester().user.user_name}>
                    <p class='flex-1 gap-2'>
                      <Avatar>
                        <AvatarImage src={requester().user.avatar} />
                        <AvatarFallback>
                          {requester().user.first_name.charAt(0).toUpperCase()}
                        </AvatarFallback>
                      </Avatar>
                      <A
                        href={'/profile/' + requester().user.user_name}
                        class='block text-sm font-bold hover:underline'
                      >
                        {requester().user.first_name}{' '}
                        {requester().user.last_name}
                      </A>
                      <time
                        class='text-xs font-light text-muted-foreground'
                        dateTime={moment(requester().creation_date).calendar()}
                        title={moment(requester().creation_date).calendar()}
                      >
                        {moment(requester().creation_date).fromNow()}
                      </time>
                    </p>
                    <Button
                      variant='ghost'
                      class='flex-1 gap-2'
                      onClick={() => {
                        handleRequest(
                          'accepted',
                          props.groupID,
                          requester().user.user_name,
                        );
                        props.requesters?.splice(i, i + 1);
                      }}
                    >
                      <FaSolidCheck class='size-4' color='green' />
                    </Button>
                    <Button
                      variant='ghost'
                      class='flex-1 gap-2'
                      color='red'
                      onClick={() => {
                        handleRequest(
                          'rejected',
                          props.groupID,
                          requester().user.user_name,
                        );
                        props.requesters?.splice(i, i + 1);
                      }}
                    >
                      <IoClose class='size-4' color='red' />
                    </Button>
                  </div>
                </div>
              </>
            )}
          </Index>
        </Tabs.Content>
      </Show>
    </Tabs>
  );
}

export function handleRequest(
  response: string,
  groupID: string,
  requester: string,
) {
  fetchWithAuth(`${config.API_URL}/join_group_res`, {
    method: 'PATCH',
    body: JSON.stringify({
      requester: requester,
      group_id: groupID,
      response: response,
    }),
  })
    .then(async (res) => {
      if (!res.ok) {
        throw new Error();
        // reason ?? 'An error occurred while responding to request',
      }
    })
    .catch((err) => {
      showToast({
        title: 'Error responding to request',
        description: err.message,
        variant: 'error',
      });
    });
  const elem = document.getElementById(requester);
  elem?.remove();
}
