import { JSXElement } from 'solid-js';
import Navbar from './components/core/navbar';
import SideBar from './components/core/sidebar';

type LayoutProps = {
  children: JSXElement;
  variant?: 'loggedin' | 'loggedout';
};

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <div class='flex h-screen w-full flex-col bg-red-400'>
      <div class='fixed flex h-[40px] w-full bg-green-400'>THIS IS NAV</div>

      <div class='flex grow bg-blue-300'>
        <div class='fixed h-full w-10 bg-neutral-600'>THIS IS SIDEBAR</div>

        <main>{props.children}</main>
      </div>

      {/* <Navbar />
      
      <SideBar></SideBar> */}
    </div>
  );
}
