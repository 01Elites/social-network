import { useParams } from '@solidjs/router';
import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import Layout from '~/Layout';
import ProfileFeed from './proFeed';
import ProfileDetails from './profileDetails';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import User from '~/types/User';

type ProfileParams = {
  username: string;
};

export default function Profile(): JSXElement {

  const [targetUser, setTargetUser] = createSignal<User | undefined>();

  const params: ProfileParams = useParams();

  createEffect(() => {
    console.log('Profile page mounted');

    // Fetch user profile details
    fetchWithAuth(config.API_URL + '/profile/' + params.username).then(async (res) => {
      const body = await res.json();
      if (res.status === 404) {
        console.log('User not found');
        return;
      }
      if (res.ok) {
        setTargetUser(body);
        console.log(body);
        return;
      }

    })
  })

  return (
    <Layout>
      <div class='grid grid-cols-1 md:grid-cols-6 m-4 '> {/* Main grid */}
        <div class='col-span-2'>
          <Show when={targetUser()}>
            <ProfileDetails targetUser={() => targetUser() as User} />
          </Show>
        </div>
        <div class='col-span-4 overflow-y-auto'>
          <Show when={targetUser()?.follow_status === "following"}>
          <ProfileFeed />
          </Show>
        </div>
      </div> {/* Main grid */}
    </Layout>
  )
}
