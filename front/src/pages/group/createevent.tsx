import { createEffect, createSignal, JSXElement, useContext } from 'solid-js';
import { Button } from '~/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';
import { Separator } from '~/components/ui/separator';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import User, { UserDetailsHook } from '~/types/User';
import PostAuthorCell from '../../components/PostAuthorCell';
import { AspectRatio } from '../../components/ui/aspect-ratio';
import { TextField,
   TextFieldTextArea,
    TextFieldInput,
  TextFieldLabel } from '../../components/ui/text-field';
import moment from 'moment';
import tailspin from '~/assets/svg-loaders/tail-spin.svg';
import config from '~/config';
import { fetchWithAuth } from '~/extensions/fetch';
import { showToast } from '../../components/ui/toast';

interface NewPostPreviewProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  groupID: string
  groupTitle: string | undefined;
}

export default function CreateEvent(props: NewPostPreviewProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [eventDescription, setEventDescription] = createSignal<string>('');
  const [title, setTitle] = createSignal<string>('');
  const [firstoption, setfirstoption] = createSignal<string>('');
  const [secondoption, setsecondoption] = createSignal<string>('');
  const [eventTime, setEventTime] = createSignal<string>('');
  const [selectedOptions, setSelectedOptions] = createSignal<String[]>([]);

  const [formProcessing, setFormProcessing] = createSignal(false);

  async function makeEvent() {
    setFormProcessing(true);
    if (moment(eventTime()).isBefore(moment().toISOString())) {
      showToast({
        title: 'Invalid date',
        description: 'Event date should be in the future',
        variant: 'error',
      });
      setFormProcessing(false);
      return;
    }
    const payload = {
      group_id: Number(props.groupID),
      title: title(),
      description: eventDescription(),
      options: [firstoption(), secondoption()],
      event_date: new Date(eventTime()).toISOString(),
    };

    fetchWithAuth(config.API_URL + '/create_event', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
      .then(async (response) => {
        setFormProcessing(false);
        if (!response.ok) {
          const errMsg = await response.json();
          showToast({
            title: 'Could not Create Event',
            description: errMsg.reason
              ? errMsg.reason
              : 'An error occurred while posting your content',
            variant: 'error',
          });
        } else {
          props.setOpen(false);
        }
      })
      .catch((error) => {
        setFormProcessing(false);
        console.error('Error posting:', error);
        showToast({
          title: 'Could not Create Event',
          description: 'An error occurred creating event',
          variant: 'error',
        });
      });
      window.location.reload();
  }

  return (
    <Dialog open={props.open} onOpenChange={props.setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create New Event</DialogTitle>
          <DialogDescription>
            set a title and description for your event  for '{props.groupTitle}' along with
             options to choose from.
          </DialogDescription>
        </DialogHeader>

        <PostAuthorCell author={userDetails() as User} date={new Date()} />
        <TextField onChange={setTitle} value={title()}>
        <TextFieldInput 
        type="text"
         id="title"
         minLength={1}
         maxLength={15}
          disabled={formProcessing()}
        placeholder="event title. (max 15 characters)" />
        </TextField>
        <TextField onChange={setEventDescription} value={eventDescription()}>
          <TextFieldTextArea
            placeholder='event description. (max 200 characters)'
            class='resize-none'
            minLength={1}
            maxLength={200}
            disabled={formProcessing()}
          />
        </TextField>

        <Separator />
        <DialogFooter class='!justify-between gap-4'>
          <TextField onChange={setfirstoption} value={firstoption()}>
        <TextFieldInput 
        type="text"
         id="option1"
         minLength={1}
         maxLength={10}
          disabled={formProcessing()}
        placeholder="option 1" />
        </TextField>
        <TextField onChange={setsecondoption} value={secondoption()}>
        <TextFieldInput 
        type="text"
         id="option2"
         minLength={1}
         maxLength={10}
          disabled={formProcessing()}
        placeholder="option 2" />
        </TextField>
        <TextField
            onChange={setEventTime}
            value={eventTime()}
            required
          >
            {/* <TextFieldLabel for='dob'></TextFieldLabel> */}
            <TextFieldInput
              class='block' // without it calendar icon gets ruined
              type='date'
              min={moment().endOf(`day`).format('YYYY-MM-DD')}
              id='dob'
            />
          </TextField>
          <Button
            class='gap-2'
            disabled={!eventDescription() ||
              !firstoption()||
              !secondoption()|| 
              !title()||
              !eventTime() ||
               formProcessing()}
            onClick={makeEvent}
          >
            {formProcessing() && <img src={tailspin} class='h-full' alt='processing' />}
            {formProcessing() ? 'Posting...' : 'Create'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
