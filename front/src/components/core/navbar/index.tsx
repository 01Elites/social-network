import { useNavigate } from '@solidjs/router';
import { JSXElement, createSignal, useContext } from 'solid-js';
import logo from '~/assets/logo.svg';
import LoginDialog from '~/components/LoginDialog';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from '~/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '~/components/ui/dropdown-menu';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

type NavbarProps = {};

export default function Navbar(prop: NavbarProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  const navigate = useNavigate();
  const [loginDialogVisible, setLoginDialogVisible] = createSignal(false);

  return (
    <>
      <LoginDialog
        open={loginDialogVisible()}
        setOpen={setLoginDialogVisible}
      />
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
                <AvatarImage src={userDetails()?.avatar_url}></AvatarImage>
                <AvatarFallback>
                  {userDetails()?.first_name.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              {userDetails()?.first_name}
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
          // This is the else
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
