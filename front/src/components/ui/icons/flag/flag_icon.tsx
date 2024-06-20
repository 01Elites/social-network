import { JSXElement } from 'solid-js';
import flag_icon from './flag.svg'

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Flag_Icon(props?: inputs): JSXElement {

  return (
    <img
      src={flag_icon}
      alt='Apps'
      class={props ? props.class : ''}
      onClick={props ? props.onClick : () => { }}

    />
  )
}
