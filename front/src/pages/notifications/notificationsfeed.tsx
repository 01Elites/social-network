import { Index, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Card } from '~/components/ui/card';
import moment from 'moment';
import { FaSolidCheck } from 'solid-icons/fa';
import { IoClose } from 'solid-icons/io';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Show } from 'solid-js';
import { handleInvite } from '../group/request';
import { A } from '@solidjs/router';
import { handleRequest } from '../group/creatorsrequest';
import { handleEventOption } from '../events/eventsfeed';
import Tooltip from '@corvu/tooltip'
import { useContext } from 'solid-js';
import NotificationsContext from '~/contexts/NotificationsContext';
import { RiBusinessCalendarEventLine } from 'solid-icons/ri'
import { Popover, PopoverContent, PopoverTrigger } from "~/components/ui/popover"

export default function NotificationsFeed(): JSXElement {
  // const [test, setnotification] = createSignal<NotificationsHook>();
  const notifications = useContext(NotificationsContext);
  for (let i = 0; i < (notifications?.store.length ?? 0); i++) {
    notifications?.markRead(notifications?.store[i].notification_id);
  }

  return (<>
    <div class="flex-row">
      <div class="flex-row">
        <Index each={notifications?.store}>
          {(notification, i) => (
            <>

              <Show when={notification().type === "FOLLOW_REQUEST"}>
                <div id={notification().metadata.requester.user_name + "follow"}>
                  <Card class="flex flex-col justify-center w-90 p-3 space-y-4 h-fit">
                    <div class="flex items-center space-x-4 text-base">
                      <Avatar class="w-20 h-20 mb-2">
                        <AvatarFallback>
                          <Show
                            when={notification().metadata.requester.avatar}
                            fallback={notification().metadata.requester.first_name.charAt(0).toUpperCase()}
                          >
                            <img
                              alt="avatar"
                              class="object-cover rounded-md"
                              loading="lazy"
                              src={`${config.API_URL}/image/${notification().metadata.requester.avatar}`}
                            />
                          </Show>
                        </AvatarFallback>
                      </Avatar>
                      <div class="flex flex-col items-start justify-center space-y-1">
                        <div><a
                          href={`/profile/${notification().metadata.requester.user_name}`}
                          class="text-base font-bold text-blue-500 hover:underline"
                        >
                          {notification().metadata.requester.first_name} {notification().metadata.requester.last_name}
                        </a></div>
                        <div>requested to follow you</div>
                        <time
                          class="text-xs font-light text-muted-foreground"
                          dateTime={moment(notification().metadata.creation_date).calendar()}
                          title={moment(notification().metadata.creation_date).calendar()}
                        >
                          {moment(notification().metadata.creation_date).fromNow()}
                        </time>
                        <div class="flex gap-2">
                          <Button
                            variant="ghost"
                            class="flex-1 gap-2"
                            onClick={() => {
                              handleFollowRequest("accepted", notification().metadata.requester.user_name);
                              notifications?.markRead(notification().notification_id, true)
                            }}
                          >
                            <FaSolidCheck class="text-green-500 size-4" />
                          </Button>
                          <Button
                            variant="ghost"
                            class="flex-1 gap-2"
                            onClick={() => {
                              handleFollowRequest("rejected", notification().metadata.requester.user_name);
                              notifications?.markRead(notification().notification_id, true)
                            }}
                          >
                            <IoClose class="text-red-500 size-4" />
                          </Button>
                        </div>
                      </div>
                    </div>
                  </Card>
                </div>
                <br></br>

              </Show>
              <Show when={notification().type == "GROUP_INVITATION"}>
                <div id={notification().metadata.id + "invite"}>
                  <Card class="flex flex-col p-3 pl-0 space-y-4 h-fit">
                    <div class="flex items-center space-x-4 text-base">
                      <div></div><Avatar class="w-20 h-20 mb-2">
                        <AvatarFallback>
                          {notification().metadata.title.charAt(0).toUpperCase()}
                        </AvatarFallback>
                      </Avatar>
                      <div class="flex flex-col items-start justify-center space-y-1">
                        <div>{<a
                          href={"/profile/" + notification().metadata.invited_by.user.user_name} class="text-base font-bold text-blue-500 hover:underline">
                          {notification().metadata.invited_by.user.first_name}  {notification().metadata.invited_by.user.last_name}</a>}
                          &nbspinvited you to join:  {<A
                            href={"/group/" + notification().metadata.id} class="text-base font-bold hover:underline">
                            {notification().metadata.title}</A>}</div>
                        <time
                          class='text-xs font-light text-muted-foreground'
                          dateTime={moment(notification().metadata.invited_by.creation_date).calendar()}
                          title={moment(notification().metadata.invited_by.creation_date).calendar()}
                        >
                          {moment(notification().metadata.invited_by.creation_date).fromNow()}</time>
                        <div class='flex gap-2'>
                          <Button
                            variant='ghost'
                            class='flex-1 gap-2'
                            onClick={() => {
                              {
                                handleInvite("accepted", notification().metadata.id, notification().metadata.invited_by.user.user_name);
                                notifications?.markRead(notification().notification_id, true)
                              }
                            }}
                          >
                            <FaSolidCheck
                              class='size-4'
                              color='green'
                            />
                          </Button>
                          <Button
                            variant='ghost'
                            class='flex-1 gap-2'
                            color="red"
                            onClick={() => {
                              {
                                handleInvite("rejected", notification().metadata.id, notification().metadata.invited_by.user.user_name);
                                notifications?.markRead(notification().notification_id, true)
                              }
                            }}
                          >
                            <IoClose class='size-4' color='red' />
                          </Button>
                        </div>
                      </div>
                    </div>
                  </Card>
                </div>
                <br></br>
              </Show>


              <Show when={notification().type == "REQUEST_TO_JOIN_GROUP"}>
                <div id={notification().metadata.id + notification().metadata.requester.user.user_name}>
                  <Card class="flex flex-col justify-center w-90 p-3  pl-0 space-y-4 h-fit">
                    <div class="flex items-center space-x-4 text-base">
                      <div></div><Avatar class="w-20 h-20 mb-2">
                        <AvatarFallback>
                          <Show when={notification().metadata.requester.user.avatar} fallback={
                            notification().metadata.requester.user.first_name.charAt(0).toUpperCase()
                          }><img
                              alt='avatar'
                              class='size-full rounded-md rounded-b-none object-cover'
                              loading='lazy'
                              src={`${config.API_URL}/image/${notification().metadata.requester.user.avatar}`}
                            /></Show></AvatarFallback>
                      </Avatar>
                      <div class="flex flex-col items-start justify-center space-y-1">
                        <div><a
                          href={"/profile/" + notification().metadata.requester.user.user_name} class='text-base font-bold text-blue-500 hover:underline'>
                          {notification().metadata.requester.user.first_name}  {notification().metadata.requester.user.last_name} </a>
                          Requested to join: {<A
                            href={"/group/" + notification().metadata.id} class="text-base font-bold hover:underline">
                            {notification().metadata.title}</A>}</div>
                        <time
                          class='text-xs font-light text-muted-foreground'
                          dateTime={moment(notification().metadata.requester.creation_date).calendar()}
                          title={moment(notification().metadata.requester.creation_date).calendar()}
                        >
                          {moment(notification().metadata.requester.creation_date).fromNow()}</time>
                        <div class='flex gap-2'>

                          <Button
                            variant='ghost'
                            class='flex-1 gap-2'
                            onClick={() => {
                              {
                                handleRequest("accepted", notification().metadata.id, notification().metadata.requester.user.user_name);
                                notifications?.markRead(notification().notification_id, true)
                              }
                            }}
                          >
                            <FaSolidCheck
                              class='size-4'
                              color='green'
                            />
                          </Button>
                          <Button
                            variant='ghost'
                            class='flex-1 gap-2'
                            color="red"
                            onClick={() => {
                              {
                                handleRequest("rejected", notification().metadata.id, notification().metadata.requester.user.user_name);
                                notifications?.markRead(notification().notification_id, true)
                              }
                            }}
                          >
                            <IoClose class='size-4' color='red' />
                          </Button>
                        </div>
                      </div>
                    </div>
                  </Card>

                </div>
                <br></br>

              </Show>

              <Show when={notification().type == "EVENT"}>
                <div id={notification().metadata.event.id + notification().metadata.event.title}>
                  <Card class="flex flex-col p-3 h-fit">
                    <div class="flex items-center space-x-4 text-base">
                      <div><RiBusinessCalendarEventLine class="w-20 h-14" />
                        <Popover>
                          <PopoverTrigger as={Button<"button">} variant={"ghost"} class="flex items-center p-4">
                            Details
                          </PopoverTrigger>
                          <PopoverContent>
                              <p class="z-100 max-w-15 block flex flex-wrap gap-2 place-items-center break-all"> {notification().metadata.event.description}
                              </p>
                              {moment(notification().metadata.event.event_time).calendar()}
                          </PopoverContent>
                        </Popover>
                      </div>
                      <div class="flex-col justify-center space-y-1">
                        {notification().metadata.event.title}<br></br>group: <A href={"/group/" + notification().metadata.group.id}
                          class="text-base font-bold hover:underline">
                          {notification().metadata.group.title}
                        </A>
                        {/* <Tooltip
                      placement="left"
                      openDelay={200}
                      // strategy='absolute'
                      floatingOptions={{
                        offset: 1,
                        flip: true,
                        shift: false,
                      }}
                    >
                      <Tooltip.Trigger>
                        Event Details
                      </Tooltip.Trigger>
                      <Tooltip.Portal>
                        <Tooltip.Content class="rounded-lg px-3 py-2 font-medium">
                          <Card class='z-100 flex flex-col justify-center items-center space-y-4 p-3'>
                            <p class="z-100 block flex flex-col gap-2 place-items-right"> {notification().metadata.event.description}
                            </p>
                        {moment(notification().metadata.event.event_time).calendar()}
                          </Card>
                        </Tooltip.Content>
                      </Tooltip.Portal>
                    </Tooltip> */}
                        <div class="flex flex-wrap gap-6">
                          <Button
                            id={"option1" + String(notification().metadata.event.id)}
                            variant='ghost'
                            class="p-2"
                            onClick={() => {
                              let elem = document.getElementById(notification().metadata.event.id + notification().metadata.event.title); elem?.remove();
                              ; notifications?.markRead(notification().notification_id, true);
                              handleEventOption(notification().metadata.event.options[0].option_id, notification().metadata.event);
                            }}

                          >
                            {notification().metadata.event.options[0].option_name}
                          </Button>
                          <Button
                            id={"option2" + String(notification().metadata.event.id)}
                            variant='ghost'
                            class="p-0"
                            color="red"
                            onClick={() => {
                              let elem = document.getElementById(notification().metadata.event.id); elem?.remove();
                              ; notifications?.markRead(notification().notification_id, true);
                              handleEventOption(notification().metadata.event.options[1].option_id, notification().metadata.event);
                            }}
                          >
                            {notification().metadata.event.options[1].option_name}
                          </Button>
                        </div>
                      </div>
                    </div>
                  </Card>
                </div>
                <br></br>
              </Show>
            </>)}
        </Index>
      </div>
    </div>
  </>)
}



function handleFollowRequest(response: string, follower: string) {
  fetchWithAuth(`${config.API_URL}/follow_response`, {
    method: 'POST',
    body: JSON.stringify({
      follower: follower,
      status: response,
    })
  })
    .then(async (res) => {
      if (!res.ok) {
        throw new Error(
          // reason ?? 'An error occurred while responding to request',
        );
      }
    })
    .catch((err) => {
      console.log('Error responding to request');
    });
  const elem = document.getElementById(follower + "follow");
  elem?.remove();
}
