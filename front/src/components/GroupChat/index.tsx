import { EmojiPicker } from 'solid-emoji-picker';
import { createSignal, JSXElement, Setter, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { UserDetailsHook } from '~/hooks/userDetails';
import { cn } from '~/lib/utils';
import { GroupChatState } from '~/pages/group/groupfeed';
import Message_Icon from '../ui/icons/message_icon';
import GroupChatMessage from './groupchatmessage';
import IconSmile from '../ui/icons/IconSmile';

interface FeedProps {
  class?: string;
  chatState?: GroupChatState;
  setChatState?: Setter<GroupChatState>;
}

export default function GroupChatPage(props: FeedProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [message, setMessage] = createSignal<string>('');
  const [messages, setMessages] = createSignal<any[]>([]);
  const useWebsocket = useContext(WebSocketContext) as WebsocketHook;

  useWebsocket.send({
    event: 'CHAT_OPENED',
    payload: {
      recipient: props.chatState?.chatWith,
      is_group: true,
    },
  });

  useWebsocket.bind('GET_MESSAGES', (data) => {
    setMessages((prevMessages) => [...prevMessages, data]);
  });

  function pickEmoji(emoji: { emoji: any }) {
    setMessage((prev) => prev + emoji.emoji);
  }

  function openEmojiPicker() {
    const emojiPicker = document.getElementById('emoji-picker');
    if (emojiPicker) {
      emojiPicker.classList.toggle('hidden');
    }
  }

  const sendMessage = () => {
    const msg = message().trim();
    if (msg === '') return;
    useWebsocket.send({
      event: 'SEND_MESSAGE',
      payload: {
        recipient: props.chatState?.chatWith,
        message: msg,
      },
    });
    setMessage(''); // Clear the input field after sending the message
  };

  return (
    <div class={cn(props.class, 'flex h-full flex-col p-3')}>
      <div class='grow overflow-y-scroll'>
        {messages().length > 0 &&
          messages().map((msg, index) => (
            <GroupChatMessage
              message={msg.messages[0].message}
              sender={msg.messages[0].sender}
              type={
                msg.messages[0].sender === userDetails()!.user_name
                  ? 'sent'
                  : 'received'
              }
            />
          ))}
      </div>
      <div
        id='emoji-picker'
        class='hidden h-32 w-96 items-end self-end overflow-y-scroll'
      >
        <EmojiPicker onEmojiClick={pickEmoji} />
      </div>
      <TextField class='flex w-full flex-row content-end items-end self-end align-bottom'>
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
              const input = document.getElementById(
                'message',
              ) as HTMLInputElement;
              setMessage(input.value); // Ensure the latest input value is captured
              sendMessage();
            }
          }}
        />
        <IconSmile
          class='ml-2 size-6 cursor-pointer self-center'
          onClick={openEmojiPicker}
        />
        <Message_Icon
          darkBack={false}
          class='ml-2 self-center'
          onClick={sendMessage}
        />
      </TextField>
    </div>
  );
}
