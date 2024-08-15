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

type eventProps = {
  events: GroupEvent[] | undefined
}

export function EventsFeed(props: eventProps): JSXElement{
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  return (<>
  <Show when={props.events?.length === 0}>
  <h1 class='text-center font-bold text-muted-foreground'>
    Hmmm, we don't seem to have any props.events :(
  </h1>
  <p class='text-center text-muted-foreground'>
    Maybe you could post some{' '}
  </p>
</Show>
<div class='flex flex flex-wrap m-4'>
<For each={props.events}>
  {(event) => (
    <div>
      <div class='flex flex-col' id={event.title}>
      <Card class='flex h-80 w-60 flex-col text-wrap items-center space-y-5 p-3'>
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
          <p class="text-lg font-light text-muted-foreground flex place-items-center">
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
              <Button
                id={"option1" + String(event.id)}
                variant='ghost'
                class='flex-1 gap-2'
                onClick={() => {
                  handleEventOption(event.options[0].option_id, event);
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
                  handleEventOption(event.options[1].option_id, event);
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
        </Show>
      </Card>
      </div>
    </div>
  )}
</For>
</div>
</>);
}

export function handleEventOption(option: number, event: GroupEvent) {
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
}
