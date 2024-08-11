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

type LayoutProps = {
  children: JSXElement;
};

export default function Layout(props: LayoutProps): JSXElement {
  const userDetailsHook = useUserDetails();
  const websocketHook = useWebsocket();
  // const notificationHook = useNotifications();

  return (
    <UserDetailsContext.Provider value={userDetailsHook}>
      <WebSocketContext.Provider value={websocketHook}>
        <Navigation>{props.children}</Navigation>
        <LoginDialog />
        <NotificationsPage />
        <SettingsPage />
        <Toaster />
      </WebSocketContext.Provider>
    </UserDetailsContext.Provider>
  );
}
