import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { useParams } from '@solidjs/router';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import type { Group } from '~/types/group';
import Layout from '~/Layout';
import GroupDetails from './details';
import GroupFeed from './groupfeed';
import { showToast } from '~/components/ui/toast';

export type GroupParams = {
  id: string;
};

export default function GroupPage(): JSXElement {

  const [targetGroup, setTargetGroup] = createSignal<Group | undefined>();

  const params: GroupParams = useParams();
  const groupID = params.id;

  createEffect(() => {
    fetchWithAuth(config.API_URL + '/group/' + groupID).then(async (res) => {
      const body = await res.json();
      if (res.status === 200) {
        console.log(body)
        setTargetGroup(body);
        return;
      } else {
        showToast({ title: 'Error', description: 'Failed to fetch group', variant: 'error' });
        return
      }
    })
  })
  return (<><Layout>
    <div class='grid grid-cols-1 md:grid-cols-6 m-4 space-x-4'>
      <Show when={targetGroup()}>
        <div class='col-span-1 mt-20'>
          <GroupDetails targetGroup={() => targetGroup() as Group} />
        </div>
      </Show>
      <Show when={targetGroup()?.ismember}>
        <div class='col-span-5 overflow-y-auto'>
          <GroupFeed groupTitle={targetGroup()?.title}
            groupID={groupID as string}
            creator={targetGroup()?.iscreator as boolean}
            requesters={targetGroup()?.requesters}
            explore={targetGroup()?.explore}
            members={targetGroup()?.members} />
        </div>
      </Show>
    </div>
  </Layout></>)
}