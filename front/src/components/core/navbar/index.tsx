import { JSXElement } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import logo from '../../../logo.svg';
import { TextField, TextFieldInput } from '../../../components/ui/text-field';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '../../../components/ui/dropdown-menu';
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from '../../..//components/ui/avatar';

type NavbarProps = {
  variant?: 'loggedin' | 'loggedout';
};

export default function Navbar(prop: NavbarProps): JSXElement {
  // Default variant is 'loggedin'
  if (!prop.variant) {
    prop.variant = 'loggedin';
  }

  const navigate = useNavigate();

  return (
    <header class='mx-5 mt-5 flex justify-between gap-4 align-middle'>
      <img
        src={logo}
        alt='Elite Logo'
        onClick={() => {
          navigate('/');
        }}
        class='w-20 cursor-pointer'
      />
      <TextField class='basis-1/2'>
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
  );
}
