import { JSXElement } from 'solid-js';

interface TextBreakerProps {
  text: string;
}
/**
 * Renders a text with line breaks.
 * @param {string} text - The text to be rendered with line breaks.
 * @returns {JSX.Element} The rendered TextBreaker component.
 * @example
 * <p><TextBreaker text='Hello\nWorld' /></p>
 */
export default function TextBreaker(props: TextBreakerProps): JSXElement {
  return (
    <>
      {props.text.split('\n').map((line, index) => {
        return (
          <>
            {line}
            {/* Add <br> except after the last line */}
            {index !== props.text.split('\n').length - 1 && <br />}
          </>
        );
      })}
    </>
  );
}
