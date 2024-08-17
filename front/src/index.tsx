import {
  ColorModeProvider,
  ColorModeScript,
  createLocalStorageManager,
} from '@kobalte/core';
import { Route, Router } from '@solidjs/router';
import { render } from 'solid-js/web';

import '~/extensions';

import UserDetailsContext from './contexts/UserDetailsContext';
import WebSocketContext from './contexts/WebSocketContext';
import { useWebsocket } from './hooks/WebsocketHook';
import { useUserDetails } from './hooks/userDetails';
import './index.css';
import EventsPage from './pages/events';
import FriendsPage from './pages/friends';
import Group from './pages/group';
import GroupsPage from './pages/groups';
import HomePage from './pages/home';
import Profile from './pages/profile';
import { useNotifications } from './hooks/NotificationsHook';
import NotificationsContext from './contexts/NotificationsContext';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

function App() {
  const storageManager = createLocalStorageManager('vite-ui-theme');
  const userDetailsHook = useUserDetails();
  const websocketHook = useWebsocket();
  const notificationHook = useNotifications({ wsCtx: websocketHook });

  return (
    <Router
      root={(props) => (
        <>
          <ColorModeScript storageType={storageManager.type} />
          <ColorModeProvider>
            <UserDetailsContext.Provider value={userDetailsHook}>
              <WebSocketContext.Provider value={websocketHook}>
                <NotificationsContext.Provider value={notificationHook}>
                  {props.children}
                </NotificationsContext.Provider>
              </WebSocketContext.Provider>

            </UserDetailsContext.Provider>
          </ColorModeProvider>
        </>
      )}
    >
      <Route path='/' component={HomePage} />
      <Route path='/group/:id' component={Group} />
      <Route path='/groups' component={GroupsPage} />
      <Route path='/friends' component={FriendsPage} />
      {/* <Route path='/events' component={EventsPage} /> */}
      <Route path='/profile/:username' component={Profile} />
    </Router>
  );
}

render(() => <App />, root!);
