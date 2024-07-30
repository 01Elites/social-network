import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement, Show } from 'solid-js';
import { createSignal } from 'solid-js'

type InviteParams = {
  username: string | undefined;
  group_id: string;
}
export default function InviteToGroup(props: InviteParams): JSXElement {


  return (<>
      
  </>)
}