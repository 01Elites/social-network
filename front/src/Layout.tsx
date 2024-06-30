import { JSXElement } from 'solid-js';
import Navbar from './components/core/navbar';
import SideBar from './components/core/sidebar';

type LayoutProps = {
  children: JSXElement;
  variant?: 'loggedin' | 'loggedout';
};

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <div class='flex h-screen w-full flex-col'>
      <div class='flex min-h-[70px] w-full'>
        <Navbar />
      </div>

      <div class='flex grow'>
        <div class='w-10'>
          <div class='h-full w-10'>
            <SideBar />
          </div>
        </div>

        <main class='h-screen w-full overflow-scroll'>{props.children}</main>
      </div>

      {/* <Navbar />
      
      <SideBar></SideBar> */}
    </div>
  );
}
