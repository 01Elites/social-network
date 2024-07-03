import { JSXElement } from 'solid-js';
import { cn } from '~/lib/utils';
import NewPostCell from './NewPostCell';
import NewPostPreview from './NewPostPreview';

interface FeedProps {
  class?: string;
}

export default function Feed(props: FeedProps): JSXElement {
  return (
    <section class={cn('flex flex-col gap-3', props.class)}>
      <h1 class='text-xl font-bold'>Feed</h1>
      <NewPostCell />
      <NewPostPreview />
    </section>
  );
}
