import { JSXElement } from "solid-js";

import logo from "../../../logo.svg";
import { A } from "@solidjs/router";

type NavbarProps = {
  variant?: 'loggedin' | 'loggedout'
}

export default function Navbar(prop: NavbarProps): JSXElement {
  // Default variant is 'loggedin'
  if (!prop.variant) {
    prop.variant = 'loggedin'
  }
  return (
    <header class="flex justify-between">
      <img src={logo} alt="" />
      <A href="/">Login</A>
    </header>
  )
  
}