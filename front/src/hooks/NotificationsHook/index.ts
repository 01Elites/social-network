import { onCleanup, useContext } from 'solid-js';
import { createStore } from 'solid-js/store';
import WebSocketContext from '~/contexts/WebSocketContext';
import { SNNotification } from '~/types/Notification';
import { WebsocketHook } from '../WebsocketHook';
import { useWebsocket } from '../WebsocketHook';

type NotificationsHook = {
  store: SNNotification[];
  markRead: (notificationId: string) => void;
};

function useNotifications(): NotificationsHook {
  const [store, setStore] = createStore([] as SNNotification[]);
  const wsCtx = useWebsocket();

  function markRead(notificationId: string): void {
    wsCtx.send({
      event: 'NOTIFICATION_READ',
      payload: { notification_id: notificationId },
    });
  }

  const nsUnbind = wsCtx.bind('NOTIFICATION', (data) => {
    setStore((prev) => {
      return [...prev, data];
    });
  });

  onCleanup(() => {
    nsUnbind();
  });

  return {
    store,
    markRead,
  };
}

export { useNotifications };
export type { NotificationsHook };
