import { JSXElement } from 'solid-js';
import Navbar from './components/core/navbar';

type LayoutProps = {
  children: JSXElement;
  variant?: 'loggedin' | 'loggedout';
};

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <>
      <Navbar />
      <main>{props.children}</main>
    </>
  );
}
