import { IconProps } from '../icon';

export default function IconEllipsis(props: IconProps) {
  return (
    <svg
      class={props.class}
      viewBox='0 0 23.1641 4.58008'
      xmlns='http://www.w3.org/2000/svg'
    >
      <g>
        <rect height='4.58008' opacity='0' width='23.1641' x='0' y='0' />
        <path
          d='M20.5176 4.56055C21.7773 4.56055 22.8027 3.54492 22.8027 2.28516C22.8027 1.01562 21.7773 0 20.5176 0C19.2578 0 18.2324 1.01562 18.2324 2.28516C18.2324 3.54492 19.2578 4.56055 20.5176 4.56055Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
        <path
          d='M11.3965 4.56055C12.666 4.56055 13.6816 3.54492 13.6816 2.28516C13.6816 1.01562 12.666 0 11.3965 0C10.1367 0 9.12109 1.01562 9.12109 2.28516C9.12109 3.54492 10.1367 4.56055 11.3965 4.56055Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
        <path
          d='M2.28516 4.56055C3.54492 4.56055 4.57031 3.54492 4.57031 2.28516C4.57031 1.01562 3.54492 0 2.28516 0C1.01562 0 0 1.01562 0 2.28516C0 3.54492 1.01562 4.56055 2.28516 4.56055Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
      </g>
    </svg>
  );
}
