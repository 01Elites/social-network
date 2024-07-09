import { createSignal, For, JSXElement } from 'solid-js';
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

interface NewPostPrivacyProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  onlyFollowersCallback: () => void;
  onlySelectedCallback: (selectedUsers: String[]) => void;
}

const fakeUsers = [
  'John Doe',
  'Jane Smith',
  'Michael Johnson',
  'Emily Brown',
  'William Davis',
  'Sophia Miller',
  'James Wilson',
  'Olivia Moore',
  'Benjamin Taylor',
  'Ava Anderson',
  'Daniel Thomas',
  'Mia Jackson',
  'Matthew White',
  'Isabella Harris',
  'David Martin',
  'Emma Thompson',
  'Andrew Garcia',
  'Charlotte Martinez',
  'Joseph Robinson',
  'Madison Clark',
];

export default function NewPostPrivacy(props: NewPostPrivacyProps): JSXElement {
  const [selectedUsers, setSelectedUsers] = createSignal<String[]>([]);

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
              <For each={fakeUsers}>
                {(user) => (
                  <TableRow>
                    <TableCell class='w-2'>
                      <Checkbox
                        onChange={() => {
                          if (selectedUsers().includes(user)) {
                            setSelectedUsers(
                              selectedUsers().filter((u) => u !== user),
                            );
                          } else {
                            setSelectedUsers([...selectedUsers(), user]);
                          }
                        }}
                        id={`npptchk-${user}`}
                      />
                    </TableCell>
                    <TableCell class='w-2'>
                      <Avatar>
                        <AvatarImage
                          loading='lazy'
                          src='https://thispersondoesnotexist.com'
                        />
                        <AvatarFallback>{user[0].toUpperCase()}</AvatarFallback>
                      </Avatar>
                    </TableCell>
                    <TableCell
                      class='cursor-pointer'
                      onClick={() => {
                        document
                          .getElementById(`npptchk-${user}-input`)
                          ?.click();
                      }}
                    >
                      {user}
                    </TableCell>
                  </TableRow>
                )}
              </For>
            </TableBody>
          </Table>
        </div>

        <Separator />

        <DialogFooter class='gap-2'>
          <Button variant='secondary' onClick={props.onlyFollowersCallback}>
            My Followers
          </Button>
          <Button
            disabled={selectedUsers().length < 1}
            onClick={() => props.onlySelectedCallback(selectedUsers())}
          >
            Only Selected
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
