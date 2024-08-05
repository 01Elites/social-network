import { JSXElement, onCleanup, useContext } from 'solid-js';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { cn } from '~/lib/utils';
import Repeat from '../core/repeat';
import { Skeleton } from '../ui/skeleton';

interface HomeEventsProps {
  class?: string;
}

export default function HomeEvents(props: HomeEventsProps): JSXElement {
  const wsCtx = useContext(WebSocketContext) as WebsocketHook;

  onCleanup(() => {});

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
