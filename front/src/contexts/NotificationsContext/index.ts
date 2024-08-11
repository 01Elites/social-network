import { createContext } from 'solid-js';
import { NotificationsHook } from '~/hooks/NotificationsHook';

const NotificationsContext = createContext<NotificationsHook>();

export default NotificationsContext;
