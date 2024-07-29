import { createSignal, For } from 'solid-js';
import { JSXElement } from 'solid-js';
import { Tabs } from '@kobalte/core/tabs';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Card } from '~/components/ui/card';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import Groups from '~/types/groups';

export default function GroupsFeed(props: { targetGroups: () => Groups | undefined }): JSXElement {
  const groups = props.targetGroups();
  console.log("groups", groups);
  const [title, setTitle] = createSignal('');
  const [description, setDescription] = createSignal('');
  const [formProcessing, setFormProcessing] = createSignal(false);

  const submitGroupData = async () => {
    setFormProcessing(true);

    const groupData = {
      title: title(),
      description: description(),
    };

    fetchWithAuth(`${config.API_URL}/create_group`, {
      method: 'POST',
      body: JSON.stringify(groupData),
    })
      .then(async (res) => {
        setFormProcessing(false);
        if (res.status === 201) {
          alert('Group created successfully');
          return;
        }

        const error = await res.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error('An error occurred while creating the group. Please try again.');
      })
      .catch((error) => {
        alert(`Error: ${error.message}`);
      })
      .finally(() => {
        setFormProcessing(false);
      });
  };

  const handleSubmit = (event: Event) => {
    event.preventDefault();
    submitGroupData();
  };

  return (
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
        <Tabs.Trigger class="tabs__trigger" value="create">
          Create
        </Tabs.Trigger>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>

      <Tabs.Content class="m-6 flex flex-wrap gap-4" value="owned">
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

      <Tabs.Content class="m-6 flex flex-wrap gap-4" value="joined">
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
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class="m-6 flex flex-wrap gap-4" value="explore">
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
            </Card>
          )}
        </For>
      </Tabs.Content>

      <Tabs.Content class="m-6 flex flex-col space-y-4" value="create">
        <form onSubmit={handleSubmit} class="flex flex-col space-y-4">
          <label class="flex flex-col">
            <span class="text-sm font-medium">Title</span>
            <input
              type="text"
              value={title()}
              onInput={(e) => setTitle(e.currentTarget.value)}
              class="p-2 border rounded"
              required
            />
          </label>
          <label class="flex flex-col">
            <span class="text-sm font-medium">Description</span>
            <textarea
              value={description()}
              onInput={(e) => setDescription(e.currentTarget.value)}
              class="p-2 border rounded"
              rows="4"
              required
            />
          </label>
          <button type="submit" class="bg-blue-500 text-white p-2 rounded" disabled={formProcessing()}>
            Create Group
          </button>
        </form>
      </Tabs.Content>
    </Tabs>
  );
}
