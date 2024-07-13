import { JSXElement } from 'solid-js';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Message_Icon(props?: inputs): JSXElement {
  return (
    <div class={props?.class} onClick={props?.onClick}>
      <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="21.8262" height="21.3965">
        <g>
          <rect height="21.3965" opacity="0" width="21.8262" x="0" y="0" />
          <path d="M12.2266 21.3965C12.9297 21.3965 13.4277 20.791 13.7891 19.8535L20.1855 3.14453C20.3613 2.69531 20.459 2.29492 20.459 1.96289C20.459 1.32812 20.0684 0.9375 19.4336 0.9375C19.1016 0.9375 18.7012 1.03516 18.252 1.21094L1.45508 7.64648C0.634766 7.95898 0 8.45703 0 9.16992C0 10.0684 0.683594 10.3711 1.62109 10.6543L6.89453 12.2559C7.51953 12.4512 7.87109 12.4316 8.29102 12.041L19.0039 2.03125C19.1309 1.91406 19.2773 1.93359 19.375 2.02148C19.4727 2.11914 19.4824 2.26562 19.3652 2.39258L9.39453 13.1445C9.01367 13.5449 8.98438 13.877 9.16992 14.5312L10.7227 19.6875C11.0156 20.6738 11.3184 21.3965 12.2266 21.3965Z" fill="hsl(var(--background))" fill-opacity="0.85" />
        </g>
      </svg>
    </div>
  );
}


