import { JSXElement } from 'solid-js';
import { Toaster } from '~/components/ui/toast';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { useUserDetails } from '~/hooks/userDetails';
import Navigation from './components/core/navigation';
import { LoginDialog } from './components/LoginDialog';
import WebSocketContext from './contexts/WebSocketContext';
import { useWebsocket } from './hooks/WebsocketHook';
import { NotificationsPage } from './pages/notifications';
import { SettingsPage } from './pages/settings';
import useLoginProviders from './hooks/LoginProvidersHook';

type LayoutProps = {
  children: JSXElement;
};

export default function Layout(props: LayoutProps): JSXElement {
  const loginProviders = useLoginProviders();

  loginProviders.postLogin();
  return (
    <>
      <Navigation>{props.children}</Navigation>
      <LoginDialog />
      <NotificationsPage />
      <SettingsPage />
      <Toaster /></>
  );
}
