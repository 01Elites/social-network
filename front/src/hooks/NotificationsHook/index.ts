import { createEffect, onCleanup, useContext } from 'solid-js';
import { createStore } from 'solid-js/store';
import WebSocketContext from '~/contexts/WebSocketContext';
import { SNNotification } from '~/types/Notification';
import { WebsocketHook } from '../WebsocketHook';
import { useWebsocket } from '../WebsocketHook';
import { showToast } from '~/components/ui/toast';

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

  createEffect(() => {
    if (wsCtx.state() === 'connected') {
      wsCtx.send({ event: 'GET_NOTIFICATIONS', payload: null });
    }
  });
  
  function markRead(notificationId: string, remove = false): void {
    // mark notification as read
    setStore((prev) => {
      return prev.map((n) => {
        if (n.notification_id === notificationId) {
          return { ...n, read: true };
        }
        return n;
      });
    });

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
    if (data.read === false) {
    showToast({
      title: "Notification",
      description: data.message,
    })
  }
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
