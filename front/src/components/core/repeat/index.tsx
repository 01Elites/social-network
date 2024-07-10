import { JSXElement } from 'solid-js';

interface RepeatProps {
  count: number;
  children: JSXElement;
}

export default function Repeat(props: RepeatProps): JSXElement {
  return (
    <>{Array.from({ length: props.count }).map((_, i) => props.children)}</>
  );
}
