import { JSXElement, useContext } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { cn } from '~/lib/utils';
import { UserDetailsHook } from '~/types/User';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';

interface HomeEventsProps {
  class?: string;
}

export default function HomeEvents(props: HomeEventsProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Events</h1>
      <section class='flex flex-col gap-4 overflow-y-scroll'>
        <Repeat count={10}>
          <div>
            <Skeleton height={80} radius={10} />
          </div>
        </Repeat>
      </section>
    </section>
  );
}
