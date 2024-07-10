import { JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import { UserDetailsHook } from '~/types/User';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';

interface ContactsProps {
  class?: string;
}

export default function Contacts(props: ContactsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Contacts</h1>
      <section class='flex flex-col gap-4 overflow-y-scroll'>
        <Repeat count={20}>
          <div class='flex items-center gap-3'>
            <Skeleton height={40} circle animate={false} />
            <div class='grow space-y-2'>
              <Skeleton height={14} radius={10} class='max-w-32' />
              <Skeleton height={12} radius={10} class='max-w-20' />
            </div>
          </div>
        </Repeat>
      </section>
    </section>
  );
}
