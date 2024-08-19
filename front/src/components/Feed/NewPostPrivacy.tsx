import { createEffect, createSignal, For, JSXElement } from 'solid-js';
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import { Button } from '../ui/button';
import { Checkbox } from '../ui/checkbox';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '../ui/dialog';
import { Separator } from '../ui/separator';
import { Table, TableBody, TableCell, TableRow } from '../ui/table';
import Friends from '~/types/friends';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';

interface NewPostPrivacyProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  onlyFollowersCallback: () => void;
  onlySelectedCallback: (selectedUsers: String[]) => void;
}

export default function NewPostPrivacy(props: NewPostPrivacyProps): JSXElement {
  const [selectedUsers, setSelectedUsers] = createSignal<String[]>([]);
  const [targetFriends, setTargetFriends] = createSignal<Friends | undefined>();

  createEffect(() => {
    // Fetch user Friends
    fetchWithAuth(config.API_URL + '/myfriends').then(async (res) => {
      const body = await res.json();
      if (res.ok) {
        setTargetFriends(body);
        return;
      } else {
        console.log('Error fetching friends');
        return;
      }
    });
  });

  function handleCheckboxChange(user: string) {
    if (selectedUsers().includes(user)) {
      setSelectedUsers(selectedUsers().filter((u) => u !== user));
    } else {
      setSelectedUsers([...selectedUsers(), user]);
    }
  }

  return (
    <Dialog open={props.open} onOpenChange={props.setOpen}>
      <DialogContent class=''>
        <DialogHeader>
          <DialogTitle>Who can see your post?</DialogTitle>
          <DialogDescription>
            Your post will be visible to everyone by default.
          </DialogDescription>
        </DialogHeader>

        <div class='max-h-[200px] overflow-x-auto'>
          <Table class=''>
            <TableBody>
              <For each={targetFriends()?.following ?? []}>
                {(user) => (
                  <TableRow>
                    <TableCell class='w-2'>
                      <Checkbox
                        checked={selectedUsers().includes(user.user_name)}
                        onChange={() => handleCheckboxChange(user.user_name)}
                        id={`npptchk-${user.user_name}`}
                      />
                    </TableCell>
                    <div onClick={() => handleCheckboxChange(user.user_name)}>
                      <TableCell class='w-2'>
                        <Avatar>
                          <AvatarImage
                            loading='lazy'
                            src={config.API_URL + '/image/' + user.avatar}
                          />
                          <AvatarFallback>{user.user_name[0].toUpperCase()}</AvatarFallback>
                        </Avatar>
                      </TableCell>
                      <TableCell class='cursor-pointer'>
                        {user.user_name}
                      </TableCell>
                    </div>
                  </TableRow>
                )}
              </For>
            </TableBody>
          </Table>
        </div>
        <Separator />

        <DialogFooter class='gap-2'>
          <Button
            variant='secondary'
            onClick={() => {
              props.onlyFollowersCallback();
              props.setOpen(false);
            }}
          >
            My Followers
          </Button>
          <Button
            disabled={selectedUsers().length < 1}
            onClick={() => {
              props.onlySelectedCallback(selectedUsers());
              props.setOpen(false);
            }}
          >
            Only Selected
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
