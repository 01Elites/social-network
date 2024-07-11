import { JSXElement } from 'solid-js';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function IconApps(props?: inputs): JSXElement {
  return (
    <div
      class='flex h-11 w-11 items-center justify-center rounded bg-accent'
      onClick={props?.onClick}
    >
      <div class={props?.class}>
        <svg
          style={{ width: '100%', height: '100%' }}
          viewBox='0 0 24 24'
          fill='none'
          xmlns='http://www.w3.org/2000/svg'
        >
          <rect
            x='2.5'
            y='2.5'
            width='8'
            height='8'
            stroke='hsl(var(--primary))'
          />
          <rect
            x='2.5'
            y='13.5'
            width='8'
            height='8'
            stroke='hsl(var(--primary))'
          />
          <rect
            x='13.5'
            y='2.5'
            width='8'
            height='8'
            stroke='hsl(var(--primary))'
          />
          <rect
            x='13.5'
            y='13.5'
            width='8'
            height='8'
            stroke='hsl(var(--primary))'
          />
        </svg>
      </div>
    </div>
  );
}
