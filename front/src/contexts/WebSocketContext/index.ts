import { createContext } from 'solid-js';
import { WebsocketHook } from '~/hooks/WebsocketHook';

const WebSocketContext = createContext<WebsocketHook>();

export default WebSocketContext;
