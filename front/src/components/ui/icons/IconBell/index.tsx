import { IconProps } from '../icon';

export function IconBell(props: IconProps) {
  return (
    <svg
      class={props.class}
      xmlns='http://www.w3.org/2000/svg'
      viewBox='0 0 23.916 25.7715'
    >
      <g>
        <rect height='25.7715' opacity='0' width='23.916' x='0' y='0' />
        <path
          d='M0 20.0293C0 20.8301 0.615234 21.3672 1.62109 21.3672L7.35352 21.3672C7.44141 23.75 9.24805 25.752 11.7773 25.752C14.3066 25.752 16.1133 23.7598 16.2012 21.3672L21.9238 21.3672C22.9395 21.3672 23.5547 20.8301 23.5547 20.0293C23.5547 18.7695 22.3047 17.6465 21.2402 16.5332C20.2832 15.5078 20.1367 13.418 19.9707 11.6309C19.8047 7.08008 18.5254 4.0332 15.2734 2.92969C14.8828 1.25977 13.5645 0 11.7773 0C9.99023 0 8.66211 1.25977 8.28125 2.92969C5.0293 4.0332 3.75 7.08008 3.58398 11.6309C3.41797 13.418 3.27148 15.5078 2.31445 16.5332C1.24023 17.6465 0 18.7695 0 20.0293ZM2.08008 19.7363L2.08008 19.5898C2.25586 19.1309 3.06641 18.3008 3.76953 17.5293C4.82422 16.377 5.10742 14.3652 5.29297 11.7871C5.48828 6.80664 7.00195 5.01953 9.17969 4.42383C9.53125 4.33594 9.72656 4.16992 9.74609 3.83789C9.82422 2.5293 10.5762 1.5918 11.7773 1.5918C12.9785 1.5918 13.7305 2.5293 13.7988 3.83789C13.8281 4.16992 14.0234 4.33594 14.375 4.42383C16.5527 5.01953 18.0664 6.80664 18.2617 11.7871C18.4473 14.3652 18.7305 16.377 19.7852 17.5293C20.4883 18.3008 21.2988 19.1309 21.4746 19.5898L21.4746 19.7363ZM9.0332 21.3672L14.5215 21.3672C14.4238 23.1152 13.3203 24.2383 11.7773 24.2383C10.2344 24.2383 9.13086 23.1152 9.0332 21.3672Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
      </g>
    </svg>
  );
}

export function IconBellActive(props: IconProps) {
  return (
    <svg
      class={props.class}
      xmlns='http://www.w3.org/2000/svg'
      viewBox='0 0 23.916 29.0332'
    >
      <g>
        <rect height='29.0332' opacity='0' width='23.916' x='0' y='0' />
        <path
          d='M13.3065 1.97674C12.9457 2.37617 12.6441 2.82893 12.4154 3.32217C12.2205 3.25697 12.0071 3.22266 11.7773 3.22266C10.5762 3.22266 9.82422 4.16016 9.74609 5.46875C9.72656 5.80078 9.53125 5.9668 9.17969 6.05469C7.00195 6.65039 5.48828 8.4375 5.29297 13.418C5.10742 15.9961 4.82422 18.0078 3.76953 19.1602C3.06641 19.9316 2.25586 20.7617 2.08008 21.2207L2.08008 21.3672L21.4746 21.3672L21.4746 21.2207C21.2988 20.7617 20.4883 19.9316 19.7852 19.1602C18.7305 18.0078 18.4473 15.9961 18.2617 13.418C18.2352 12.7427 18.1845 12.1262 18.1092 11.5681C18.7001 11.5251 19.2668 11.3869 19.7944 11.1685C19.8873 11.8223 19.9437 12.5222 19.9707 13.2617C20.1367 15.0488 20.2832 17.1387 21.2402 18.1641C22.3047 19.2773 23.5547 20.4004 23.5547 21.6602C23.5547 22.4609 22.9395 22.998 21.9238 22.998L16.2012 22.998C16.1133 25.3906 14.3066 27.3828 11.7773 27.3828C9.24805 27.3828 7.44141 25.3809 7.35352 22.998L1.62109 22.998C0.615234 22.998 0 22.4609 0 21.6602C0 20.4004 1.24023 19.2773 2.31445 18.1641C3.27148 17.1387 3.41797 15.0488 3.58398 13.2617C3.75 8.71094 5.0293 5.66406 8.28125 4.56055C8.66211 2.89062 9.99023 1.63086 11.7773 1.63086C12.3379 1.63086 12.8524 1.75482 13.3065 1.97674ZM9.0332 22.998C9.13086 24.7461 10.2344 25.8691 11.7773 25.8691C13.3203 25.8691 14.4238 24.7461 14.5215 22.998Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
        {/* The Dot */}
        <path
          d='M17.6465 10.0391C19.9609 10.0391 21.8848 8.125 21.8848 5.80078C21.8848 3.47656 19.9609 1.5625 17.6465 1.5625C15.3125 1.5625 13.4082 3.47656 13.4082 5.80078C13.4082 8.125 15.3125 10.0391 17.6465 10.0391Z'
          fill='#ff453a'
        />
      </g>
    </svg>
  );
}
