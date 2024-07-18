import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { useParams } from '@solidjs/router';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import type {Group} from '~/types/group';
import Layout from '~/Layout';
import type {Post} from '~/types/Post';
import GroupDetails from './details';
import GroupFeed from './groupfeed';

  type GroupParams = {
    id: string;
  };

export default function Group(): JSXElement {

  const [targetGroup, setTargetGroup] = createSignal<Group | undefined>();
  
  const [groupPosts, setGroupPosts] = createSignal<Post[]>();
  const params: GroupParams = useParams();
  
  createEffect(() => {
  fetchWithAuth(config.API_URL + '/group/' + params.id).then(async (res) => {
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
createEffect(() => {
  fetchWithAuth(config.API_URL + '/group/' + params.id + '/posts').then(async (res) => {
    const body = await res.json();
    if (res.status === 404) {
      console.log('User not found');
      return;
    }
    if (res.ok) {
      setGroupPosts(body);
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
        <GroupFeed />
      </div>
        </Show>
    </div> {/* Main grid */}
  </Layout></>)
}