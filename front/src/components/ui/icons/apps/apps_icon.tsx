import { JSXElement } from 'solid-js';
import apps_icon from './apps.svg';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Apps_Icon(props?: inputs): JSXElement {
  return (
    <div class='flex h-11 w-11 items-center justify-center rounded bg-accent'>
      <img
        src={apps_icon}
        alt='Apps'
        class={props ? props.class : ''}
        onClick={props ? props.onClick : () => {}}
      />
    </div>
  );
}
