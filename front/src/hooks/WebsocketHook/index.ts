import { Accessor, createSignal } from 'solid-js';
import config from '~/config';

type WSMessage = {
  event: string;
  payload: any;
};

type WebSocketState = 'disconnected' | 'connected' | 'connecting';

type WebsocketHook = {
  state: Accessor<WebSocketState>;
  error: Accessor<string | null>;
  /**
   *
   * @param event the event to listen to
   * @param callback a function that is ran when the event is received
   * @returns a function to unbind the callback, must be called or can cause a memory leak
   */
  bind: (event: string, callback: (data: any) => void) => () => void;
  send: (event: WSMessage) => void;
};

/**
 * @private This is the private interface of the WebsocketHook
 */
type WebsocketHookPrivate = WebsocketHook & {
  connect: () => void;
};

function useWebsocket(): WebsocketHook {
  const [socketState, setSocketState] =
    createSignal<WebSocketState>('disconnected');
  const [socketError, setSocketError] = createSignal<string | null>(null);

  const listnersMap = new Map<
    { event: string; id: string },
    (event: WSMessage) => void
  >();

  let ws: WebSocket;

  function connect() {
    if (ws) {
      ws.close();
    }

    if (!localStorage.getItem('SN_TOKEN')) {
      setSocketError('No token found');
      return;
    }

    ws = new WebSocket(
      `${config.WS_URL}?token=${localStorage.getItem('SN_TOKEN')}`,
    );

    ws.onopen = () => {
      setSocketState('connected');
    };

    ws.onclose = () => {
      setSocketState('disconnected');
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
      setSocketError(error as any);
    };

    ws.onmessage = async (_message) => {
      const message: WSMessage = await JSON.parse(_message.data);
      let broadcasted = false;
      listnersMap.forEach((callback, key) => {
        if (key.event === message.event) {
          broadcasted = true;
          callback(message.payload);
        }
      });

      if (!broadcasted) {
        console.warn('Unhandled event:', message);
      }
    };
  }

  // Bind a callback to a specific event, and return a function to unbind it
  function bind(event: string, callback: (event: WSMessage) => void) {
    const id = Math.random().toString(36).slice(2);
    listnersMap.set({ event, id }, callback);

    return () => {
      listnersMap.delete({ event, id });
    };
  }

  function send(event: WSMessage) {
    ws.send(JSON.stringify(event));
  }
  connect();
  return {
    state: socketState,
    error: socketError,
    bind: bind,
    send: send,
    connect: connect,
  } as WebsocketHook;
}

export { useWebsocket };
export type { WebsocketHook, WebsocketHookPrivate };
