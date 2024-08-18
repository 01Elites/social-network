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

  useWebsocket.send(
    {
      event: "CHAT_OPENED",
      payload: {
        recipient: props.chatState?.chatWith,
        is_group: true,
      }
    });

  useWebsocket.bind('GET_MESSAGES', (data) => {
    // console.log('Received messages:', messages());
    setMessages(prevMessages => [...prevMessages, data]);
  });

  return (
    <div class={cn(props.class, "flex flex-col h-full")}>
      <div class="overflow-y-scroll grow">
        {messages().length > 0 &&
          messages().map((message, index) => {
            if (message.messages[0].sender == userDetails()!.user_name)
              return <GroupChatMessage message={message.messages[0].message} type="sent" />
            else {
              return <GroupChatMessage message={message.messages[0].message} sender={message.messages[0].sender } type="received" />
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
          type='text'
          id='message'
          placeholder='Type a message'
          onChange={(event: { currentTarget: { value: any } }) => {
            setMessage(event.currentTarget.value);
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
          }}
        />
      </TextField>
    </div>
  );
}
