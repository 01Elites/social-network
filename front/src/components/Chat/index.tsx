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
import { EmojiPicker } from 'solid-emoji-picker';
import { IoArrowBackCircle } from 'solid-icons/io'
import IconSmile from '../ui/icons/IconSmile';

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
  function pickEmoji(emoji: { emoji: any }) {
    var input = document.getElementById('message')! as HTMLInputElement;
    if (input) {
      input.value += emoji.emoji
    }
    setMessage(input?.value);
  }
  function openEmojiPicker() {
    const emojiPicker = document.getElementById('emoji-picker');
    if (emojiPicker) {
      emojiPicker.classList.toggle('hidden');
    }
  }
  const sendMessage = () => {
    if (message().trim() === '') return;
    useWebsocket.send({
      event: 'SEND_MESSAGE',
      payload: {
        recipient: props.chatState?.chatWith,
        message: message(),
      },
    });
    setMessage(''); // Clear the input field after sending the message
  };

  return (
    <div class={cn(props.class, "flex flex-col h-full")}>
      {/* <div class="flex h-12 bg-primary-foreground text-primary-background"> */}
      <div class='flex justify-left self-start'><IoArrowBackCircle size="40" onClick={() => {
        console.log('Close chat');
        props.setChatState!({
          isOpen: false,
          chatWith: '',
        });
        useWebsocket.send({
          event: 'CHAT_CLOSED',
          payload: {},
        });
      }}
        class="hover:cursor-pointer" /></div>
      {/* </div> */}
      <div class='flex flex-col justify-center place-items-center text-center'>
        <a href={`/profile/${props.chatState?.chatWith}`}
          class='text-base font-bold '
        >{props.chatState?.chatWith}
        </a></div>
      <div class="overflow-y-scroll grow">
        {messages().length > 0 &&
          messages().map((message, index) => {
            if (message.messages[0].sender == userDetails()!.user_name) {
              return <ChatMessage message={message.messages[0].message} type="sent" />
            } else {
              return <ChatMessage message={message.messages[0].message} type="received" />
            }
          })
        }
      </div>
      <div id="emoji-picker" class="items-end self-end hidden h-44 w-96 overflow-y-scroll"><EmojiPicker onEmojiClick={pickEmoji} /></div>
      <TextField class='flex flex-row w-full content-end items-end self-end align-bottom'>
        <Button
          class='self-center hover:cursor-pointer mr-2'
          onClick={() => {
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
          class='w-lbw'
          value={message()}
          placeholder='Type a message'
          onChange={(event: { currentTarget: { value: any } }) => {
            setMessage(event.currentTarget.value);
          }}
          onKeyUp={(e: KeyboardEvent) => {
            if (e.key === 'Enter') {
              e.preventDefault();
              const input = document.getElementById('message') as HTMLInputElement;
              setMessage(input.value); // Ensure the latest input value is captured
              sendMessage();
            }
          }}
        />

        {/* <FiSmile size="30" class="self-center emoji-button ml-2 cursor-pointer" onclick={openEmojiPicker} /> */}

        <IconSmile class="size-6 self-center cursor-pointer ml-2" onClick={openEmojiPicker} />
        <Message_Icon
          darkBack={false}
          class='self-center hover:cursor-pointer ml-2'
          onClick={sendMessage}
        />
      </TextField>
    </div>
  );
}

