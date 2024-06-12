import { JSXElement } from "solid-js";

type LayoutProps = {
    children: JSXElement
}

export default function Layout(props: LayoutProps): JSXElement {
    return (
        <>
        <header class=" bg-slate-950">
            This is header
        </header>
        <main>
            {props.children}
        </main>
        <footer>this is footer</footer>
        </>
    )
}