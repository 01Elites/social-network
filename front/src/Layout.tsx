import { JSXElement } from 'solid-js';
import Navbar from './components/core/navbar';
import SideBar from './components/core/sidebar';

type LayoutProps = {
  children: JSXElement;
  variant?: 'loggedin' | 'loggedout';
};

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <>
      <Navbar />
      <SideBar>
        <main>{props.children}</main>
      </SideBar>
    </>
  );
}
