import 'solid-devtools';
import { createEffect, JSXElement, useContext } from 'solid-js';
import Feed from '~/components/Feed';
import HomeContacts from '~/components/HomeContacts';
import HomeEvents from '~/components/HomeEvents';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  const webSocketCtx = useContext(WebSocketContext) as WebsocketHook;

  createEffect(() => {
    console.log('WebSocket State:', webSocketCtx.state());
  });

  createEffect(() => {
    console.log('WebSocket Error:', webSocketCtx.error());
  });

  return (
    <Layout>
      <section class='flex h-full gap-4'>
        <HomeEvents class='hidden w-5/12 max-w-60 overflow-hidden md:flex' />
        <Feed class='grow overflow-hidden' />
        {/* <ChatPage class='grow flex-col place-content-end' /> */}
        <HomeContacts class='hidden w-1/3 max-w-52 overflow-hidden md:flex' />
      </section>
    </Layout>
  );
}
