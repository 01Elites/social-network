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
import { UserDetailsHook } from '~/hooks/userDetails';
import { showToast } from '../../components/ui/toast';
import { GroupEvent } from '~/types/group/index';
import moment from 'moment';
import { Card } from '~/components/ui/card';
import { Button } from '~/components/ui/button';
import Tooltip from '@corvu/tooltip'
import { RiBusinessCalendarEventLine } from 'solid-icons/ri'
import NotificationsContext from '~/contexts/NotificationsContext';

type eventProps = {
  events: GroupEvent[] | undefined
}

export function EventsFeed(props: eventProps): JSXElement{
  const [notificationId, setNotificationId] = createSignal<string>('');
  const notifications = useContext(NotificationsContext);
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
 function handleEventOption(option: number, event: GroupEvent) {
    console.log(event);
  let option1Count = 0;
  let option2Count = 0;
  if (!event.options[0].usernames) {
  option1Count = 0;
  } else {
  option1Count = event.options[0].usernames.length;
  }
  if (!event.options[1].usernames) {
  option2Count = 0;
  } else {
  option2Count = event.options[1].usernames.length;
  }
  fetchWithAuth(`${config.API_URL}/event_response`, {
  method: 'POST',
  body: JSON.stringify({
    event_id: event.id,
    option_id: option,
  })
  })
  .then(async (res) => {
    if (res.ok) {
      if (event.options[0].option_id == option) {
        option1Count++;
      } else {
        option2Count++;
      }
      var button1 = document.getElementById("option1" + String(event.id));
  button1?.setAttribute('disabled', '');
  button1 ? button1.innerHTML = `${event.options[0].option_name} (${option1Count})` : null;
  
  var button2 = document.getElementById("option2" + String(event.id));
  button2?.setAttribute('disabled', '');
  button2 ? button2.innerHTML = `${event.options[1].option_name} (${option2Count})` : null;
  let data = await res.json();
  console.log(data)
  setNotificationId(data);
  }
  })
  .catch((err) => {
    showToast({
      title: 'Error responding to event',
      description: err.message,
      variant: 'error', 
    });
  });
  return 
  }
  return (<>
  <Show when={props.events?.length === 0}>
  <h1 class='text-center font-bold text-muted-foreground'>
    Hmmm, we don't seem to have any events :(
  </h1>
  <p class='text-center text-muted-foreground'>
    Maybe you could post some{' '}
  </p>
</Show>
<div class='flex flex-wrap -ml-3'>
<For each={props.events}>
  {(event) => (
    <div>
      <div class='flex flex-col mr-3 mb-4' id={event.title}>
      <Card class="flex flex-col p-3 w-96 h-36">
      <Tooltip
          placement="bottom"
          openDelay={200}
          floatingOptions={{
            offset: 1,
            flip: true,
            shift: true,
          }}
        ><div class="grid grid-cols-10 mb-3">
          <RiBusinessCalendarEventLine class="w-10 h-6 mt-1" />
          <p class='block text-xl border= "white" gap-1 font-bold col-span-7 ml-1'>
          {event.title}</p> 
          <Tooltip.Trigger
            class="my-auto rounded-full bg-corvu-100 transition-all duration-100 hover:bg-corvu-200 active:translate-y-2"
          >
                  <div class="mr-2">Details</div> 
          </Tooltip.Trigger>
         </div>
          <p class="text-lg font-light text-muted-foreground ml-10">
                  <Show when={moment(event.event_time).isAfter(moment())} fallback={<p>Event Done</p>}>
                  event&nbsp<time>{moment(event.event_time).fromNow()}</time></Show>
                </p>
          <Tooltip.Portal>
            <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
              <Card class='flex flex-col break-after-page justify-center items-center space-y-4 p-3'><p class="block flex flex-col gap-2 place-items-right">
                {event.description}</p>
                <p>{moment(event.event_time).calendar()}</p>
                <p>created by {event.creator.first_name} {event.creator.last_name}
                </p></Card>
              <Tooltip.Arrow class="text-corvu-100" />
            </Tooltip.Content>
          </Tooltip.Portal>
        </Tooltip>

        <Show
          when={(event.options[0].usernames && event.options[1].usernames) ||
            (!event.options[1].usernames && event.options[0].usernames?.includes(userDetails().user_name))||
             (!event.options[0].usernames && event.options[1].usernames?.includes(userDetails().user_name))}
          fallback={
            <>
            <div class="flex flex-row ml-6">
              <Button
                id={"option1" + String(event.id)}
                variant='ghost'
                class='flex-col w-22'
                onClick={() => {
                  handleEventOption(event.options[0].option_id, event);
                  setTimeout(() => {
                  notifications?.markRead(notificationId(), true)
                  }, 1000)
                }}
                
              >
                {event.options[0].option_name}
              </Button>
              <Button
                id={"option2" + String(event.id)}
                variant='ghost'
                class='flex-col w-22'
                color="red"
                onClick={() => {
                  (handleEventOption(event.options[1].option_id, event))
                }}
              >
                {event.options[1].option_name}
              </Button>
              </div>
            </>
          }
        >
          <div class="flex flex-row ml-7">
          <Tooltip
            placement="bottom"
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

              {event.options[0].option_name} (<Show when={event.options[0].usernames == undefined}>0</Show>{event.options[0].usernames?.length})
            </Tooltip.Trigger>
            <Tooltip.Portal>
              <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                <Card class='flex flex-col justify-center items-center space-y-4 p-3'>
                Choosen By:
                  <For each={event.options[0].fullnames}>
                    {(user) => (
                        <p class='block text-sm font-bold hover:underline'>
                          {user}
                        </p>
                    )}
                  </For>
                </Card>
                <Tooltip.Arrow class="text-corvu-100" />
              </Tooltip.Content>
            </Tooltip.Portal>
          </Tooltip>
          <Tooltip
            placement="bottom"
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
              {event.options[1].option_name} (<Show when={event.options[1].usernames == undefined}>0</Show>{event.options[1].usernames?.length})
            </Tooltip.Trigger>
            <Tooltip.Portal>
              <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                <Card class='flex flex-col justify-center items-center space-y-4 p-3'>
                  Choosen By:
                  <For each={event.options[1].fullnames}>
                    {(user) => (
                        <p class='block text-sm font-bold hover:underline'>
                          {user}
                        </p>
                    )}
                  </For>
                </Card>
                <Tooltip.Arrow class="text-corvu-100" />
              </Tooltip.Content>
            </Tooltip.Portal>
          </Tooltip>                                               
          </div>                                                                            
        </Show>
      </Card>
      </div>
    </div>
  )}
</For>
</div>
</>);
}
