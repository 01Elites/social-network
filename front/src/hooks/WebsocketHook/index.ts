import { createSignal } from 'solid-js';
import config from '~/config';

type WSEvent = {
  type: string;
  payload: any;
};

type WebSocketState = 'disconnected' | 'connected' | 'connecting';

type WebsocketHook = {
  state: WebSocketState;
  error: string | null;
  bind: (event: string, callback: (event: WSEvent) => void) => () => void;
  send: (event: WSEvent) => void;
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
    const event = JSON.parse(message.data);
    let broadcasted = false;

    listnersMap.forEach((callback, key) => {
      if (key.event === event.type) {
        broadcasted = true;
        callback(event);
      }
    });

    if (!broadcasted) {
      console.warn('Unhandled event:', event);
    }
  };

  // Bind a callback to a specific event, and return a function to unbind it
  function bind(event: string, callback: (event: WSEvent) => void) {
    const id = Math.random().toString(36).slice(2);
    listnersMap.set({ event, id }, callback);

    return () => {
      listnersMap.delete({ event, id });
    };
  }

  function send(event: WSEvent) {
    ws.send(JSON.stringify(event));
  }

  return {
    state: socketState(),
    error: socketError(),
    bind: bind,
    send: send,
  };
}

export { useWebsocket };
export type { WebsocketHook };
