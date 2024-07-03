import { JSXElement, createEffect, createSignal, useContext } from 'solid-js';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '~/components/ui/dialog';

import logo from '~/assets/logo.svg';
import rebootLogo from '~/assets/reboot_01_logo.png';
import tailspin from '~/assets/svg-loaders/tail-spin.svg';

import { Button } from '~/components/ui/button';
import { Checkbox } from '~/components/ui/checkbox';
import { Label } from '~/components/ui/label';
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
import { showToast } from '~/components/ui/toast';
import config from '~/config';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { fetchWithAuth } from '~/extensions/fetch';
import { UserDetailsHook } from '~/types/User';
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
  const { fetchUserDetails } = useContext(
    UserDetailsContext,
  ) as UserDetailsHook;

  const [showLogin, setShowLogin] = createSignal(true);

  const [formProcessing, setFormProcessing] = createSignal(false);

  // -------- Login Dialog --------
  const [loginEmail, setLoginEmail] = createSignal('');
  const [loginPassword, setLoginPassword] = createSignal('');

  function handleLoginForm(e: SubmitEvent) {
    e.preventDefault();
    setFormProcessing(true);

    fetchWithAuth(config.API_URL + '/auth/signin', {
      method: 'POST',
      body: JSON.stringify({ email: loginEmail(), password: loginPassword() }),
    })
      .then(async (res) => {
        setFormProcessing(false);
        if (res.status === 200) {
          fetchUserDetails();
          props.setOpen(false);
          return;
        }

        const error = await res.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error(
          'An error occurred while logging you in. Please try again.',
        );
      })
      .catch((error: Error) => {
        showToast({
          title: 'An error occurred',
          description: error.message,
          variant: 'error',
        });
      });
  }

  function handleLoginWithReboot() {
    console.error('Login with Reboot01 is not implemented yet');
  }

  // -------- Signup Dialog --------
  const [signupFirstName, setSignupFirstName] = createSignal('');
  const [signupLastName, setSignupLastName] = createSignal('');
  const [signupEmail, setSignupEmail] = createSignal('');
  const [signupDOB, setSignupDOB] = createSignal('');
  const [signupNickname, setSignupNickname] = createSignal('');
  const [signupGender, setSignupGender] = createSignal<'female' | 'male'>();
  const [signupAbout, setSignupAbout] = createSignal('');
  const [signupPassword, setSignupPassword] = createSignal('');
  const [signupPrivate, setSignupPrivate] = createSignal(false);
  const [signupConfirmPassword, setSignupConfirmPassword] = createSignal('');

  const [signupPasswordValidation, setSignupPasswordValidation] = createSignal<
    'valid' | 'invalid'
  >('valid');

  function handleSignupForm(e: SubmitEvent) {
    e.preventDefault();
    setFormProcessing(true);
    fetchWithAuth(config.API_URL + '/auth/signup', {
      method: 'POST',
      body: JSON.stringify({
        first_name: signupFirstName(),
        last_name: signupLastName(),
        email: signupEmail(),
        date_of_birth: new Date(signupDOB()).toISOString(),
        nick_name: signupNickname(),
        profile_privacy: signupPrivate() ? 'private' : 'public',
        about: signupAbout(),
        password: signupPassword(),
        gender: signupGender(), // Hardcoded for testing
      }),
    })
      .then(async (res) => {
        setFormProcessing(false);
        if (res.status === 201) {
          showToast({
            title: 'Account created',
            description: 'Your account has been created successfully',
            variant: 'success',
          });
          props.setOpen(false);
          setShowLogin(true);
          return;
        }

        const error = await res.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error(
          'An error occurred while creating your account. Please try again.',
        );
      })
      .catch((error: Error) => {
        showToast({
          title: 'An error occurred',
          description: error.message,
          variant: 'error',
        });
      });
  }

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
          <form class='flex flex-col gap-4' onSubmit={handleLoginForm}>
            <Button
              variant='outline'
              class='gap-4'
              onClick={handleLoginWithReboot}
              disabled={formProcessing()}
            >
              <img src={rebootLogo} class='h-5'></img>
              Login with Reboot01
            </Button>
            <TextField
              class='grid w-full items-center gap-1.5'
              onChange={setLoginEmail}
              required
            >
              <TextFieldLabel for='email'>Email</TextFieldLabel>
              <TextFieldInput type='email' id='email' placeholder='Email' />
            </TextField>

            <TextField
              class='grid w-full items-center gap-1.5'
              onChange={setLoginPassword}
              required
            >
              <TextFieldLabel for='password'>Password</TextFieldLabel>
              <TextFieldInput
                type='password'
                id='password'
                placeholder='Password'
              />
            </TextField>

            <Button
              disabled={
                loginEmail() === '' ||
                loginPassword() === '' ||
                formProcessing()
              }
              type='submit'
              class='gap-4'
            >
              {formProcessing() && <img src={tailspin} class='h-full' />}
              {formProcessing() ? 'Logging in...' : 'Login'}
            </Button>
            <p class='text-center'>
              Don't have an account?{' '}
              {/* TODO: remove show underline after we add our acccent */}
              <Button
                variant='link'
                class='p-0 text-base underline'
                onClick={() => setShowLogin(false)}
                disabled={formProcessing()}
              >
                Sign up for Free
              </Button>
            </p>
          </form>
        )}

        {/* Sign up form */}
        {!showLogin() && (
          <form
            class='grid grid-cols-2 gap-4 w-full'
            onSubmit={handleSignupForm}
          >
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupFirstName}
              required
            >
              <TextFieldLabel for='fname'>First Name</TextFieldLabel>
              <TextFieldInput type='text' id='fname' placeholder='Yaman' />
            </TextField>
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupLastName}
              required
            >
              <TextFieldLabel for='lname'>Last Name</TextFieldLabel>
              <TextFieldInput type='text' id='lname' placeholder='Almasri' />
            </TextField>
            <TextField
              class='grid w-full items-center gap-1.5 col-span-1'
              onChange={setSignupEmail}
              required
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
              required
            >
              <TextFieldLabel for='dob'>Date of Birth</TextFieldLabel>
              <TextFieldInput
                class='block' // without it calendar icon gets ruined
                type='date'
                id='dob'
                placeholder='30/6/2024'
              />
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
              onChange={setSignupGender}
            >
              <TextFieldLabel>Gender</TextFieldLabel>

              <Select
                class='w-full col-span-1'
                placeholder='Select your Gender'
                itemComponent={(props) => (
                  <SelectItem item={props.item}>
                    {props.item.rawValue}
                  </SelectItem>
                )}
                options={['male', 'female']}
                defaultValue='male'
              >
                <SelectTrigger aria-label='profile privacy' class='w-full'>
                  <SelectValue<string>>
                    {(state) => {
                      setSignupGender(state.selectedOption() as any);
                      return state.selectedOption();
                    }}
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
              required
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
              required
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

            <div class='items-top flex space-x-2 col-span-2'>
              <Checkbox
                id='terms1'
                checked={signupPrivate()}
                onChange={setSignupPrivate}
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
              disabled={formProcessing()}
            >
              {formProcessing() && <img src={tailspin} class='h-full' />}
              Become a Looser
            </Button>
            <Button
              variant='link'
              class='p-0 text-base underline justify-start'
              onClick={() => setShowLogin(true)}
              disabled={formProcessing()}
            >
              I am already a looser
            </Button>
          </form>
        )}
      </DialogContent>
    </Dialog>
  );
}
