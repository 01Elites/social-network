/* @refresh reload */
import { render } from 'solid-js/web';
import './index.css';
import type { JSXElement } from 'solid-js';
import { Router, Route } from '@solidjs/router';
import HomePage from './pages/home';

export default function App(): JSXElement {
  return (
    <Router>
      <Route path='/' component={HomePage} />
      <Route path='/login' component={HomePage} />
    </Router>
  );
}

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

render(
  () => (
    <Router>
      <Route path='/' component={HomePage} />
      <Route path='/login' component={HomePage} />
      <Route path='/profile' />
    </Router>
  ),
  root!,
);
