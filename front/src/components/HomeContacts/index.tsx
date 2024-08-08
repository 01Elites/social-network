import { JSXElement, useContext, createEffect, createSignal, For } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import { UserDetailsHook } from '~/types/User';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import PostAuthorCell from '../PostAuthorCell';

interface HomeContactsProps {
  class?: string;
}

interface Contact {
  first_name: string
  last_name: string,
  username: string,
  state: string
}

export default function HomeContacts(props: HomeContactsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const useWebsocket = useContext(WebSocketContext) as WebsocketHook;
  const [contacts, setContacts] = createSignal<Contact[] | undefined>()
  const [loadOnce, setLoadOnce] = createSignal(0)

  createEffect(() => {
    useWebsocket.bind("USERLIST", (data) => {
      if (data != null && loadOnce() == 0) {
        setContacts(data.users)
        setLoadOnce(1)
      }
    })
  })
  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Contacts</h1>
      <section class='flex flex-col gap-4 overflow-y-scroll'>
        <For each={contacts()}>{(item) => <li>{item.first_name} {item.last_name}</li>}</For>

        <Repeat count={20}>
          <div class='flex items-center gap-3'>
            <Skeleton height={40} circle animate={false} />
            <div class='grow space-y-2'>
              <Skeleton height={14} radius={10} />
              <Skeleton height={12} radius={10} class='max-w-20' />
            </div>
          </div>
        </Repeat>
      </section>
    </section>
  );
}
