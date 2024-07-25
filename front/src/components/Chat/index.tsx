import { TextField, TextFieldInput, TextFieldLabel } from "~/components/ui/text-field"
import { JSXElement } from "solid-js";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "../ui/card";
import ChatMessage from "./chatMessage";

interface FeedProps {
  class?: string;
}

export default function ChatPage(props: FeedProps): JSXElement {
  return (
    <div class={props.class}>
      <ChatMessage message="Hello" type="sent" />
      <ChatMessage message="Hiiiii" type="received" />
      <ChatMessage message="How is it gdsg goooooinnnngg!!!fdsfjsdjfOJOfnsdjjfsd,klkojojfsd,f mdsknGDSifhiudshughIS!" type="received" />
      <TextField class="align-bottom items-end content-end self-end mb-4">
        <TextFieldInput type="text" id="message" placeholder="Type a message" />
      </TextField>
    </div>
  );
}
