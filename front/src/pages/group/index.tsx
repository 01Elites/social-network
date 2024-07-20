import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { useParams } from '@solidjs/router';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import type {Group} from '~/types/group';
import Layout from '~/Layout';
import GroupDetails from './details';
import GroupPostFeed from './groupfeed';

  export type GroupParams = {
    id: string;
  };

export default function Group(): JSXElement {

  const [targetGroup, setTargetGroup] = createSignal<Group | undefined>();
  
  const params: GroupParams = useParams();
  const groupID = params.id;
  
  createEffect(() => {
  fetchWithAuth(config.API_URL + '/group/' + groupID).then(async (res) => {
    const body = await res.json();
    if (res.status === 404) {
      console.log('User not found');
      return;
    }
    if (res.ok) {
      setTargetGroup(body);
      return;
    }
  })
})

  return (<><Layout>
    <div class='grid grid-cols-1 md:grid-cols-6 m-4 '> {/* Main grid */}
      <div class='col-span-2'>
          <GroupDetails targetGroup={() => targetGroup() as Group} />
      </div>
      <Show when={targetGroup()?.ismember}>
      <div class='col-span-4 overflow-y-auto'>
        <GroupPostFeed groupID={groupID as string} />
      </div>
        </Show>
    </div> {/* Main grid */}
  </Layout></>)
}