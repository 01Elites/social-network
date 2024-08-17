import { JSXElement, useContext, createEffect, createSignal, For, Show, Setter } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { Avatar, AvatarFallback, AvatarImage } from '../../components/ui/avatar';
import config from '../../config';
import { Card } from '../../components/ui/card';
import { UserDetailsHook } from '~/hooks/userDetails';
import { ChatState } from '~/pages/home';
import Contact from '~/types/Contact';

interface HomeContactsProps {
  class?: string;
  setChatState?: Setter<ChatState>
}


interface Section {
  name: string;
  users: Contact[];
}

export default function HomeContacts(props: HomeContactsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const wsCtx = useContext(WebSocketContext) as WebsocketHook;
  const [sections, setSections] = createSignal<Section[]>([]);

  createEffect(() => {
    if (wsCtx.state() === 'connected') {
      wsCtx.send({ event: 'USERLIST', payload: null });
    }
  });

  createEffect(() => {
    wsCtx.bind('USERLIST', (section: Section) => {
      console.log('Received userlist', section);
      if (!section) {
        return;
      }
      if (section.users === null || section.users.length === 0) {
        return;
      }

      setSections(prev => {
        const newSections = prev.filter(s => s.name !== section.name);
        newSections.push(section);
        return newSections;
      })
    });
  })


  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Contacts</h1>
      <section class='flex flex-col gap-2 overflow-y-scroll'>
        {/* Skeleton for guests */}
        <Show when={!userDetails()}>
          <Repeat count={10}>
            <div class='flex items-center gap-4'>
              <Skeleton height={40} circle animate={false} />
              <div class='grow space-y-2'>
                <Skeleton height={14} radius={10} />
                <Skeleton height={12} radius={10} class='max-w-20' />
              </div>
            </div>
          </Repeat>
        </Show>
        <Show when={userDetails() && sections().length !== 0} fallback={
          <>
            {userDetails() && (
              <>
                <h1 class='font-semibold text-muted-foreground'>
                  Hmmm, you don't have any friends added yet :(
                </h1>
                <p class='text-muted-foreground'>
                  Maybe you could add some
                </p>
              </>
            )}
          </>
        }>
          <For each={sections()}>
            {(section) => (
              <>
                <h1 class='text-md font-semibold text-primary/80'>{section.name}</h1>
                < Show when={section.users.length > 0} fallback={
                  <Repeat count={4}>
                    <div class='flex items-center gap-3'>
                      <Skeleton height={40} circle animate={false} />
                      <div class='grow space-y-2'>
                        <Skeleton height={14} radius={10} />
                        <Skeleton height={12} radius={10} class='max-w-20' />
                      </div>
                    </div>
                  </Repeat>
                }>
                  <For each={section.users}>{(user) => (
                    // set chat state when a user is clicked
                    <Card class='cursor-pointer' onClick={() => {
                      if (props.setChatState != null) {
                        console.log('Opening chat with', user.user_name);
                        props.setChatState({
                          isOpen: true,
                          chatWith: user.user_name
                        });
                      }
                    }}>
                      <div class='flex items-center gap-3 relative'>
                        <div class='relative'>
                          <Avatar>
                            <AvatarImage src={`${config.API_URL}/image/${user.avatar}`} />
                            <AvatarFallback>{user.first_name.charAt(0).toUpperCase()}</AvatarFallback>
                          </Avatar>
                          <div class={cn('absolute top-0 right-0 w-3 h-3 rounded-full z-10', user.state === 'online' ? 'bg-green-500' : 'bg-red-500')}></div>
                        </div>
                        <div class='grow space-y-2'>
                          <div class='flex flex-col items-center justify-center space-x-1'>
                            <div>{user.first_name} {user.last_name}</div>
                            <div>{user.user_name}</div>
                          </div>
                        </div>
                      </div>
                    </Card>
                  )}</For>
                </Show>
              </>
            )}
          </For>
        </Show>


      </section>
    </section >
  );
}
