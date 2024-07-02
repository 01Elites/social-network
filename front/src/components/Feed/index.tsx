import { JSXElement } from 'solid-js';
import { cn } from '~/lib/utils';

interface FeedProps {
  class?: string;
}

export default function Feed(props: FeedProps): JSXElement {
  return (
    <section class={cn(props.class)}>
      <h1 class='text-lg font-bold'>Feed</h1>
    </section>
  );
}
