import { Tabs } from "@kobalte/core/tabs";
import { JSXElement, useContext } from 'solid-js';
import Repeat from '~/components/core/repeat';
import "./style.css";
import { Skeleton } from '~/components/ui/skeleton';
import { cn } from "~/lib/utils";

export default function ProfileFeed(): JSXElement {

  return (
    <Tabs aria-label="Main navigation" class="tabs">
      <Tabs.List class="tabs__list">
        <Tabs.Trigger class="tabs__trigger" value="posts">Posts</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill">ما ادري ويش</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill2">وش تمبا</Tabs.Trigger>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="posts">
        <div class={cn('flex flex-col gap-4 p-2')}>
          <Repeat count={10}>
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
          </Repeat>
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill2">still NOTHING!!!</Tabs.Content>
    </Tabs>
  )
}

