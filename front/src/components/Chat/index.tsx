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
    console.log('Received messages:', messages());
    setMessages(prevMessages => [...prevMessages, data]);
  });

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
            if (message.messages[0].sender == userDetails()!.user_name)
              return <ChatMessage message={message.messages[0].message} type="sent" />
            else {
              return <ChatMessage message={message.messages[0].message} type="received" />
            }
          })
        }
      </div>

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
          value={message()}
          type='text'
          id='message'
          placeholder='Type a message'
          onChange={(event: { currentTarget: { value: any } }) => {
            setMessage(event.currentTarget.value);
          }}
          onKeyUp={(event) => {
            if (event.key != "Enter")
              return
            useWebsocket.send({
              event: 'SEND_MESSAGE',
              payload: {
                recipient: props.chatState?.chatWith,
                message: message(),
              },
            });
            setMessage('') // Reset message field
            event.currentTarget.value = message()
          }}
        />
        <Message_Icon
          darkBack={false}
          class='ml-2 self-center'
          onClick={() => {
            // send the message
            useWebsocket.send({
              event: 'SEND_MESSAGE',
              payload: {
                recipient: props.chatState?.chatWith,
                message: message(),
              },
            });
            setMessage('') // Reset message field
          }}
        />
      </TextField>
    </div>
  );
}
