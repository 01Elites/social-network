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
        <div class='w-10 m-9'>
          <div class='w-10 h-full'>
            <SideBar />
          </div>
        </div>

        <main class='h-[90vh] w-full overflow-scroll justify-center self-center items-center'>{props.children}</main>
      </div>

      {/* <Navbar />
      
      <SideBar></SideBar> */}
    </div>
  );
}
