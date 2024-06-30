import { JSXElement } from 'solid-js';
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

type NavbarProps = {
  variant?: 'loggedin' | 'loggedout';
};

export default function Navbar(prop: NavbarProps): JSXElement {
  // Default variant is 'loggedin'
  if (!prop.variant) {
    prop.variant = 'loggedin';
  }

  const navigate = useNavigate();

  // const { userDetails, fetchUserDetails, error, updateUserDetails } =
  // useUserDetails();

  // updateUserDetails({ email: 'Natheer' });

  return (
    <>
      <div class='xs:hidden'>
        {/* Mobile Sidemenu, if not wrapped in div will cause hidden padding */}
        <Sheet>
          <SheetTrigger>
            <Button variant='outline' class='mx-5 mt-5 xs:hidden'>
              E
            </Button>
          </SheetTrigger>
          <SheetContent position='left' class='flex flex-col'>
            <img
              src={logo}
              alt='Elite Logo'
              onClick={() => {
                navigate('/');
              }}
              class='w-20 cursor-pointer'
            />
            <TextField>
              <TextFieldInput
                type='search'
                placeholder='Search friends, groups, posts...'
              />
            </TextField>
            <DropdownMenu>
              <DropdownMenuTrigger class='mt-auto flex flex-row items-center gap-2'>
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
          </SheetContent>
        </Sheet>
      </div>

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
