import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { useParams } from '@solidjs/router';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import type {Group} from '~/types/Group';
import type {Post} from '~/types/Post';

  type GroupParams = {
    groupid: number;
  };

export default function Group(): JSXElement {

  const [targetGroup, setTargetGroup] = createSignal<Group | undefined>();
  
  const [groupPosts, setGroupPosts] = createSignal<Post[]>();
  const params: GroupParams = useParams();
  
  createEffect(() => {
  fetchWithAuth(config.API_URL + '/group/' + params.groupid).then(async (res) => {
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
  fetchWithAuth(config.API_URL + '/group/' + params.groupid).then(async (res) => {
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

  return <></>
}