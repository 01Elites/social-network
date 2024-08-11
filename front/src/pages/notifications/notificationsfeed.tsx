import { For, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Card } from '~/components/ui/card';
import moment from 'moment';
import { FaSolidCheck } from 'solid-icons/fa';
import { IoClose } from 'solid-icons/io';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { Show } from 'solid-js';
import FollowRequest from '../profile/followRequest';
import { Notifications } from '~/types/notifications';
import { createSignal } from 'solid-js'
import { EventsFeed } from '../events/eventsfeed';
import { handleInvite } from '../group/request';
import { A } from '@solidjs/router';
import { handleRequest } from '../group/creatorsrequest';
import { handleEventOption } from '../events/eventsfeed';
import Tooltip from '@corvu/tooltip'
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { useContext } from 'solid-js';
import { UserDetailsHook } from '~/hooks/userDetails';

export default function NotificationsFeed(): JSXElement {
  const [notifications, setnotification] = createSignal<Notifications[]>([]);
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
// setnotification([{
//   event: "NOTIFICATION",
//   id: 2,
//   payload: {
//       type: "FOLLOW_REQUEST",
//       message: "You have a new follow request",
//       notification_id: 666,
//       read: false,
//       metadata: {
//           requester: {
//               first_name: "John",
//               last_name: "Doe",
//               username: "johndoe",
//               avatar:""
//           },
//        creation_date:"2025-07-07 00:00:00 +0000 UTC"
//   }
// }
// },
// {
//   event: "NOTIFICATION",
//   id: 1,
//   payload: {
//       type: "GROUP_INVITATION",
//       message: "You have a new group invitation",
//       notification_id: 777,
//       read: false,
//       metadata: {
//               id: "1",
//               title: "Group 1",
//              invited_by:{
//               user:{ first_name:"test",
//               last_name:"next",
//               user_name:"user",
//               avatar:"",},
//              created_by: "2025-07-07 00:00:00 +0000 UTC"
//           }
//       }
// }
// },
// {
//   event: "NOTIFICATION",
//   id: 3,
//   payload: {
//       type: "REQUEST_TO_JOIN_GROUP",
//       message: "You have a new request to join the group",
//       notification_id: 888,
//       read: false,
//       metadata: {
//           group: {
//               id: "1",
//               title: "Group 1"
//           },
//           requester: {
//             user:{
//              first_name: "John",
//              last_name: "Doe",
//              username: "johndoe",
//              avatar:"",
//               },
//             creation_date:"2024-12-31 00:00:00 +0000 UTC"
//          }
//       }
//   }
// },
// {
//   event: "NOTIFICATION",
//   id: 4,
//   payload: {
//       type: "EVENT",
//       message: "You have a new event in the group",
//       notification_id: 222,
//       read: false,
//       metadata: {
//           group: {
//               id: "1",
//               title: "Group 1",
//           },
//           event: {
//               id: "1",
//               title: "Event 1",
//               description: "description",
//               event_time:"2025-07-07 00:00:00 +0000 UTC",
//               options: [{
//                        id: 0,
//                        name:"option1",
//           }, {
//             id: 1,
//             name:"option2",
//           }]
//       }
//   }
// }
// }
// ])
  return (<>
    <div class="flex-row">
      <div class="flex-row">
        <For each={notifications()}>
          {(notification) =>(
            <>
            <Show when={notification.payload.type == "FOLLOW_REQUEST"}>
        <div id={notification.payload.metadata.requester.user_name}>
        <Card class='flex h-80 w-80 flex-col justify-center items-center space-y-4 p-3 justfi'>
            <a
              href={`/profile/${notification.payload.metadata.requester.user_name}`}
              class='flex flex-col items-center text-base font-bold hover:underline text-blue-500'
            >
              <Avatar class='w-[5rem] h-[5rem] mb-2'>
                <AvatarFallback>
                  <Show when={notification.payload.metadata.requester.avatar} fallback={
                    notification.payload.metadata.requester.first_name.charAt(0).toUpperCase()
                  }><img
                      alt='avatar'
                      class='size-full rounded-md rounded-b-none object-cover'
                      loading='lazy'
                      src={`${config.API_URL}/image/${notification.payload.metadata.requester.avatar}`}
                    /></Show></AvatarFallback>
              </Avatar><br></br>
              <div class='flex flex-wrap items-center justify-center space-x-1'>
                <div>{notification.payload.metadata.requester.first_name}</div>
                <div>{notification.payload.metadata.requester.last_name}</div>
              </div>
            </a>
            requsted to follow you
            <time
              class='text-xs font-light text-muted-foreground'
              dateTime={moment(notification.payload.metadata.creation_date).calendar()}
              title={moment(notification.payload.metadata.creation_date).calendar()}
            >
              {moment(notification.payload.metadata.creation_date).fromNow()}</time>
            <div class='flex flex-row gap-2'>
              <Button
                variant='ghost'
                class='flex-1 gap-2'
                onClick={() => { handleFollowRequest("accepted", notification.payload.metadata.requester.user_name); }}
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
                onClick={() => { handleFollowRequest("rejected", notification.payload.metadata.requester.user_name) }}
              >
                <IoClose class='size-4' color='red' />
              </Button>
            </div>
          </Card>
        </div>
      <br></br>
      </Show>
      <Show when={notification.payload.type == "GROUP_INVITATION"}>
      <div id={notification.payload.metadata.invited_by.user.user_name}>
        <Card class='flex h-80 w-80 flex-col justify-center items-center space-y-4 p-3 justfi'>
          <Avatar class='w-[5rem] h-[5rem] mb-2'>
            <AvatarFallback>
              {notification.payload.metadata.title.charAt(0).toUpperCase()}
            </AvatarFallback>
          </Avatar>
          <p class="flex-col justify-center items-center">
            {<A
              href={"/profile/" + notification.payload.metadata.invited_by.user.user_name} class='flex flex-col justify-center items-center'>
              {notification.payload.metadata.invited_by.user.first_name}  {notification.payload.metadata.invited_by.user.last_name}</A>}
            invited you to join {<A
              href={"/group/" + notification.payload.metadata.id} class='flex flex-col justify-center items-center'>
              {notification.payload.metadata.title}<br></br></A>}
            <time
              class='text-xs font-light text-muted-foreground'
              dateTime={moment(notification.payload.metadata.invited_by.creation_date).calendar()}
              title={moment(notification.payload.metadata.invited_by.creation_date).calendar()}
            >
              {moment(notification.payload.metadata.invited_by.creation_date).fromNow()}</time>
          </p>
          <div class='flex flex-row gap-2'>
            <Button
              variant='ghost'
              class='flex-1 gap-2'
              onClick={() => { handleInvite("accepted", notification.payload.metadata.group.id, notification.payload.metadata.invited_by.user.user_name); }}
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
              onClick={() => { handleInvite("rejected", notification.payload.metadata.group.id, notification.payload.metadata.invited_by.user.user_name) }}
            >
              <IoClose class='size-4' color='red' />
            </Button>
          </div>
        </Card></div>
      <br></br>
      </Show>
      <Show when={notification.payload.type == "REQUEST_TO_JOIN_GROUP"}>
      <div id={notification.payload.metadata.requester.user.user_name} class="flex items-center">
        <Card id={notification.payload.metadata.requester.user.user_name} class='flex h-80 w-80 flex-col justify-center items-center space-y-4 p-3 justfi'>
          <p class="flex flex-col gap-2 place-items-center">
            <Avatar class='w-[5rem] h-[5rem] mb-2'>
              <AvatarFallback>
                <Show when={notification.payload.metadata.requester.user.avatar} fallback={
                  notification.payload.metadata.requester.user.first_name.charAt(0).toUpperCase()
                }><img
                    alt='avatar'
                    class='size-full rounded-md rounded-b-none object-cover'
                    loading='lazy'
                    src={`${config.API_URL}/image/${notification.payload.metadata.requester.user.avatar}`}
                  /></Show></AvatarFallback>
            </Avatar><A
              href={"/profile/" + notification.payload.metadata.requester.user.user_name} class='block text-sm font-bold hover:underline'>
              {notification.payload.metadata.requester.user.first_name}  {notification.payload.metadata.requester.user.last_name} </A>
            Requested to join {notification.payload.metadata.group.title}<br></br>
            <time
              class='text-xs font-light text-muted-foreground'
              dateTime={moment(notification.payload.metadata.requester.creation_date).calendar()}
              title={moment(notification.payload.metadata.requester.creation_date).calendar()}
            >
              {moment(notification.payload.metadata.requester.creation_date).fromNow()}</time></p>
          <div class='flex flex-row gap-2'>

            <Button
              variant='ghost'
              class='flex-1 gap-2'
              onClick={() => { handleRequest("accepted", notification.payload.metadata.requester.id, notification.payload.metadata.requester.user.user_name) }}
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
              onClick={() => { handleRequest("rejected", notification.payload.metadata.requester.id, notification.payload.metadata.requester.user.user_name) }}
            >
              <IoClose class='size-4' color='red' />
            </Button>
          </div>
        </Card>

      </div>
      <br></br>
      </Show>
      <Show when={notification.payload.type == "EVENT"}>
        <div id={notification.payload.metadata.group.title}>
          <Card class='flex h-80 w-80 flex-col justify-center items-center space-y-4 p-3 justfi'>
            <Tooltip
              placement="bottom"
              openDelay={200}
              floatingOptions={{
                offset: 1,
                flip: true,
                shift: true,
              }}
            >
              <Avatar class='w-[5rem] h-[5rem] mb-2'>
                <AvatarFallback>
                  {notification.payload.metadata.group.title.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <p class='block text-xl border= "white" gap-4 font-bold flex flex-col place-items-center'>{notification.payload.metadata.group.title}</p>
              <p class='block text-xl border= "white" gap-4 font-bold flex flex-col place-items-center'>{notification.payload.metadata.event.title}</p>
              <Tooltip.Trigger
                class="my-auto rounded-full bg-corvu-100 p-3 transition-all duration-100 hover:bg-corvu-200 active:translate-y-2"
              >
                Event Description
              </Tooltip.Trigger>
              <p class="text-lg font-light text-muted-foreground flex place-items-center">
                <Show when={moment(notification.payload.metadata.event.event_time).isAfter(moment())} fallback={<p>Event Done</p>}>
                  event&nbsp<time>{moment(notification.payload.metadata.event.event_time).fromNow()}</time></Show>
              </p>
              <Tooltip.Portal>
                <Tooltip.Content class="rounded-lg bg-corvu-100 px-3 py-2 font-medium corvu-open:animate-in corvu-open:fade-in-50 corvu-open:slide-in-from-bottom-1 corvu-closed:animate-out corvu-closed:fade-out-50 corvu-closed:slide-out-to-bottom-1">
                  <Card class='flex flex-col break-after-page justify-center items-center space-y-4 p-3'><p class="block flex flex-col gap-2 place-items-right">
                    {notification.payload.metadata.event.description}</p>
                  </Card>
                  <Tooltip.Arrow class="text-corvu-100" />
                </Tooltip.Content>
              </Tooltip.Portal>
            </Tooltip>
              <Button
                id={"option1" + String(notification.payload.metadata.event.id)}
                variant='ghost'
                class='flex-1 gap-2'
                onClick={() => {
                  handleEventOption(notification.payload.metadata.event.options[0].option_id, notification.payload.metadata.event);
                }}

              >
                {notification.payload.metadata.event.options[0].option_name}
              </Button>
              <Button
                id={"option2" + String(notification.payload.metadata.event.id)}
                variant='ghost'
                class='flex-1 gap-2'
                color="red"
                onClick={() => {
                  handleEventOption(notification.payload.metadata.event.options[1].option_id, notification.payload.metadata.event);
                }}
              >
                {notification.payload.metadata.event.options[1].option_name}
              </Button>
          </Card>
        </div>
        </Show>
      </>)}
      </For>
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
  const elem = document.getElementById(follower);
  elem?.remove();
}
