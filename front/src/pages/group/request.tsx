import type { Group} from '~/types/group';
import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement } from 'solid-js';
import { Show } from 'solid-js';


type GroupRequestParams= {
  groupID: string;
}

export default function RequestToJoin(props: { targetGroup: () => Group}):JSXElement{
  function fetchRequest(data: string){
  fetchWithAuth(config.API_URL + data,{
    method:'POST',
    body:JSON.stringify({
        group_id:props.targetGroup().id
    })
  }).then(async (res) => {
    if (res.status === 200) {
      if (data === '/join_group_req'){
        props.targetGroup().request_made = true
        console.log('RequestMade');
      } else {
        props.targetGroup().request_made = false
      }
      return;
    } else {
      console.log('Error making request');
      return
    }
  })
}
  return (<><Show when={!props.targetGroup().ismember}
  fallback={<Button class="flex grow" variant="default" onclick={[fetchRequest,"/exit_group"]} name="request">
  <Follow_Icon /> exit group
</Button>}> 
<Show when={!props.targetGroup().request_made}
fallback={<Button class="flex grow" variant="default" onclick={[fetchRequest,"/cancel_join_req"]} name="request">
<Follow_Icon /> Cancel Request
</Button>
}>
<Button class="flex grow" variant="default" onclick={[fetchRequest,"/join_group_req"]} name="request">
    <Follow_Icon /> Request to join
  </Button>
</Show>
</Show>
 <div class='flex flex-row w-full justify-between gap-2 m-4'>
   <div class='flex gap-2'> {/* Follow button */}
   </div> {/* Follow button */}
 </div>
     </>)
}