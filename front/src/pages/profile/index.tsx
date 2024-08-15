import { useParams } from '@solidjs/router';
import {
  createEffect,
  createSignal,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/hooks/userDetails';
import Layout from '~/Layout';
import User from '~/types/User';
import ProfileFeed from './proFeed';
import ProfileDetails from './profileDetails';

type ProfileParams = {
  username: string;
};

export default function Profile(): JSXElement {
  const userCtx = useContext(UserDetailsContext) as UserDetailsHook;
  const [targetUser, setTargetUser] = createSignal<User | undefined>();

  const params: ProfileParams = useParams();

  createEffect(() => {
    console.log('Profile page mounted');

    if (
      userCtx.userDetails() !== null &&
      userCtx.userDetails()!.user_name !== null &&
      userCtx.userDetails()!.user_name === params.username
    ) {
      setTargetUser(userCtx.userDetails()!);
      return;
    }
    // Fetch user profile details
    fetchWithAuth(config.API_URL + '/profile/' + params.username).then(
      async (res) => {
        const body = await res.json();
        if (res.status === 404) {
          console.log('User not found');
          return;
        }
        if (res.ok) {
          setTargetUser(body);
          return;
        }
      },
    );
  });

  return (
    <Layout>
      <div class='m-4 grid grid-cols-1 md:grid-cols-6'>
        <div class='col-span-2'>
          <Show when={targetUser()}>
            <ProfileDetails targetUser={() => targetUser() as User} />
          </Show>
        </div>
        <div class='col-span-4 overflow-y-auto'>
          <ProfileFeed targetUser={() => targetUser() as User} />
        </div>
      </div>{' '}
      {/* Main grid */}
    </Layout>
  );
}
