import { JSX, Match, Show, Switch } from "solid-js";

interface Props {
  message: string;
  type: "sent" | "received";
};

export default function ChatMessage(props: Props): JSX.Element {
  return (
    <>
      <Switch>
        <Match when={props.type == "sent"}>
          <SentMessage message={props.message} type="sent" />
        </Match>
        <Match when={props.type == "received"}>
          <RecivedMessage message={props.message} type="received" />
        </Match>
      </Switch>
    </>
  );
}

function RecivedMessage(props: Props): JSX.Element {
  return (
    <div class="flex flex-row w-full justify-start">
      <div class="flex flex-row w-fit h-fit p-2 m-2 rounded-xl bg-[hsl(var(--secondary))] text-[hsl(var(--secondary-foreground))] flex-end">
        <p class="break-words max-w-[40vw]">{props.message}</p>
      </div>
    </div >
  );
}

function SentMessage(props: Props): JSX.Element {
  return (
    <div class="flex flex-row w-full justify-end">
      <div class="flex flex-row w-fit h-fit p-2 m-2 rounded-xl bg-[hsl(var(--primary))] text-[hsl(var(--primary-foreground))]">
        <p class="break-words max-w-[40vw]">{props.message}</p>
      </div>
    </div>
  );
}
