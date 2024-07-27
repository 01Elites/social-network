import 'solid-devtools';
import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import Layout from '~/Layout';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import Friends from '~/types/friends';
import FriendsFeed from './friendsFeed';

export default function FriendsPage(): JSXElement {
  const [targetFriends, setTargetFriends] = createSignal<Friends | undefined>();

  createEffect(() => {
    // Fetch user Friends
    fetchWithAuth(config.API_URL + '/myfriends').then(async (res) => {
      const body = await res.json();
      if (res.ok) {
        setTargetFriends(body);
        return;
      } else {
        console.log('Error fetching friends');
        return;
      }
    });
  });

  return (
    <Layout>
      <section class='flex h-full flex-col gap-4'>
        <h1>Friends</h1>
        <Show when={targetFriends()}>
          <div class='m-4 grid grid-cols-1'>
            <FriendsFeed targetFriends={() => targetFriends() as Friends} />
          </div>
        </Show>
      </section>
    </Layout>
  );
}
