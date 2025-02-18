import { JSXElement, useContext, createEffect, createSignal, For, Show, Setter, onCleanup } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { Avatar, AvatarFallback, AvatarImage } from '../../components/ui/avatar';
import config from '../../config';
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

  const handleUserListEvent = (data: { type: string; metadata: any; }) => {
    if (!data) return;

    switch (data.type) {
      case 'init':
        handleUserListInit(data.metadata);
        break;
      case 'update':
        handleUserListUpdate(data.metadata);
        break;
      case 'add':
        handleUserListAdd(data.metadata);
        break;
      default:
        console.warn('Unknown data type:', data.type);
    }
  };
  // Function to handle 'init' event
  function handleUserListInit(section: { name: string; users: Contact[]; }) {
    if (!section) return;
    if (section.users == null || section.users.length === 0) return;
    setSections(prev => {
      const newSections = prev.filter(s => s.name !== section.name);
      newSections.push(section);
      return newSections;
    })
  }

  // Function to handle 'update' event
  function handleUserListUpdate(metadata: { user_name: string; state: 'online' | 'offline' }) {
    if (!metadata || !metadata.user_name || !metadata.state) return;
    setSections(prevSections =>
      prevSections.map(section => ({
        ...section,
        users: section.users.map(user =>
          user.user_name === metadata.user_name
            ? { ...user, state: metadata.state }
            : user
        )
      }))
    );
  }

  // Function to handle 'add' event
  function handleUserListAdd(metadata: { name: string; user: Contact }) {
    if (!metadata || !metadata.name || !metadata.user) return;
    setSections(prevSections => {
      return prevSections.map(section =>
        section.name === metadata.name
          ? { ...section, users: [...section.users, metadata.user] }
          : section
      );
    });
  }

  createEffect(() => {
    if (wsCtx.state() === 'connected') {
      wsCtx.send({ event: 'USERLIST', payload: null });
    }
  });

  createEffect(() => {
    if (userDetails == null || wsCtx == null) return;
    const unbindUserList = wsCtx.bind('USERLIST', handleUserListEvent);

    onCleanup(unbindUserList);
  });

  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold hidden md:flex'>Contacts</h1>
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
                <h1 class='text-md font-semibold text-primary/80 hidden md:flex'>{section.name}</h1>
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
                    <div class='flex items-center gap-3 relative cursor-pointer select-none hover:bg-secondary/80 rounded-md p-2'
                      onClick={() => {
                        if (props.setChatState != null) {
                          props.setChatState!({
                            isOpen: false,
                            chatWith: '',
                          });
                          wsCtx.send({
                            event: "CHAT_CLOSED",
                            payload: {}
                          });

                          props.setChatState({
                            isOpen: true,
                            chatWith: user.user_name
                          });
                        }
                      }}>
                      <div class='relative'>
                        <Avatar>
                          <AvatarImage src={`${config.API_URL}/image/${user.avatar}`} />
                          <AvatarFallback>{user.first_name.charAt(0).toUpperCase()}</AvatarFallback>
                        </Avatar>
                        <div class={cn('absolute bottom-0 -right-1 size-4 rounded-full border-2 border-background', user.state === 'online' ? 'bg-green-500' : 'bg-red-500')}></div>
                      </div>
                      <div class='grow space-y-2'>
                        <div class='flex flex-col items-start justify-center'>
                          <h3 class='font-semibold text-sm text-primary/90'>{user.first_name} {user.last_name}</h3>
                          <h4 class='font-medium text-xs text-primary/90'>{user.user_name}</h4>
                        </div>
                      </div>
                    </div>
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
