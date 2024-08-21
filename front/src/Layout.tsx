import { createEffect, JSXElement, useContext } from 'solid-js';
import { Toaster } from '~/components/ui/toast';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import Navigation from './components/core/navigation';
import { LoginDialog, showLogin } from './components/LoginDialog';
import useLoginProviders from './hooks/LoginProvidersHook';
import { NotificationsPage } from './pages/notifications';
import { SettingsPage } from './pages/settings';

type LayoutProps = {
  children: JSXElement;
};

export default function Layout(props: LayoutProps): JSXElement {
  const userCtx = useContext(UserDetailsContext);

  const loginProviders = useLoginProviders();
  loginProviders.postLogin();

  createEffect(() => {
    if (!userCtx?.userDetails()) {
      showLogin();
    }
  });

  return (
    <>
      <Navigation>{props.children}</Navigation>
      <LoginDialog />
      <NotificationsPage />
      <SettingsPage />
      <Toaster />
    </>
  );
}
