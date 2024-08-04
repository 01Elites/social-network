import { Accessor, createSignal } from 'solid-js';
import config from '~/config';

type WSEvent = {
  type: string;
  payload: any;
};

type WebSocketState = 'disconnected' | 'connected' | 'connecting';

type WebsocketHook = {
  state: Accessor<WebSocketState>;
  error: Accessor<string | null>;
  bind: (event: string, callback: (event: WSEvent) => void) => () => void;
};

function useWebsocket(): WebsocketHook {
  const [socketState, setSocketState] =
    createSignal<WebSocketState>('disconnected');
  const [socketError, setSocketError] = createSignal<string | null>(null);

  const ws = new WebSocket(
    `${config.WS_URL}?token=${localStorage.getItem('SN_TOKEN')}`,
  );
  const listnersMap = new Map<
    { event: string; id: string },
    (event: WSEvent) => void
  >();

  ws.onopen = () => {
    console.log('Connected to the WebSocket server');
    setSocketState('connected');
  };

  ws.onclose = () => {
    console.log('Disconnected from the WebSocket server');
    setSocketState('disconnected');
  };

  ws.onerror = (error) => {
    console.error('WebSocket error:', error);
    setSocketState('disconnected');
    setSocketError(error as any);
  };

  ws.onmessage = (message) => {
    console.log('Received message:', message);

    const event = JSON.parse(message.data);

    listnersMap.forEach((callback, key) => {
      if (key.event === event.type) {
        callback(event);
      }
    });
  };

  // Bind a callback to a specific event, and return a function to unbind it
  function bind(event: string, callback: (event: WSEvent) => void) {
    const id = Math.random().toString(36).slice(2);
    listnersMap.set({ event, id }, callback);

    return () => {
      listnersMap.delete({ event, id });
    };
  }

  return {
    state: socketState,
    error: socketError,
    bind: bind,
  };
}

export { useWebsocket };
export type { WebsocketHook };
