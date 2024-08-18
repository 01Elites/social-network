import moment from 'moment';
import { JSXElement, Show, createSignal, useContext } from 'solid-js';
import logo from '~/assets/logo.svg';
import tailspin from '~/assets/svg-loaders/tail-spin.svg';
import { Button } from '~/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';
import {
  TextField,
  TextFieldInput,
  TextFieldLabel,
  TextFieldTextArea,
} from '~/components/ui/text-field';
import { showToast } from '~/components/ui/toast';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/hooks/userDetails';
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import { Checkbox } from '../ui/checkbox';
import { Label } from '../ui/label';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../ui/select';

const [editOpen, setEditOpen] = createSignal(false);

function showEditProfile() {
  setEditOpen(true);
}

function ProfileEditDialog(): JSXElement {
  const { userDetails, fetchUserDetails } = useContext(
    UserDetailsContext,
  ) as UserDetailsHook;

  const [formProcessing, setFormProcessing] = createSignal(false);

  const [userUploadedImage, setUploadedImage] = createSignal<File | null>(null);
  const [userNickName, setNickName] = createSignal(userDetails()?.nick_name);
  const [userFirstName, setFirstName] = createSignal(userDetails()?.first_name);
  const [userLastName, setLastName] = createSignal(userDetails()?.last_name);
  const [userDOB, setDOB] = createSignal(
    userDetails()?.date_of_birth
      ? moment(userDetails()?.date_of_birth).format('YYYY-MM-DD')
      : '',
  );
  const [userPrivate, setPrivate] = createSignal(
    userDetails()?.profile_privacy === 'private',
  );
  const [userGender, setGender] = createSignal<'female' | 'male' | 'undefined'>(
    userDetails()?.gender || 'undefined',
  );
  const [userAbout, setAbout] = createSignal(userDetails()?.about);

  function handleImageUpload(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      setUploadedImage(target.files[0]);
    }
  }

  async function handleEditProfileForm(e?: SubmitEvent) {
    e?.preventDefault();
    setFormProcessing(true);

    const payload = {
      nick_name: userNickName(),
      first_name: userFirstName(),
      last_name: userLastName(),
      date_of_birth: new Date(userDOB()).toISOString(),
      profile_privacy: userPrivate() ? 'private' : 'public',
      gender: userGender(),
      about: userAbout(),
      avatar: '',
    };

    if (userUploadedImage()) {
      try {
        const base64 = await userUploadedImage()?.toBase64();
        payload.avatar = base64 as string;
      } catch (error) {
        console.error('Error converting avatar to base64:', error);
      }
    }

    fetchWithAuth(config.API_URL + '/profile', {
      method: 'PATCH',
      body: JSON.stringify(payload),
    })
      .then(async (res) => {
        setFormProcessing(false);
        if (res.status === 200) {
          showToast({
            title: 'Profile Updated',
            description: 'Your Profile has been updated successfully',
            variant: 'success',
          });
          console.log(userDetails());
          setEditOpen(false);
          fetchUserDetails();
          console.log(userDetails());
          return;
        }

        const error = await res.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error(
          'An error occurred while editing your profile. Please try again.',
        );
      })
      .catch((error: Error) => {
        showToast({
          title: 'An error occurred',
          description: error.message,
          variant: 'error',
        });
      })
      .finally(() => {
        setFormProcessing(false);
      });
  }

  return (
    <Dialog
      open={editOpen()}
      onOpenChange={(isOpen) => {
        setEditOpen(isOpen);
      }}
    >
      <DialogContent>
        <DialogHeader>
          <div class={'flex justify-center xs:justify-start'}>
            <img src={logo} alt='Elite Logo' class='w-28' />
          </div>
          <DialogTitle class={'pt-2 text-center text-2xl xs:text-left'}>
            {'Edit profile'}
          </DialogTitle>
        </DialogHeader>
        <form
          class='grid w-full grid-cols-1 gap-4 xs:grid-cols-2'
          onSubmit={handleEditProfileForm}
        >
          <div class='col-span-2 flex justify-center'>
            <input
              placeholder='Upload Image'
              class='hidden'
              type='file'
              id='signupImageUpload'
              accept='image/*'
              onChange={handleImageUpload}
            />
            <Avatar class='size-20'>
              <Show when={userUploadedImage()}>
                <AvatarImage src={URL.createObjectURL(userUploadedImage()!)} />
              </Show>
              <AvatarFallback class='text-xl'>
                {userFirstName()?.[0]}
                {userLastName()?.[0]}
              </AvatarFallback>
              <button
                type='button'
                onClick={() => {
                  if (userUploadedImage()) {
                    URL.revokeObjectURL(userUploadedImage()!.name);
                    setUploadedImage(null);
                  }
                  document.getElementById('signupImageUpload')?.click();
                }}
                class='absolute bottom-0 w-full bg-primary/80 text-center text-primary-foreground'
              >
                {userUploadedImage() ? 'change' : 'set'}
              </button>
            </Avatar>
          </div>
          <TextField
            class='col-span-2 grid w-full items-center gap-1.5'
            onChange={setNickName}
            value={userNickName()}
          >
            <TextFieldLabel for='nickname'>NickName</TextFieldLabel>
            <TextFieldInput
              type='text'
              id='nickname'
              placeholder='Your Nickname'
            />
          </TextField>
          <TextField
            class='col-span-2 grid w-full items-center gap-1.5 xs:col-span-1'
            onChange={setFirstName}
            value={userFirstName()}
            required
          >
            <TextFieldLabel for='fname'>First Name</TextFieldLabel>
            <TextFieldInput type='text' id='fname' placeholder='Yaman' />
          </TextField>
          <TextField
            class='col-span-2 grid w-full items-center gap-1.5 xs:col-span-1'
            onChange={setLastName}
            value={userLastName()}
            required
          >
            <TextFieldLabel for='lname'>Last Name</TextFieldLabel>
            <TextFieldInput type='text' id='lname' placeholder='Almasri' />
          </TextField>
          <TextField
            class='col-span-2 grid w-full items-center gap-1.5 xs:col-span-1'
            onChange={setDOB}
            value={userDOB()}
            required
          >
            <TextFieldLabel for='dob'>Date of Birth</TextFieldLabel>
            <TextFieldInput
              class='block' // without it calendar icon gets ruined
              type='date'
              max={moment().subtract(18, 'years').format('YYYY-MM-DD')}
              id='dob'
            />
          </TextField>

          <TextField
            class='col-span-2 grid w-full items-center gap-1.5 xs:col-span-1'
            onChange={setGender}
          >
            <TextFieldLabel>Gender</TextFieldLabel>

            <Select
              class='col-span-2 w-full xs:col-span-1'
              placeholder='Select your Gender'
              itemComponent={(props) => (
                <SelectItem item={props.item}>{props.item.rawValue}</SelectItem>
              )}
              options={['male', 'female']}
              defaultValue={userGender()}
            >
              <SelectTrigger aria-label='profile privacy' class='w-full'>
                <SelectValue<string>>
                  {(state) => {
                    setGender(state.selectedOption() as any);
                    return state.selectedOption();
                  }}
                </SelectValue>
              </SelectTrigger>
              <SelectContent />
            </Select>
          </TextField>
          <TextField
            class='col-span-2 grid w-full items-center gap-1.5'
            onChange={setAbout}
            value={userAbout()}
          >
            <TextFieldLabel for='about'>About</TextFieldLabel>
            <TextFieldTextArea id='about' placeholder='about me' rows={2} />
          </TextField>

          <div class='items-top col-span-2 flex space-x-2'>
            <Checkbox
              id='terms1'
              checked={userPrivate()}
              onChange={setPrivate}
            />
            <div class='grid gap-1.5 leading-none'>
              <Label for='terms1-input'>Make my profile private</Label>
              <p class='text-sm text-muted-foreground'>
                I am a big looser and I don't want anyone to know about me
              </p>
            </div>
          </div>

          <Button
            type='submit'
            class='col-span-2 gap-4'
            disabled={
              !userFirstName() ||
              !userLastName() ||
              !userDOB() ||
              !userGender() ||
              formProcessing()
            }
          >
            {formProcessing() && <img alt='' src={tailspin} class='h-full' />}
            {'Save Changes'}
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export { ProfileEditDialog, showEditProfile };