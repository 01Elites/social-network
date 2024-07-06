import { Route, Router } from '@solidjs/router';
import { render } from 'solid-js/web';
import '~/extensions/arrays';
import './index.css';
import HomePage from './pages/home';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

render(
  () => (
    <Router>
      {/* <AuthGuard> */}
      <Route path='/' component={HomePage} />
      {/* </AuthGuard> */}
      {/* <Route path='/login' component={HomePage} /> */}
      {/* <AuthGuard> */}
      {/* <Route path='/profile' /> */}
      {/* </AuthGuard> */}
    </Router>
  ),
  root!,
);
