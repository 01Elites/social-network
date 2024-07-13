import { JSXElement } from 'solid-js';
import Navbar from '~/components/core/navbar';
import SideBar from '~/components/core/sidebar';
import { Toaster } from '~/components/ui/toast';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { useUserDetails } from '~/hooks/userDetails';
import { SettingsPage } from './pages/settings';

type LayoutProps = {
  children: JSXElement;
};

export default function Layout(props: LayoutProps): JSXElement {
  const userDetailsHook = useUserDetails();

  return (
    <UserDetailsContext.Provider value={userDetailsHook}>
      <SettingsPage />
      <div class='flex h-screen w-full flex-col'>
        <div class='flex min-h-[70px] w-full'>
          <Navbar />
        </div>

        <div class='flex grow'>
          <div class='m-9 w-10'>
            <div class='h-full w-10'>
              <SideBar />
            </div>
          </div>

          <main class='h-[90vh] w-full items-center justify-center self-center overflow-scroll'>
            {props.children}
          </main>
        </div>

        {/* <Navbar />
      
      <SideBar></SideBar> */}
      </div>
      <Toaster />
    </UserDetailsContext.Provider>
  );
}
