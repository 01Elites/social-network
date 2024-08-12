import type { Group} from '~/types/group';
import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement, Show } from 'solid-js';
import { createSignal } from 'solid-js'
import moment from 'moment';
import { IoClose } from 'solid-icons/io';
import { FaSolidCheck } from 'solid-icons/fa';
import { A } from '@solidjs/router';
import { Card } from '~/components/ui/card';

export default function RequestToJoin(props: { targetGroup: () => Group}):JSXElement{
  var [buttonData, setButtonData] = createSignal(["", ""]);
  console.log(props.targetGroup())
  if (props.targetGroup().iscreator){
    return null
  }else if (props.targetGroup().ismember){
    setButtonData(["Exit", "/exit_group"])
  } else if (props.targetGroup().request_made){
  setButtonData(["Requested", "/cancel_join_req"])
} else {
    setButtonData(["Join", "/join_group_req"])
  }
  function sendRequestApi(){
    console.log(buttonData())
    if (buttonData()[1] === ""){
      return
    }
  fetchWithAuth(config.API_URL + buttonData()[1],{
    method:'POST',
    body:JSON.stringify({
        group_id:props.targetGroup().id
    })
  }).then(async (res) => {
    if (res.status === 200) {
      if (buttonData()[1] === "/join_group_req"){
        setButtonData(["Requested", "/cancel_join_req"])
        props.targetGroup().request_made = true
        console.log('RequestMade');
      } else if (buttonData()[1] === "/cancel_join_req"){
        setButtonData(["Join", "/join_group_req"])      
        props.targetGroup().request_made = false
        console.log('RequestCancelled');
      } else if (buttonData()[1] === "/exit_group"){
        setButtonData(["Join", "/join_group_req"])
        props.targetGroup().ismember = false
        props.targetGroup().request_made = false
        console.log('GroupExited');
      }
      return;
    } else {
      console.log(res.body);
      console.log('Error making request');
      return
    }
  })
}
  return (<>
    <Show when={props.targetGroup().invited_by.user.first_name !== ""}>
      <div id={props.targetGroup().invited_by.user.user_name}>
      <Card class='flex flex-col items-center space-y-4 p-3'>
      <p class="flex-col justify-center items-center">
            {<A
        href={"/profile/" + props.targetGroup().invited_by.user.user_name} class='flex flex-col justify-center items-center'>
    {props.targetGroup().invited_by.user.first_name}  {props.targetGroup().invited_by.user.last_name}</A> }
    invited you to join this group<br></br>
    <time
    class='text-xs font-light text-muted-foreground'
    dateTime={moment(props.targetGroup().invited_by.creation_date).calendar()}
    title={moment(props.targetGroup().invited_by.creation_date).calendar()}
  >
    {moment(props.targetGroup().invited_by.creation_date).fromNow()}</time>
    </p>
    <div class='flex flex-row gap-2'>
      <Button
      variant='ghost'
      class='flex-1 gap-2'
      onClick={() => {handleInvite("accepted", props.targetGroup().id, props.targetGroup().invited_by.user.user_name);}}
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
      onClick={() => {handleInvite("rejected", props.targetGroup().id, props.targetGroup().invited_by.user.user_name)}}
    >
    <IoClose class='size-4' color='red'/>
    </Button>
    </div>
  </Card></div>      
</Show>
<Show when={buttonData()[0] == "Exit" || props.targetGroup().invited_by.user.first_name == ""}>
<Button class="flex grow" variant="default" onClick={sendRequestApi}>
<Show when={buttonData()[0] == "Join"}>
  <Follow_Icon />
</Show>
{buttonData()[0]}
</Button>
</Show>
     </>)
}

export function handleInvite(response: string, groupID: number, invitee: string) {
  fetchWithAuth(`${config.API_URL}/invitation_response`, {
    method: 'PATCH',
    body: JSON.stringify({
      group_id: Number(groupID),
      response: response,
  })})
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
window.location.reload();
}