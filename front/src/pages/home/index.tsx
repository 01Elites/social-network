import 'solid-devtools';
import { createSignal, JSXElement, Show } from 'solid-js';
import Feed from '~/components/Feed';
import HomeContacts from '~/components/HomeContacts';
import HomeEvents from '~/components/HomeEvents';
import Layout from '../../Layout';
import ChatPage from '~/components/Chat';

export type ChatState = {
  isOpen: boolean; // Whether the chat window is open
  chatWith: string; // The recipient's username
};

export default function HomePage(): JSXElement {
  const [chatState, setChatState] = createSignal<ChatState>({
    isOpen: false,
    chatWith: ''
  });

  return (
    <Layout>
      <section class='flex h-full gap-4'>
        {/* <HomeEvents class='hidden w-5/12 max-w-60 overflow-hidden md:flex' /> */}
        <Show when={chatState().isOpen} fallback={<Feed class='grow overflow-hidden' />}>
          <ChatPage class='grow place-content-end overflow-hidden' chatState={chatState()} setChatState={setChatState} />
        </Show>
        <HomeContacts class='hidden w-1/3 max-w-52 overflow-hidden md:flex' setChatState={setChatState} />
      </section>
    </Layout>
  );
}
