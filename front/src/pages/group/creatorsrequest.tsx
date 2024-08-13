import { Index } from 'solid-js';
import { Button } from "~/components/ui/button";
import { IoClose } from 'solid-icons/io'
import { FaSolidCheck } from 'solid-icons/fa'
import { A } from "@solidjs/router";
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import moment from 'moment';;
import { JSXElement } from 'solid-js';
import { showToast } from '~/components/ui/toast';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { requester } from "~/pages/group/groupfeed";
import { Card } from '~/components/ui/card';

interface GroupRequestParams {
  requesters: requester[] | undefined;
  groupID: string;
}

export function GroupRequests(params: GroupRequestParams): JSXElement {
  return (
    <Index each={params.requesters}>
      {(requester, i) => <> <div class='flex w-full space-x-1 '>
        <Card id={params.groupID+requester().user.user_name} class='flex w-44 flex-col justify-center items-center space-y-4 p-3 justfi'>
          <p class="flex flex-col gap-2 place-items-center"><Avatar>
            <AvatarImage src={requester().user.avatar} />
            <AvatarFallback>
              {requester().user.first_name.charAt(0).toUpperCase()}
            </AvatarFallback>
          </Avatar><A
            href={"/profile/" + requester().user.user_name} class='block text-sm font-bold hover:underline'>
              {requester().user.first_name}  {requester().user.last_name}</A>
            <time
              class='text-xs font-light text-muted-foreground'
              dateTime={moment(requester().creation_date).calendar()}
              title={moment(requester().creation_date).calendar()}
            >
              {moment(requester().creation_date).fromNow()}</time></p>
              <div class='flex flex-row gap-2'>

              <Button
                variant='ghost'
                class='flex-1 gap-2'
                onClick={() => { handleRequest("accepted", params.groupID, requester().user.user_name); params.requesters?.splice(i, i + 1) }}
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
            onClick={() => { handleRequest("rejected", params.groupID, requester().user.user_name); params.requesters?.splice(i, i + 1) }}
          >
            <IoClose class='size-4' color='red' />
          </Button>
          </div>
        </Card></div></>}
    </Index>
  )
}

export function handleRequest(response: string, groupID: string, requester: string) {
  fetchWithAuth(`${config.API_URL}/join_group_res`, {
    method: 'PATCH',
    body: JSON.stringify({
      requester: requester,
      group_id: Number(groupID),
      response: response,
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
      showToast({
        title: 'Error responding to request',
        description: err.message,
        variant: 'error',
      });
    });
  const elem = document.getElementById(groupID + requester);
  elem?.remove();
  
}