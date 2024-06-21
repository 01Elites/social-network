import { JSXElement } from 'solid-js';
import two_persons from './two_persons.svg'

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Two_Persons_Icon(props?: inputs): JSXElement {

  return (
    <img
      src={two_persons}
      alt='Two Persons'
      class={props ? props.class : ''}
      onClick={props ? props.onClick : () => { }}

    />
  )
}
