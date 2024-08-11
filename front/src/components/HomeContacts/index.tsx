import { JSXElement, useContext, createEffect, createSignal, For, Show } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { Avatar, AvatarFallback } from '../../components/ui/avatar';
import config from '../../config';
import { Card } from '../../components/ui/card';
import { UserDetailsHook } from '~/hooks/userDetails';

interface HomeContactsProps {
  class?: string;
}

interface Contact {
  first_name: string;
  last_name: string;
  username: string;
  state: string;
  avatar: string;
}

interface Section {
  name: string;
  users: Contact[];
}

export default function HomeContacts(props: HomeContactsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const wsCtx = useContext(WebSocketContext) as WebsocketHook;
  const [sections, setSections] = createSignal<Section[]>([
    { name: 'Following', users: [] },
    { name: 'Direct Messages', users: [] }
  ]);
  if (userDetails != null && wsCtx != null) {
    createEffect(() => {
      wsCtx.bind('USERLIST', (data) => {
        if (data != null) {
          setSections((prevSections) => {
            const updatedSections = [...prevSections];
            const sectionIndex = updatedSections.findIndex(section => section.name === data.name);

            if (sectionIndex > -1) {
              // Create a map for updating users within the section
              const userMap = new Map(updatedSections[sectionIndex].users.map(user => [user.username, user]));

              // Update or add new users
              data.users.forEach(user => {
                userMap.set(user.username, user);
              });

              // Update the section with the new list of users
              updatedSections[sectionIndex] = {
                name: data.name,
                users: Array.from(userMap.values())
              };
            } else {
              // If the section does not exist, add it
              updatedSections.push({
                name: data.name,
                users: data.users
              });
            }

            return updatedSections;
          });
        }
      });
    })
  };

  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Contacts</h1>
      <section class='flex flex-col gap-2 overflow-y-scroll'>
        <For each={sections()}>
          {(section) => (
            <>
              <h1 class='text-xl'>{section.name}</h1>
              <Show when={section.users.length > 0} fallback={
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
                <For each={section.users}>{(item) => (
                  <Card class='bg-primary/10'>
                    <div class='flex items-center gap-3 relative'>
                      <div class='relative'>
                        <Avatar>
                          <AvatarFallback>
                            <Show when={item.avatar} fallback={item.first_name.charAt(0).toUpperCase()}>
                              <img
                                alt='avatar'
                                class='size-full rounded-md rounded-b-none object-cover'
                                loading='lazy'
                                src={`${config.API_URL}/image/${item.avatar}`}
                              />
                            </Show>
                          </AvatarFallback>
                        </Avatar>
                        <div class={cn('absolute top-0 right-0 w-3 h-3 rounded-full z-10', {
                          'bg-green-500': item.state === 'online',
                          'bg-red-500': item.state !== 'online',
                        })}></div>
                      </div>
                      <div class='grow space-y-2'>
                        <div class='flex flex-col items-center justify-center space-x-1'>
                          <div>{item.first_name} {item.last_name}</div>
                          <div>{item.username}</div>
                        </div>
                      </div>
                    </div>
                  </Card>
                )}</For>
              </Show>
            </>
          )}
        </For>
      </section>
    </section>
  );
}
