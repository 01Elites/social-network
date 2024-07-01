import { JSXElement, createSignal, useContext } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import logo from '~/logo.svg';
import {
  TextField,
  TextFieldInput,
  TextFieldLabel,
} from '~/components/ui/text-field';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '~/components/ui/dropdown-menu';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import { Sheet, SheetContent, SheetTrigger } from '~/components/ui/sheet';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '~/components/ui/dialog';

type NavbarProps = {};

export default function Navbar(prop: NavbarProps): JSXElement {
  const { userDetails, updateUserDetails, fetchUserDetails } = useContext(
    UserDetailsContext,
  ) as UserDetailsHook;

  const navigate = useNavigate();
  const [loginDialogVisible, setLoginDialogVisible] = createSignal(false);

  return (
    <>
      <Dialog
        open={loginDialogVisible()}
        onOpenChange={(isOpen) => setLoginDialogVisible(isOpen)}
      >
        <DialogContent>
          <DialogHeader>
            <div class='flex justify-center'>
              <img src={logo} alt='Elite Logo' class='w-20' />
            </div>
            <DialogTitle class='text-center text-3xl'>Welcome back</DialogTitle>
            <DialogDescription class='text-center'>
              Glad to see you again 👋 <br />
              Login to your account below
            </DialogDescription>
          </DialogHeader>
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
            <Button variant='link' class='p-0 text-base'>
              Sign up for Free
            </Button>
          </p>
        </DialogContent>
      </Dialog>

      <header
        style={{
          width: 'calc(100% - 40px)',
        }}
        class='fixed mx-5 hidden h-[70px] items-center justify-between gap-4 align-middle xs:flex'
      >
        <img
          src={logo}
          alt='Elite Logo'
          onClick={() => {
            navigate('/');
          }}
          class='w-20 cursor-pointer'
        />
        <TextField class='flex basis-1/2 items-center'>
          <TextFieldInput
            type='search'
            placeholder='Search friends, groups, posts...'
          />
        </TextField>
        {/* If Logged in show details, else show  */}
        {userDetails() ? (
          <DropdownMenu>
            <DropdownMenuTrigger class='flex flex-row items-center gap-2'>
              <Avatar>
                <AvatarImage src='https://thispersondoesnotexist.com/'></AvatarImage>
                <AvatarFallback>N</AvatarFallback>
              </Avatar>
              Natheer
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => {
                  navigate('/profile');
                }}
              >
                Profile
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => {
                  fetch('/api/auth/logout', { method: 'DELETE' }).finally(
                    () => {
                      navigate('/');
                    },
                  );
                }}
              >
                Logout
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button
            variant='secondary'
            onClick={() => {
              setLoginDialogVisible(true);
            }}
          >
            Login
          </Button>
        )}
      </header>
    </>
  );
}
