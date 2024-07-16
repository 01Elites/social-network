import { JSXElement } from 'solid-js';
import { Skeleton } from '../ui/skeleton';

export default function FeedPostCellSkeleton(): JSXElement {
  return (
    <div class='space-y-4'>
      <div class='flex items-center space-x-4'>
        <Skeleton height={40} circle animate={false} />
        <div class='w-full space-y-2'>
          <Skeleton height={16} radius={10} class='max-w-40' />
          <Skeleton height={16} radius={10} class='max-w-32' />
        </div>
      </div>
      <Skeleton height={150} radius={10} />
    </div>
  );
}
