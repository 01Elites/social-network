import { JSXElement, createEffect, createSignal } from 'solid-js';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';

import logo from '~/assets/logo.svg';
import rebootLogo from '~/assets/reboot_01_logo.png';

import { Button } from '~/components/ui/button';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '~/components/ui/select';
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

  // -------- Login Dialog --------
  const [loginEmail, setLoginEmail] = createSignal('');
  const [loginPassword, setLoginPassword] = createSignal('');

  async function handleLogin() {}

  // -------- Signup Dialog --------
  const [signupFirstName, setSignupFirstName] = createSignal('');
  const [signupLastName, setSignupLastName] = createSignal('');
  const [signupEmail, setSignupEmail] = createSignal('');
  const [signupDOB, setSignupDOB] = createSignal('');
  const [signupNickname, setSignupNickname] = createSignal('');
  const [signupPrivacy, setSignupPrivacy] = createSignal<
    'public' | 'private'
  >();
  const [signupAbout, setSignupAbout] = createSignal('');
  const [signupPassword, setSignupPassword] = createSignal('');
  const [signupConfirmPassword, setSignupConfirmPassword] = createSignal('');

  const [signupPasswordValidation, setSignupPasswordValidation] = createSignal<
    'valid' | 'invalid'
  >('valid');

  createEffect(() => {
    // if empty, don't show validation
    if (signupConfirmPassword() === '') {
      setSignupPasswordValidation('valid');
    } else if (signupPassword() !== signupConfirmPassword()) {
      setSignupPasswordValidation('invalid');
    } else {
      setSignupPasswordValidation('valid');
    }
  });

  // createEffect(() => {
  //   console.log('firstName', signupFirstName());
  //   console.log('lastName', signupLastName());
  //   console.log('email', signupEmail());
  //   console.log('dob', signupDOB());
  //   console.log('nickname', signupNickname());
  //   console.log('privacy', signupPrivacy());
  //   console.log('about', signupAbout());
  //   console.log('password', signupPassword());
  //   console.log('confirm password', signupConfirmPassword());
  // });

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
            <TextField
              class='grid w-full items-center gap-1.5'
              onChange={setLoginEmail}
            >
              <TextFieldLabel for='email'>Email</TextFieldLabel>
              <TextFieldInput type='email' id='email' placeholder='Email' />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5'
              onChange={setLoginPassword}
            >
              <TextFieldLabel for='password'>Password</TextFieldLabel>
              <TextFieldInput
                type='password'
                id='password'
                placeholder='Password'
              />
            </TextField>

            <Button disabled={loginEmail() === '' || loginPassword() === ''}>
              Login
            </Button>
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
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupFirstName}
            >
              <TextFieldLabel for='fname'>First Name</TextFieldLabel>
              <TextFieldInput type='text' id='fname' placeholder='Yaman' />
            </TextField>
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupLastName}
            >
              <TextFieldLabel for='lname'>Last Name</TextFieldLabel>
              <TextFieldInput type='text' id='lname' placeholder='Almasri' />
            </TextField>
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupEmail}
            >
              <TextFieldLabel for='email'>Email</TextFieldLabel>
              <TextFieldInput
                type='email'
                id='email'
                placeholder='yaman@reboot01.com'
              />
            </TextField>
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupDOB}
            >
              <TextFieldLabel for='dob'>Date of Birth</TextFieldLabel>
              <TextFieldInput type='date' id='dob' placeholder='30/6/2024' />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupNickname}
            >
              <TextFieldLabel for='nickname'>Nickname</TextFieldLabel>
              <TextFieldInput
                type='text'
                id='nickname'
                placeholder='yalmasri'
              />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupPrivacy}
            >
              <TextFieldLabel for='privacy'>Profile Privacy</TextFieldLabel>

              <Select
                class='w-full col-span-1'
                placeholder='Profile Privacy'
                itemComponent={(props) => (
                  <SelectItem item={props.item}>
                    {props.item.rawValue}
                  </SelectItem>
                )}
                options={['public', 'private']}
                defaultValue={'public'}
              >
                <SelectTrigger aria-label='profile privacy' class='w-full'>
                  <SelectValue<string>>
                    {(state) => state.selectedOption()}
                  </SelectValue>
                </SelectTrigger>
                <SelectContent />
              </Select>
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5 col-span-2'
              onChange={setSignupAbout}
            >
              <TextFieldLabel for='about'>About you</TextFieldLabel>
              <TextFieldTextArea
                class=' resize-none'
                id='about'
                placeholder='The biggest Looser 🤡'
                autoResize={false}
              />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupPassword}
            >
              <TextFieldLabel for='password'>Password</TextFieldLabel>
              <TextFieldInput
                type='password'
                id='password'
                placeholder='your password'
              />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupConfirmPassword}
              validationState={signupPasswordValidation()}
            >
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
