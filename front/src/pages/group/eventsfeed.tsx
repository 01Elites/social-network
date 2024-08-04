import {
  createEffect,
  createSignal,
  For,
  Index,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/types/User';
import { showToast } from '../../components/ui/toast';
import { GroupEvent } from '~/types/group/index';
import moment from 'moment';
import { Card } from '~/components/ui/card';
import { Button } from '~/components/ui/button';
import Tooltip from '@corvu/tooltip'

interface FeedPostsProps {
  groupID: string;
}

export default function EventsFeed(props: FeedPostsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [events, setEvents] = createSignal<GroupEvent[]>();

  // function updatePost(updatedPost: Post) {
  //   const updatedPosts = events()?.map((event) =>
  //     post.post_id === updatedPost.post_id ? updatedPost : post,
  //   );
  //   setPosts(updatedPosts);
  // }
  createEffect(() => {
    if (!userDetails()) return;
    fetchWithAuth(config.API_URL + `/group/${props.groupID}/events`)
      .then(async (res) => {
        const body = await res.json();
        if (res.status === 404) {
          setEvents([]);
          return;
        }
        if (res.ok) {
          setEvents(body);
          return;
        }
        throw new Error(
          body.reason ? body.reason : 'An error occurred while fetching events',
        );
      })
      .catch((err) => {
        showToast({
          title: 'Error fetching events',
          description: err.message,
          variant: 'error',
        });
      });
  });
  return (<>
    <Show when={events()?.length === 0}>
      <h1 class='text-center font-bold text-muted-foreground'>
        Hmmm, we don't seem to have any events :(
      </h1>
      <p class='text-center text-muted-foreground'>
        Maybe you could post some{' '}
      </p>
    </Show>
    <div class='flex w-4m-6 flex-wrap m-4 space-x-4'>
      <For each={events()}>
        {(event) => (
          <div>
            <div class='flex flex-col' id={event.title}>
              <Card class='flex h-60 min-w-52 flex-col text-wrap justify-center items-center space-y-4 p-3'>
                <Tooltip
                  placement="bottom"
                  openDelay={200}
                  floatingOptions={{
                    offset: 1,
                    flip: true,
                    shift: true,
                  }}
                >
                  <p class='block text-xl border= "white" gap-4 font-bold flex flex-col place-items-center'>{event.title}</p>
                  <Tooltip.Trigger
                    class="my-auto rounded-full bg-corvu-100 p-3 transition-all duration-100 hover:bg-corvu-200 active:translate-y-2"
                  >
                    Event Details
                  </Tooltip.Trigger>
                  <Tooltip.Portal>
                    <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                      <Card class='flex flex-col break-after-page justify-center items-center space-y-4 p-3'><p class="block flex flex-col gap-2 place-items-right">
                        {event.description}</p>
                        <p class="text-lg font-light text-muted-foreground flex flex-col place-items-center">
                          event well happen <time>{moment(event.event_time).fromNow()}</time>
                        </p></Card>
                      <Tooltip.Arrow class="text-corvu-100" />
                    </Tooltip.Content>
                  </Tooltip.Portal>
                </Tooltip>

                <Show
                  when={event.responded_users?.includes(userDetails().user_name)}
                  fallback={
                    <>
                      <Button
                        id={"option1" + String(event.id)}
                        variant='ghost'
                        class='flex-1 gap-2'
                        onClick={() => {
                          event.responded_users?.push(userDetails().user_name);
                          handleEventOption(event.options[0].option_id, event.id);
                        }}

                      >
                        {event.options[0].option_name}
                      </Button>
                      <Button
                        id={"option2" + String(event.id)}
                        variant='ghost'
                        class='flex-1 gap-2'
                        color="red"
                        onClick={() => {
                          event.responded_users?.push(userDetails().user_name);
                          handleEventOption(event.options[1].option_id, event.id);
                        }}
                      >
                        {event.options[1].option_name}
                      </Button>
                    </>
                  }
                >
                  <Tooltip
                    placement="right"
                    openDelay={200}
                    floatingOptions={{
                      offset: 1,
                      flip: true,
                      shift: true,
                    }}
                  >
                    <Tooltip.Trigger
                      class="my-auto rounded-full bg-corvu-100 p-3 transition-all duration-100 hover:bg-corvu-200 active:translate-y-2"
                    >
                      {event.options[0].option_name}
                    </Tooltip.Trigger>
                    <Tooltip.Portal>
                      <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                        <Card class='flex flex-col justify-center items-center space-y-4 p-3'>
                          Choosen By:
                          <Index each={event.responded_users}>
                            {(user, i) => (
                              <Show when={event.choices[i] == event.options[0].option_name}>
                                <p class='block text-sm font-bold hover:underline'>
                                  {event.full_names[i]}
                                </p>
                              </Show>
                            )}
                          </Index>
                        </Card>
                        <Tooltip.Arrow class="text-corvu-100" />
                      </Tooltip.Content>
                    </Tooltip.Portal>
                  </Tooltip>
                  <Tooltip
                    placement="right"
                    openDelay={200}
                    floatingOptions={{
                      offset: 1,
                      flip: true,
                      shift: true,
                    }}
                  >
                    <Tooltip.Trigger
                      class="my-auto rounded-full bg-corvu-100 p-3 transition-all duration-100 hover:bg-corvu-200 active:translate-y-2"
                    >
                      {event.options[1].option_name}
                    </Tooltip.Trigger>
                    <Tooltip.Portal>
                      <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                        <Card class='flex flex-col justify-center items-center space-y-4 p-3'>
                          Choosen By:
                          <Index each={event.responded_users}>
                            {(user, i) => (
                              <Show when={event.choices[i] == event.options[1].option_name}>
                                <p class='block text-sm font-bold hover:underline'>
                                  {event.full_names[i]}
                                </p>
                              </Show>
                            )}
                          </Index>
                        </Card>
                        <Tooltip.Arrow class="text-corvu-100" />
                      </Tooltip.Content>
                    </Tooltip.Portal>
                  </Tooltip>
                </Show>
              </Card>
            </div>
          </div>
        )}
      </For>
    </div>
  </>);
}

export function handleEventOption(option: number, eventID: number) {
  fetchWithAuth(`${config.API_URL}/event_response`, {
    method: 'POST',
    body: JSON.stringify({
      event_id: eventID,
      option_id: option,
    })
  })
    .then(async (res) => {
      if (!res.ok) {
        throw new Error(
          // reason ?? 'An error occurred while responding to request',
        );
      }
      // window.location.reload();
    })
    .catch((err) => {
      showToast({
        title: 'Error responding to request',
        description: err.message,
        variant: 'error',
      });
    });
  var button = document.getElementById("option1" + String(eventID));
  button?.setAttribute('disabled', '');
  var button = document.getElementById("option2" + String(eventID));
  button?.setAttribute('disabled', '');
}
