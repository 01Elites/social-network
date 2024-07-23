import type { Group} from '~/types/group';
import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement } from 'solid-js';


type GroupRequestParams= {
  groupID: string;
}

export default function RequestToJoin(props: { targetGroup: () => Group}):JSXElement{
  function fetchRequest(){
  fetchWithAuth(config.API_URL + '/join_group_req',{
    method:'POST',
    body:JSON.stringify({
        group_id:props.targetGroup().id
    })
  }).then(async (res) => {
    const body = await res.json();
    if (res.status === 200) {
      console.log('RequestMade');
      console.log(body)
      props.targetGroup().request_made = true
      return;
    } else {
      console.log('Error making request');
      return
    }
  })
}
  return ( <Button class="flex grow" variant="default" onclick={fetchRequest}>
    <Follow_Icon /> Request to join
  </Button>)
}