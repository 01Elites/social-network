import 'solid-devtools';
import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import Layout from '~/Layout';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import Groups from '~/types/groups';
import GroupsFeed from './groupsFeed';

export default function GroupsPage(): JSXElement {
  const [targetGroups, setTargetGroups] = createSignal<Groups | undefined>();

  createEffect(() => {
    // Fetch user Groups
    fetchWithAuth(config.API_URL + '/mygroups').then(async (res) => {
      const body = await res.json();
      if (res.ok) {
        setTargetGroups(body);
        return;
      } else {
        console.log('Error fetching groups');
        return;
      }
    });
  });
  return (

    <Layout>
      <section class='flex h-full flex-col gap-4'>
        <h1 class='text-xl font-bold'>Groups</h1>
        <Show when={targetGroups()}>
          <div class='grid grid-cols-1'>
            <GroupsFeed targetGroups={targetGroups} />
          </div>
        </Show>
      </section>
    </Layout>
  );
}
