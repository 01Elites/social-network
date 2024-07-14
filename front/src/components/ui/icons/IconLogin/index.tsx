import { JSXElement } from 'solid-js';
import { IconProps } from '../icon';

export default function IconLogin(props: IconProps): JSXElement {
  return (
    <svg
      class={props.class}
      viewBox='0 0 31.2598 27.998'
      xmlns='http://www.w3.org/2000/svg'
    >
      <g>
        <rect height='27.998' opacity='0' width='31.2598' x='0' y='0' />
        <path
          d='M25.8105 3.32031L25.8105 24.6777C25.8105 26.6992 24.4629 27.998 22.373 27.998L8.52539 27.998C6.43555 27.998 5.08789 26.6992 5.08789 24.6777L5.08789 17.0508L6.82617 17.0508L6.82617 24.4043C6.82617 25.5859 7.51953 26.2598 8.74023 26.2598L22.168 26.2598C23.3887 26.2598 24.082 25.5859 24.082 24.4043L24.082 3.59375C24.082 2.40234 23.3887 1.72852 22.168 1.72852L8.74023 1.72852C7.51953 1.72852 6.82617 2.40234 6.82617 3.59375L6.82617 10.9375L5.08789 10.9375L5.08789 3.32031C5.08789 1.29883 6.43555 0 8.52539 0L22.373 0C24.4629 0 25.8105 1.29883 25.8105 3.32031Z'
          fill='white'
          fill-opacity='0.85'
        />
        <path
          d='M0 13.9844C0 14.4531 0.380859 14.8438 0.839844 14.8438L11.123 14.8438L12.9785 14.7754L12.168 15.5566L9.94141 17.6758C9.76562 17.832 9.67773 18.0664 9.67773 18.2715C9.67773 18.7207 10 19.0527 10.4395 19.0527C10.6738 19.0527 10.8398 18.9551 11.0059 18.7891L15.0781 14.6094C15.2832 14.3945 15.3613 14.209 15.3613 13.9844C15.3613 13.7695 15.2832 13.584 15.0781 13.3691L11.0059 9.18945C10.8398 9.02344 10.6738 8.93555 10.4395 8.93555C10 8.93555 9.67773 9.27734 9.67773 9.7168C9.67773 9.92188 9.76562 10.1465 9.94141 10.3027L12.168 12.4316L12.9688 13.1934L11.123 13.1348L0.839844 13.1348C0.380859 13.1348 0 13.5254 0 13.9844Z'
          fill={props.fill ?? 'hsl(var(--primary))'}
          fill-opacity='0.85'
        />
      </g>
    </svg>
  );
}
