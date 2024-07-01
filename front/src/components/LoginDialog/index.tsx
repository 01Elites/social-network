import { JSXElement } from 'solid-js';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';

import logo from '~/logo.svg';
import rebootLogo from '~/reboot_01_logo.png';

import {
  TextField,
  TextFieldInput,
  TextFieldLabel,
} from '~/components/ui/text-field';

import { Button } from '~/components/ui/button';
interface LoginDialogProps {
  open: boolean;
  setOpen: (open: boolean) => void;
}

const loginMessages = [
  'Waste your time here ✨',
  "Guess it's time to complain aabout your life 🤔",
];
export default function LoginDialog(props: LoginDialogProps): JSXElement {
  return (
    <Dialog open={props.open} onOpenChange={(isOpen) => props.setOpen(isOpen)}>
      <DialogContent>
        <DialogHeader>
          <div class='flex justify-center'>
            <img src={logo} alt='Elite Logo' class='w-20' />
          </div>
          <DialogTitle class='text-center text-3xl'>Oh, no life?</DialogTitle>
          <DialogDescription class='text-center'>
            {loginMessages.random()}
          </DialogDescription>
        </DialogHeader>

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
          <Button variant='link' class='p-0 text-base underline'>
            Sign up for Free
          </Button>
        </p>
      </DialogContent>
    </Dialog>
  );
}
