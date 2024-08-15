import { onCleanup, useContext } from 'solid-js';
import { createStore } from 'solid-js/store';
import WebSocketContext from '~/contexts/WebSocketContext';
import { SNNotification } from '~/types/Notification';
import { WebsocketHook } from '../WebsocketHook';

type NotificationsHook = {
  store: SNNotification[];
  markRead: (notificationId: string, remove?: boolean) => void;
};

type UseNotificationsProps = {
  wsCtx?: WebsocketHook;
};

function useNotifications(props?: UseNotificationsProps): NotificationsHook {
  const [store, setStore] = createStore<SNNotification[]>([]);

  let wsCtx: WebsocketHook = props?.wsCtx as WebsocketHook;

  if (!wsCtx) {
    wsCtx = useContext(WebSocketContext) as WebsocketHook;
    if (!wsCtx) {
      throw new Error('WebSocketContext cannot be initialized');
    }
  }

  function markRead(notificationId: string, remove = false): void {
    if (remove) {
      setStore((prev) => {
        return prev.filter((n) => n.notification_id !== notificationId);
      });
    }

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
