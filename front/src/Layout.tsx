import { JSXElement } from 'solid-js';
import { Toaster } from '~/components/ui/toast';
import Navigation from './components/core/navigation';
import { LoginDialog } from './components/LoginDialog';
import { SettingsPage } from './pages/settings';

type LayoutProps = {
  children: JSXElement;
};

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <>
      <Navigation>{props.children}</Navigation>
      <LoginDialog />
      <SettingsPage />
      <Toaster />
    </>
  );
}
