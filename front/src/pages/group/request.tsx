import type { Group} from '~/types/group';
import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { createEffect, JSXElement } from 'solid-js';
import { createSignal } from 'solid-js';


type GroupRequestParams= {
  groupID: string;
}

export default function RequestToJoin(props: { targetGroup: () => Group}):JSXElement{
  var [buttonData, setButtonData] = createSignal(["", ""]);
  if (props.targetGroup().ismember){
    setButtonData(["Exit Group", "/exit_group"])
  } else if (props.targetGroup().request_made){
  setButtonData(["Cancel Request", "/cancel_join_req"])
} else {
    setButtonData(["Request to Join", "/join_group_req"])
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
        setButtonData(["Cancel Request", "/cancel_join_req"])
        props.targetGroup().request_made = true
        console.log('RequestMade');
      } else if (buttonData()[1] === "/cancel_join_req"){
        setButtonData(["Request to Join", "/join_group_req"])      
        props.targetGroup().request_made = false
        console.log('RequestCancelled');
      } else if (buttonData()[1] === "/exit_group"){
        setButtonData(["Request to Join", "/join_group_req"])
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
<Button class="flex grow" variant="default" onClick={sendRequestApi}>
<Follow_Icon />{buttonData()[0]}
</Button>
     </>)
}