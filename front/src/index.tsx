import {
  ColorModeProvider,
  ColorModeScript,
  createLocalStorageManager,
} from '@kobalte/core';
import { Route, Router } from '@solidjs/router';
import { render } from 'solid-js/web';

import '~/extensions';

import './index.css';
import HomePage from './pages/home';
import Profile from './pages/profile';
import Group from './pages/group';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

function App() {
  const storageManager = createLocalStorageManager('vite-ui-theme');

  return (
    <Router
      root={(props) => (
        <>
          <ColorModeScript storageType={storageManager.type} />
          <ColorModeProvider>{props.children}</ColorModeProvider>
        </>
      )}
    >
      {/* <AuthGuard> */}
      <Route path='/' component={HomePage} />
      {/* </AuthGuard> */}
       <Route path='/group/:id' component={Group} /> 
      {/* <AuthGuard> */}
      <Route path='/profile/:username' component={Profile} />
      {/* </AuthGuard> */}
    </Router>
  );
}

render(() => <App />, root!);
