import { JSXElement } from 'solid-js';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function IconFlag(props?: inputs): JSXElement {
  return (
    <div class={props?.class} onClick={props?.onClick}>
      <svg
        style={{ width: '100%', height: '100%' }}
        viewBox='0 0 24 24'
        fill='none'
        xmlns='http://www.w3.org/2000/svg'
      >
        <g clip-path='url(#clip0_1_2)'>
          <path
            d='M2.25 2.2615L20.8254 11L2.25 19.7385L2.25 2.2615Z'
            stroke='hsl(var(--primary))'
          />
          <line x1='2.5' y1='2' x2='2.5' y2='22' stroke='hsl(var(--primary))' />
        </g>
        <defs>
          <clipPath id='clip0_1_2'>
            <rect
              width='20'
              height='20'
              fill='white'
              transform='translate(2 2)'
            />
          </clipPath>
        </defs>
      </svg>
    </div>
  );
}
