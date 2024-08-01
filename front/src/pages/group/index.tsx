import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { useParams } from '@solidjs/router';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import type { Group } from '~/types/group';
import Layout from '~/Layout';
import GroupDetails from './details';
import GroupFeed from './groupfeed';
import NewGroupPostCell from "~/components/Feed/NewGroupPostCell";
import GroupContacts from "./groupcontacts";

export type GroupParams = {
  id: string;
};

export default function GroupPage(): JSXElement {

  const [targetGroup, setTargetGroup] = createSignal<Group | undefined>();

  const params: GroupParams = useParams();
  const groupID = params.id;

  createEffect(() => {
    fetchWithAuth(config.API_URL + '/group/' + groupID).then(async (res) => {
      console.log(config.API_URL + '/group/' + groupID)
      const body = await res.json();
      if (res.status === 200) {
        console.log('User fetched');
        console.log(body)
        setTargetGroup(body);
        return;
      } else {
        console.log('Error fetching group');
        return
      }
    })
  })
  return (<><Layout>
    <section class='flex h-full gap-4'>
    <div class='grid grid-cols-1 md:grid-cols-6 m-4'>
      <Show when={targetGroup()}>
        <div class='col-span-2'>
          <GroupDetails targetGroup={() => targetGroup() as Group} />
        </div>
      </Show>
      <Show when={targetGroup()?.ismember}>
        <div class='col-span-4 overflow-y-auto'>
          <NewGroupPostCell targetGroup={() => targetGroup() as Group} />
            <GroupFeed groupTitle={targetGroup()?.title}
            groupID={groupID as string}
              creator={targetGroup()?.iscreator as boolean}
              requesters={targetGroup()?.requesters}
              explore={targetGroup()?.explore} />
        </div>
      </Show>
    </div> 
      <GroupContacts members={targetGroup()?.members} class='hidden w-1/3 max-w-52 overflow-hidden md:flex' />
    </section>
  </Layout></>)
}