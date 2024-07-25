import { TextField, TextFieldInput, TextFieldLabel } from "~/components/ui/text-field"
import { JSXElement } from "solid-js";

interface FeedProps {
  class?: string;
}

export default function ChatPage(props: FeedProps): JSXElement {
  return (
    <div class={props.class}>
      <TextField class="align-bottom items-end content-end self-end mb-4">
        <TextFieldInput type="text" id="message" placeholder="Type a message" />
      </TextField>
    </div>
  );
}
