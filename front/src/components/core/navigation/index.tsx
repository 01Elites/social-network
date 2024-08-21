import { A, useLocation, useNavigate } from '@solidjs/router';
import {
  createEffect,
  createSignal,
  For,
  JSXElement,
  Show,
  useContext,
} from 'solid-js';
import { showLogin } from '~/components/LoginDialog';
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button, buttonVariants } from '~/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '~/components/ui/dropdown-menu';
import { IconBell, IconBellActive } from '~/components/ui/icons/IconBell';
import { IconElites, IconElitesSmall } from '~/components/ui/icons/IconElites';
import IconFlag from '~/components/ui/icons/IconFlag';
import IconGroup from '~/components/ui/icons/IconGroup';
import IconHome from '~/components/ui/icons/IconHome';
import IconLogin from '~/components/ui/icons/IconLogin';
import IconSettings from '~/components/ui/icons/IconSettings';
import IconTwoPerson from '~/components/ui/icons/IconTwoPerson';
import { Separator } from '~/components/ui/separator';
import config from '~/config';
import NotificationsContext from '~/contexts/NotificationsContext';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import WebSocketContext from '~/contexts/WebSocketContext';
import { fetchWithAuth } from '~/extensions/fetch';
import useLoginProviders from '~/hooks/LoginProvidersHook';
import { WebsocketHookPrivate } from '~/hooks/WebsocketHook';
import { cn } from '~/lib/utils';
import { showNotifications } from '~/pages/notifications';
import { showSettings } from '~/pages/settings';

interface NavigationProps {
  children: JSXElement;
}

type NavItem = {
  label: string;
  href: string;
  icon?: JSXElement;
  variant: 'default' | 'ghost';
};

export default function Navigation(props: NavigationProps): JSXElement {
  const navigate = useNavigate();
  const userCtx = useContext(UserDetailsContext);
  const notificationsCtx = useContext(NotificationsContext);
  const wsCtx = useContext(WebSocketContext);

  const loginProviders = useLoginProviders();
  loginProviders.postLogin();

  createEffect(() => {
    if (!userCtx?.userDetails()) {
      showLogin();
    }
  });

  const location = useLocation();
  // eh writing the same line every where? sucks
  function cpFill(path: string) {
    return location.pathname === path ? 'white' : undefined;
  }

  function itemVariant(path: string) {
    return location.pathname === path ? 'default' : 'ghost';
  }

  // Define the navItems array with default item based on currentPath match
  const [navItems] = createSignal<NavItem[]>([
    {
      label: 'Home',
      href: '/',
      icon: <IconHome class='size-4' fill={cpFill('/')} />,
      variant: itemVariant('/'),
    },
    {
      label: 'Friends',
      href: '/friends',
      icon: <IconTwoPerson class='size-4' fill={cpFill('/friends')} />,
      variant: itemVariant('/friends'),
    },
    {
      label: 'Groups',
      href: '/groups',
      icon: <IconGroup class='size-4' fill={cpFill('/groups')} />,
      variant: itemVariant('/groups'),
    },
    {
      label: 'Events',
      href: '/events',
      icon: <IconFlag class='size-4' fill={cpFill('/events')} />,
      variant: itemVariant('/events'),
    },
  ]);

  return (
    <div class='flex h-screen w-screen flex-row gap-4 p-4'>
      <nav class='flex min-w-fit max-w-40 flex-col items-center gap-4 overflow-hidden border-r border-border pr-2 md:w-1/4 md:min-w-32'>
        <div class='mb-4 cursor-pointer'>
          <IconElites class='hidden h-full w-full md:block' />
          <IconElitesSmall class='block h-full w-full md:hidden' />
        </div>

        <For each={navItems()}>
          {(navItem) => (
            <A
              href={navItem.href}
              class={cn(
                buttonVariants({ variant: navItem.variant }),
                'w-fit justify-start gap-2 md:w-full',
                navItem.variant === 'default' &&
                  'dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white',
              )}
            >
              {navItem?.icon}
              <span class='hidden md:block'>{navItem.label}</span>
            </A>
          )}
        </For>
        <Show when={userCtx!.userDetails()}>
          <Button
            variant='ghost'
            class='mt-auto w-fit justify-start gap-2 md:w-full'
            color='red'
            onClick={showNotifications}
          >
            <Show
              when={notificationsCtx?.store.find(
                (notification) => notification.read === false,
              )}
              fallback={<IconBell class='size-5' />}
            >
              <IconBellActive class='size-5' />
            </Show>
            <span class='hidden md:block'>Notifications</span>
          </Button>
        </Show>
        <Button
          variant='ghost'
          class={cn(
            'w-fit justify-start gap-2 md:w-full',
            userCtx!.userDetails() ? '' : 'mt-auto',
          )}
          onClick={showSettings}
        >
          <IconSettings class='size-4' />
          <span class='hidden md:block'>Settings</span>
        </Button>
        <Separator />
        <Show
          when={userCtx!.userDetails()}
          fallback={
            <Button
              variant='secondary'
              class='w-full gap-2'
              onClick={showLogin}
            >
              <IconLogin class='size-4' />
              <span class='hidden md:block'>Login</span>
            </Button>
          }
        >
          <DropdownMenu>
            <DropdownMenuTrigger class='flex w-full flex-row items-center justify-center gap-2 px-2 md:justify-start'>
              <Avatar>
                <AvatarImage
                  src={`${config.API_URL}/image/${userCtx!.userDetails()?.avatar}`}
                ></AvatarImage>
                <AvatarFallback>
                  {userCtx!.userDetails()?.first_name.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <span class='hidden md:block'>
                {userCtx!.userDetails()?.first_name}
              </span>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => {
                  navigate(`/profile/${userCtx!.userDetails()?.user_name}`);
                }}
              >
                Profile
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={() => {
                  fetchWithAuth(config.API_URL + '/auth/logout', {
                    method: 'DELETE',
                  }).finally(() => {
                    localStorage.removeItem('SN_TOKEN');
                    userCtx!.setUserDetails(null);
                    // .connect() closes open connection
                    (wsCtx as WebsocketHookPrivate).connect();
                    if (location.pathname !== '/') {
                      navigate('/');
                    }
                  });
                }}
              >
                Logout
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </Show>
      </nav>
      <main class='w-full overflow-scroll'>{props.children}</main>
    </div>
  );
}
