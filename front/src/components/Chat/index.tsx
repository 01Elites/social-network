import { createSignal, JSXElement, Setter, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { UserDetailsHook } from '~/hooks/userDetails';
import { ChatState } from '~/pages/home';
import Message_Icon from '../ui/icons/message_icon';
import ChatMessage from './chatMessage';
import { Button } from '../ui/button';
import { cn } from '~/lib/utils';
import { FiSmile } from 'solid-icons/fi'
import {
  EmojiPicker,
  useEmojiComponents,
  useEmojiData,
  Emoji,
  loadEmojiData
} from 'solid-emoji-picker';

interface FeedProps {
  class?: string;
  chatState?: ChatState;
  setChatState?: Setter<ChatState>
}

export default function ChatPage(props: FeedProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [message, setMessage] = createSignal<string>('');
  const [messages, setMessages] = createSignal<any[]>([]);
  const useWebsocket = useContext(WebSocketContext) as WebsocketHook;

  useWebsocket.send(
    {
      event: "CHAT_OPENED",
      payload: {
        recipient: props.chatState?.chatWith,
        is_group: false,
      }
    });

  useWebsocket.bind('GET_MESSAGES', (data) => {
    setMessages(prevMessages => [...prevMessages, data]);
  });
  function pickEmoji(emoji) {
    var input = document.getElementById('message')
    if (input) {
      input.value += emoji.emoji
    }
    setMessage(emoji.emoji);
  }
  function openEmojiPicker() {
    const emojiPicker = document.getElementById('emoji-picker');
    if (emojiPicker) {
      emojiPicker.classList.toggle('hidden');
    }
  }

  return (
    <div class={cn(props.class, "flex flex-col h-full")}>
      <div class="flex justify-center items-center h-12 bg-primary-foreground text-primary-background">
        <a href={`/profile/${props.chatState?.chatWith}`}
          class='items-center text-base font-bold '
        >{props.chatState?.chatWith}
        </a>
      </div>
      <div class="overflow-y-scroll grow">
        {messages().length > 0 &&
          messages().map((message, index) => {
            if (message.messages[0].sender == userDetails()!.user_name){
              return <ChatMessage message={message.messages[0].message} type="sent" />
             } else {
              return <ChatMessage message={message.messages[0].message} type="received" />
            }
          })
        }
      </div>
        <div id="emoji-picker" class="items-end self-end hidden h-32 w-96 overflow-y-scroll"><EmojiPicker onEmojiClick={pickEmoji}/></div>
      <TextField class='flex flex-row w-full content-end items-end self-end align-bottom'>
        <Button onClick={() => {
          console.log('Close chat');
          props.setChatState!({
            isOpen: false,
            chatWith: '',
          });
          useWebsocket.send({
            event: 'CHAT_CLOSED',
            payload: {},
          });
        }}>Close</Button>
        <TextFieldInput
          type='text'
          id='message'
          placeholder='Type a message'
          onChange={(event: { currentTarget: { value: any } }) => {
            setMessage(event.currentTarget.value);
          }}
        />
            <button class="emoji-button size-12" onclick={openEmojiPicker}>
            <FiSmile size="34"/>
            </button>
        <Message_Icon
          darkBack={false}
          class='self-center hover:cursor-pointer'
          onClick={() => {
            // send the message
            useWebsocket.send({
              event: 'SEND_MESSAGE',
              payload: {
                recipient: props.chatState?.chatWith,
                message: message(),
              },
            });
            document.getElementById('message')!.value = '';
          }}
        />
      </TextField>
    </div>
  );
}

    