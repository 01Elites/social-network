import { createSignal, JSXElement } from 'solid-js';
import { Button } from '~/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';
import { TextField, TextFieldTextArea } from '~/components/ui/text-field';
import { showToast } from '~/components/ui/toast';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';

interface NewGroupPreviewProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

export default function NewGroupPreview(
  props: NewGroupPreviewProps,
): JSXElement {
  const [groupName, setGroupName] = createSignal<string>('');
  const [groupDescription, setGroupDescription] = createSignal<string>('');
  const [formProcessing, setFormProcessing] = createSignal(false);

  async function createGroup() {
    setFormProcessing(true);

    const payload = {
      title: groupName(),
      description: groupDescription(),
    };

    fetchWithAuth(config.API_URL + '/create_group', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
      .then(async (response) => {
        setFormProcessing(false);
        if (!response.ok) {
          const errMsg = await response.json();
          showToast({
            title: 'Could not create group',
            description: errMsg.reason
              ? errMsg.reason
              : 'An error occurred while creating the group',
            variant: 'error',
          });
        } else {
          showToast({
            title: 'Group created',
            description: 'Your group has been created successfully',
            variant: 'success',
          });
          props.setOpen(false);
          window.location.reload();
        }
      })
      .catch((error) => {
        setFormProcessing(false);
        console.error('Error creating group:', error);
        showToast({
          title: 'Could not create group',
          description: 'An error occurred while creating the group',
          variant: 'error',
        });
      });
  }

  return (
    <Dialog open={props.open} onOpenChange={props.setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create New Group</DialogTitle>
          <DialogDescription>
            Share a brief description and title for your new group.
          </DialogDescription>
        </DialogHeader>

        <TextField value={groupName()} onChange={(e) => setGroupName(e)}>
          <TextFieldTextArea
            placeholder='Group Name'
            class='resize-none'
            minLength={1}
            maxLength={13}
            disabled={formProcessing()}
          />
        </TextField>

        <TextField
          value={groupDescription()}
          onChange={(e) => setGroupDescription(e)}
        >
          <TextFieldTextArea
            placeholder='Group Description'
            class='resize-none'
            minLength={1}
            maxLength={200}
            disabled={formProcessing()}
          />
        </TextField>

        <DialogFooter class='!justify-between gap-4'>
          <Button
            class='gap-2'
            disabled={
              formProcessing() ||
              !groupName().trim() ||
              !groupDescription().trim()
            }
            onClick={createGroup}
          >
            {formProcessing() ? 'Creating...' : 'Create Group'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
