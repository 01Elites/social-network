import { createSignal, JSXElement, Setter, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { UserDetailsHook } from '~/hooks/userDetails';
import { GroupChatState } from '~/pages/group/groupfeed';
import Message_Icon from '../ui/icons/message_icon';
import GroupChatMessage from './groupchatmessage';
import { Button } from '../ui/button';
import { cn } from '~/lib/utils';
import { EmojiPicker } from 'solid-emoji-picker';
import { FiSmile } from 'solid-icons/fi';

interface FeedProps {
  class?: string;
  chatState?: GroupChatState;
  setChatState?: Setter<GroupChatState>
}

export default function GroupChatPage(props: FeedProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [message, setMessage] = createSignal<string>('');
  const [messages, setMessages] = createSignal<any[]>([]);
  const useWebsocket = useContext(WebSocketContext) as WebsocketHook;

  useWebsocket.send({
    event: "CHAT_OPENED",
    payload: {
      recipient: props.chatState?.chatWith,
      is_group: true,
    }
  });

  useWebsocket.bind('GET_MESSAGES', (data) => {
    setMessages(prevMessages => [...prevMessages, data]);
  });
  function pickEmoji(emoji: { emoji: any }) {
    let input = document.getElementById('message')! as HTMLInputElement;
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
      <div class="overflow-y-scroll grow">
        {messages().length > 0 &&
          messages().map((msg, index) => (
            <GroupChatMessage
              message={msg.messages[0].message}
              sender={msg.messages[0].sender}
              type={msg.messages[0].sender === userDetails()!.user_name ? "sent" : "received"}
            />
          ))
        }
      </div>
      <div id="emoji-picker" class="items-end self-end hidden h-32 w-96 overflow-y-scroll"><EmojiPicker onEmojiClick={pickEmoji} /></div>
      <TextField class='flex flex-row w-full content-end items-end self-end align-bottom'>
        {/* <Button onClick={() => {
          console.log('Close chat');
          props.setChatState!({
            isOpen: false,
            chatWith: '',
          });
          useWebsocket.send({
            event: 'CHAT_CLOSED',
            payload: {},
          });
        }}>Close</Button> */}
        <TextFieldInput
          type='text'
          id='message'
          value={message()}
          placeholder='Type a message'
          onChange={(event: { currentTarget: { value: string } }) => {
            setMessage(event.currentTarget.value);
          }}
          onKeyPress={(e: KeyboardEvent) => {
            if (e.key === 'Enter') {
              e.preventDefault();
              sendMessage();
            }
          }}
        />
        <Button
          title='emoji picker'
          class="emoji-button ml-2" onclick={openEmojiPicker}>
          <FiSmile size="30" />
        </Button>
        <Message_Icon
          darkBack={false}
          class='ml-2 self-center'
          onClick={sendMessage}
        />
      </TextField>
    </div>
  );
}
