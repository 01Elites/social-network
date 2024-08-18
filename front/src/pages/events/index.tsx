import 'solid-devtools';
import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import Layout from '~/Layout';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import { GroupEvent } from '~/types/group'; 
import { Tabs } from '@kobalte/core/tabs';
import { EventsFeed } from './groupeventsfeed';


function EventsPage(): JSXElement {
  const [pendingEvents, setPendingEvents] = createSignal<GroupEvent[] | undefined>();
  const [upcomingEvents, setUpcomingEvents] = createSignal<GroupEvent[] | undefined>();
  const [pastEvents, setPastEvents] = createSignal<GroupEvent[] | undefined>();

  createEffect(() => {
    // Fetch user Friends
    fetchWithAuth(config.API_URL + '/myEvents').then(async (res) => {
      const body = await res.json();
      if (res.ok) {
        setPendingEvents(body.pending)
        setUpcomingEvents(body.upcoming)
        setPastEvents(body.past)
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
        <h1 class='text-xl font-bold'>Events</h1>
        <Tabs aria-label='Main navigation' class='tabs'>
      <Tabs.List class='tabs__list'>
        <Tabs.Trigger class='tabs__trigger' value='pending'>
          Pending ({pendingEvents()?.length || 0})
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='upcoming'>
          Upcoming ({upcomingEvents?.length || 0})
        </Tabs.Trigger>
        <Tabs.Trigger class='tabs__trigger' value='past'>
          Past ({pastEvents()?.length || 0})
        </Tabs.Trigger>

        <Tabs.Indicator class='tabs__indicator' />
      </Tabs.List>

      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='pending'>
      < EventsFeed events={pendingEvents()} />
      </Tabs.Content>
      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='upcoming'>
      < EventsFeed events={upcomingEvents()} />
      </Tabs.Content>
      <Tabs.Content class='m-6 flex flex-wrap gap-4' value='past'>
      < EventsFeed events={pastEvents()} />

      </Tabs.Content>

    </Tabs>
      </section>
    </Layout>
  );
}

export default EventsPage;


// export default function FriendsPage(): JSXElement {


//   return (
//     <Layout>
//       <section class='flex h-full flex-col gap-4'>
//         <h1 class='text-xl font-bold'>Friends</h1>
//         <Show when={targetFriends()}>
//           <div class='grid grid-cols-1'>
//             <FriendsFeed targetFriends={() => targetFriends() as Friends} />
//           </div>
//         </Show>
//       </section>
//     </Layout>
//   );
// }

