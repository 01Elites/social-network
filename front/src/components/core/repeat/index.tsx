import { JSXElement } from 'solid-js';

interface RepeatProps {
  count: number;
  children: JSXElement;
}

/**
 * Renders the children a specified number of times.
 * @param {number} count - The number of times to render the children.
 * @param {JSX.Element} children - The children to render.
 * @returns {JSX.Element} The rendered Repeat component.
 * @example
 * <><Repeat count={5}><p>Hello, World!</p></Repeat></>
 */
export default function Repeat(props: RepeatProps): JSXElement {
  return (
    <>{Array.from({ length: props.count }).map((_, i) => props.children)}</>
  );
}
