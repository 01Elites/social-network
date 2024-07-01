import { render } from 'solid-js/web';
import './index.css';
import { Router, Route } from '@solidjs/router';
import HomePage from './pages/home';
import Login from './hooks/login';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

// Example usage of the Login function
// Login('jane.smith@example.com', 'S3cur3P@ss').then((res) => {
//   console.log(res);
// })

render(
  () => (
    <Router>
      {/* <AuthGuard> */}
      <Route path='/' component={HomePage} />
      {/* </AuthGuard> */}
      <Route path='/login' component={HomePage} />
      {/* <AuthGuard> */}
      <Route path='/profile' />
      {/* </AuthGuard> */}
    </Router>
  ),
  root!,
);
