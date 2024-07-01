import { JSXElement, useContext } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import logo from '~/logo.svg';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
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

type NavbarProps = {};

export default function Navbar(prop: NavbarProps): JSXElement {
  const navigate = useNavigate();

  const { userDetails, updateUserDetails, fetchUserDetails } = useContext(
    UserDetailsContext,
  ) as UserDetailsHook;

  return (
    <>
      <header
        style={{
          width: 'calc(100% - 40px)',
        }}
        class='fixed mx-5 hidden h-[70px] justify-between gap-4 align-middle xs:flex'
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
                fetch('/api/auth/logout', { method: 'DELETE' }).finally(() => {
                  navigate('/');
                });
              }}
            >
              Logout
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </header>
    </>
  );
}
