import 'solid-devtools';
import { createEffect, createSignal, JSXElement } from 'solid-js';
import Layout from '~/Layout';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import Friends from '~/types/friends';

// type FriendsParams = {
//   username: string;
// };

export default function FriendsPage(): JSXElement {
  const [targetFriends, setTargetFriends] = createSignal<Friends | undefined>();

  // const params: FriendsParams = useParams();
  // console.log('/friends/' + params.username);

  createEffect(() => {
    // Fetch user Friends
    fetchWithAuth(config.API_URL + '/myfriends').then(async (res) => {
      console.log(config.API_URL + '/myfriends');

      const body = await res.json();
      if (res.ok) {
        setTargetFriends(body);
        console.log(body);
        return;
      } else {
        console.log('Error fetching friends');
        return;
      }
    });
  });

  return (
    <Layout>
      <section class='flex h-full gap-4'>
        <h1>Friends</h1>
      </section>
    </Layout>
  );
}
