import { createSignal, For, Index, useContext } from 'solid-js';
import { JSXElement } from 'solid-js';
import { Tabs } from '@kobalte/core/tabs';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Card } from '~/components/ui/card';
import Groups from '~/types/groups';
import NewGroupPreview from '~/components/Feed/NewGroupPreview';
import { Button } from '~/components/ui/button';
import RequestToJoin from "../group/request";
import { Group } from '~/types/group';
import { UserDetailsHook } from '~/hooks/userDetails';
import UserDetailsContext from '~/contexts/UserDetailsContext';

export default function GroupsFeed(props: { targetGroups: () => Groups | undefined }): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const groups = props.targetGroups();
  // console.log("groups", groups);
  const [groupPreviewOpen, setGroupPreviewOpen] = createSignal(false);

  return (
    <div >
    <div class='flex gap-2 rounded border-[1px] p-2'>
    <Avatar>
        <AvatarImage src={userDetails()?.avatar} />
        <AvatarFallback>
          {userDetails()?.first_name.charAt(0).toUpperCase()}
        </AvatarFallback>
      </Avatar>
        <NewGroupPreview setOpen={setGroupPreviewOpen} open={groupPreviewOpen()} />
        <Button
        variant='ghost'
        class='w-full justify-start text-muted-foreground'
        onClick={() => setGroupPreviewOpen(true)}
      >
        Create New Group
      </Button>
      </div>
      <Tabs aria-label="Main navigation" class="tabs">
        <Tabs.List class="tabs__list">
          <Tabs.Trigger class="tabs__trigger" value="owned">
            Owned ({groups?.owned?.length || 0})
          </Tabs.Trigger>
          <Tabs.Trigger class="tabs__trigger" value="joined">
            Joined ({groups?.joined?.length || 0})
          </Tabs.Trigger>
          <Tabs.Trigger class="tabs__trigger" value="explore">
            Explore ({groups?.explore?.length || 0})
          </Tabs.Trigger>

          <Tabs.Indicator class="tabs__indicator" />
        </Tabs.List>

        <Tabs.Content class="tabs__content  m-6 flex flex-wrap gap-4" value="owned">
          <For each={groups?.owned ?? []}>
            {(group) => (
              <Card class="flex w-44 flex-col items-center space-y-4 p-3">
                <a
                  href={`/group/${group.id}`}
                  class="flex flex-col items-center text-base font-bold text-blue-500"
                >
                  <Avatar class="mb-3 h-20 w-20">
                    <AvatarImage />
                    <AvatarFallback>{group.title.charAt(0).toUpperCase()}</AvatarFallback>
                  </Avatar>
                  {group.title}
                </a>
              </Card>
            )}
          </For>
        </Tabs.Content>

        <Tabs.Content class="tabs__content  m-6 flex flex-wrap gap-4" value="joined">
          <For each={groups?.joined ?? []}>
            {(group) => (
              <Card class="flex w-44 flex-col items-center space-y-4 p-3">
                <a
                  href={`/group/${group.id}`}
                  class="flex flex-col items-center text-base font-bold text-blue-500"
                >
                  <Avatar class="mb-3 h-20 w-20">
                    <AvatarImage />
                    <AvatarFallback>{group.title.charAt(0).toUpperCase()}</AvatarFallback>
                  </Avatar>
                  {group.title}
                </a>
                <RequestToJoin targetGroup={() => group as Group} />
              </Card>
            )}
          </For>
        </Tabs.Content>

        <Tabs.Content class="tabs__content  m-6 flex flex-wrap gap-4" value="explore">
          <For each={groups?.explore ?? []}>
            {(group) => (
              <Card class="flex w-44 flex-col items-center space-y-4 p-3">
                <a
                  href={`/group/${group.id}`}
                  class="flex flex-col items-center text-base font-bold text-blue-500"
                >
                  <Avatar class="mb-3 h-20 w-20">
                    <AvatarImage />
                    <AvatarFallback>{group.title.charAt(0).toUpperCase()}</AvatarFallback>
                  </Avatar>
                  {group.title}
                </a>
                <RequestToJoin targetGroup={() => group as Group} />
              </Card>
            )}
          </For>
        </Tabs.Content>
      </Tabs>
    </div>
  );
}
