import { JSXElement } from 'solid-js';
import apps_icon from './apps.svg'

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Apps_Icon(props?: inputs): JSXElement {

  return (
    <div class='rounded bg-accent h-11 w-11 flex items-center justify-center'>
      <img
        src={apps_icon}
        alt='Apps'
        class={props ? props.class : ''}
        onClick={props ? props.onClick : () => { }}
      />
    </div>
  )
}
