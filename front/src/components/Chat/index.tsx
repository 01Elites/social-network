import { createSignal, JSXElement, useContext } from 'solid-js';
import { TextField, TextFieldInput } from '~/components/ui/text-field';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import WebSocketContext from '~/contexts/WebSocketContext';
import { WebsocketHook } from '~/hooks/WebsocketHook';
import { UserDetailsHook } from '~/hooks/userDetails';
import { ChatState } from '~/pages/home';
import Message_Icon from '../ui/icons/message_icon';
import ChatMessage from './chatMessage';

interface FeedProps {
  class?: string;
  chatState?: ChatState;
}

export default function ChatPage(props: FeedProps): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;
  const [message, setMessage] = createSignal<string>('');
  const useWebsocket = useContext(WebSocketContext) as WebsocketHook;

  return (
    <div class={props.class}>
      <ChatMessage message='Hello' type='sent' />
      <ChatMessage message='Hiiiii' type='received' />
      <ChatMessage
        message='How is it gdsg goooooinnnngg!!!fdsfjsdjfOJOfnsdjjfsd,klkojojfsd,f mdsknGDSifhiudshughIS!'
        type='received'
      />
      <TextField class='mb-4 flex flex-row content-end items-end self-end align-bottom'>
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
            useWebsocket.send({
              event: 'CHAT_OPENED',
              payload: {
                recipient: 'las0947',
                is_group: false,
              },
            });
            useWebsocket.send({
              event: 'SEND_MESSAGE',
              payload: {
                recipient: 'las0947',
                message: 'This is a test message',
              },
            });
          }}
        />
      </TextField>
    </div>
  );
}
