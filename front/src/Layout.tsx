import { JSXElement } from "solid-js";

type LayoutProps = {
  children: JSXElement
}

export default function Layout(props: LayoutProps): JSXElement {
  return (
    <>
      <header class="pt-5 pb-5">
        This is header
      </header>
      <main>
        {props.children}
      </main>
      <footer>this is footer</footer>
    </>
  )
}
