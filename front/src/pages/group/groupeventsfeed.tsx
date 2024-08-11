import {
  createEffect,
  createSignal,
  For,
  Index,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/types/User';
import { showToast } from '../../components/ui/toast';
import { GroupEvent } from '~/types/group/index';
import {EventsFeed} from "../events/eventsfeed";

interface FeedPostsProps {
  groupID: string;
}

export default function GroupEventsFeed(props: FeedPostsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [events, setEvents] = createSignal<GroupEvent[]>(); 
  createEffect(() => {
    if (!userDetails()) return;
    fetchWithAuth(config.API_URL + `/group/${props.groupID}/events`)
      .then(async (res) => {
        const body = await res.json();
        if (res.status === 404) {
          setEvents([]);
          return;
        }
        if (res.ok) {
          console.log(body);
          setEvents(body);
          return;
        }
        throw new Error(
          body.reason ? body.reason : 'An error occurred while fetching events',
        );
      })
      .catch((err) => {
        showToast({
          title: 'Error fetching events',
          description: err.message,
          variant: 'error',
        });
      });
  });
  return (<>
  < EventsFeed events={events()}/>
  </>
)}
