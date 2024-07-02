import { JSXElement, createSignal } from 'solid-js';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';

import logo from '~/logo.svg';
import rebootLogo from '~/reboot_01_logo.png';

import { Button } from '~/components/ui/button';
import {
  TextField,
  TextFieldInput,
  TextFieldLabel,
  TextFieldTextArea,
} from '~/components/ui/text-field';
interface LoginDialogProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

const loginMessages = [
  'Waste your time here ✨',
  "Guess it's time to complain about your life 🤔",
];

const signUpMessages = [
  'Join the Bullies Community 🌟',
  'You are just one step away from being a bully 🤩',
];

export default function LoginDialog(props: LoginDialogProps): JSXElement {
  const [showLogin, setShowLogin] = createSignal(true);

  return (
    <Dialog
      open={props.open}
      onOpenChange={(isOpen) => {
        props.setOpen(isOpen);
        setShowLogin(true);
      }}
    >
      <DialogContent>
        <DialogHeader>
          <div class={showLogin() ? 'flex justify-center' : ''}>
            <img src={logo} alt='Elite Logo' class='w-20' />
          </div>
          <DialogTitle
            class={showLogin() ? 'text-center text-3xl' : 'text-3xl text-left'}
          >
            {showLogin() ? 'Oh, no life?' : 'Sign Up'}
          </DialogTitle>
          <DialogDescription class={showLogin() ? 'text-center' : 'text-left'}>
            {showLogin() ? loginMessages.random() : signUpMessages.random()}
          </DialogDescription>
        </DialogHeader>

        {showLogin() && (
          <form class='flex flex-col gap-4'>
            <Button variant='outline' class='gap-4'>
              <img src={rebootLogo} class='h-5'></img>
              Login with Reboot01
            </Button>
            <TextField class='grid w-full items-center gap-1.5'>
              <TextFieldLabel for='email'>Email</TextFieldLabel>
              <TextFieldInput type='email' id='email' placeholder='Email' />
            </TextField>

            <TextField class='grid w-full items-center gap-1.5'>
              <TextFieldLabel for='password'>Password</TextFieldLabel>
              <TextFieldInput
                type='password'
                id='password'
                placeholder='Password'
              />
            </TextField>

            <Button>Login</Button>
            <p class='text-center'>
              Don't have an account?{' '}
              {/* TODO: remove show underline after we add our acccent */}
              <Button
                variant='link'
                class='p-0 text-base underline'
                onClick={() => setShowLogin(false)}
              >
                Sign up for Free
              </Button>
            </p>
          </form>
        )}

        {!showLogin() && (
          <form class='grid grid-cols-2 gap-4 w-full'>
            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='fname'>First Name</TextFieldLabel>
              <TextFieldInput type='text' id='fname' placeholder='Yaman' />
            </TextField>
            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='lname'>Last Name</TextFieldLabel>
              <TextFieldInput type='text' id='lname' placeholder='Almasri' />
            </TextField>
            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='email'>Email</TextFieldLabel>
              <TextFieldInput
                type='email'
                id='email'
                placeholder='yaman@reboot01.com'
              />
            </TextField>
            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='dob'>Date of Birth</TextFieldLabel>
              <TextFieldInput type='date' id='dob' placeholder='30/6/2024' />
            </TextField>

            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='nickname'>Nickname</TextFieldLabel>
              <TextFieldInput
                type='text'
                id='nickname'
                placeholder='yalmasri'
              />
            </TextField>

            {/* TODO: Solid-ui was down at 1 am, tomorrow when its up switch this to a ComboBox */}
            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='privacy'>Profile Privacy</TextFieldLabel>
              <select class='border-input border-[1px] h-full p-2 rounded-md'>
                <option value='public'>Public</option>
                <option value='private'>Private</option>
              </select>
            </TextField>

            <TextField class='grid w-full items-center gap-1.5 col-span-2'>
              <TextFieldLabel for='about'>About you</TextFieldLabel>
              <TextFieldTextArea
                class=' resize-none'
                id='about'
                placeholder='The biggest Looser 🤡'
                autoResize={false}
              />
            </TextField>

            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='password'>Password</TextFieldLabel>
              <TextFieldInput
                type='password'
                id='password'
                placeholder='your password'
              />
            </TextField>

            <TextField class='grid w-full items-center gap-1.5 col-span-1'>
              <TextFieldLabel for='confirm-password'>
                Confirm Password
              </TextFieldLabel>
              <TextFieldInput
                type='password'
                id='confirm-password'
                placeholder='confirm your password'
              />
            </TextField>

            <Button class='col-span-2'>Become a Looser</Button>
            <Button
              variant='link'
              class='p-0 text-base underline justify-start'
              onClick={() => setShowLogin(true)}
            >
              I am already a looser
            </Button>
          </form>
        )}
      </DialogContent>
    </Dialog>
  );
}
