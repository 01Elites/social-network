/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';
import { Route, Router } from 'solid-app-router';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

function Users(props: any) {
  return (
    <div>Please
      <h1>Fukng womt</h1>
    </div>
  )
}

function Home(props: any) {
  return (
    <div>Please Home
      <h1>Fukng w omt</h1>
    </div>
  )
}

render(() => (
  <Router>
    <Route path="/users" component={Users} />
    <Route path="/" component={Home} />
  </Router>
), root!);