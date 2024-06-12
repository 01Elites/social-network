import type { JSXElement } from 'solid-js';
import { Router, Routes, Route } from "solid-app-router";
import Layout from './Layout'

type PageWithLayoutProps = {
  component: () => JSXElement;
};

// Not all routes will have the Full App Layout
function PageWithLayout(props: PageWithLayoutProps): JSXElement {
  return (
    <Layout>
      <props.component />
    </Layout>
  );
};


function HomePage(): JSXElement {
  return (
    <h1 class=' text-red-400'>Test homePage</h1>
  )
}

export default function App(): JSXElement {
  return (
    <Router>
        <Route path="/" component={HomePage} />
    </Router>
  );
};

