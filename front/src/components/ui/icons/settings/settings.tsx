import { JSXElement } from 'solid-js';
import settings_icon from './settings.svg'

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Settings_Icon(props?: inputs): JSXElement {

  return (
    <img
      src={settings_icon}
      alt='Two Persons'
      class={props ? props.class : ''}
      onClick={props ? props.onClick : () => { }}

    />
  )
}
