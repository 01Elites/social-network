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
    console.log("marking as read", notificationId)
    console.log(remove)
    for (let i=0;i<store.length;i++){
      if (Number(store[i].notification_id) === Number(notificationId)){
        console.log(store[i], "doesIt")
      }
    }
    if (remove) {
      setStore((prev) => {
        console.log("it reaches here", notificationId)
        console.log(prev)
        return prev.filter((n) => Number(n.notification_id) !== Number(notificationId));
      });
    }else {    
    wsCtx.send({
      event: 'NOTIFICATION_READ',
      payload: { notification_id: notificationId },
    });
  }
    setStore((prev) => {
      return prev.map((n) => {
        if (n.notification_id === notificationId) {
          return { ...n, read: true };
        }
        return n;
      });
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
