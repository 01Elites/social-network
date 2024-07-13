import { cn } from '~/lib/utils';
import { IconProps } from '../icon';

export function IconElites(props: IconProps) {
  return (
    <svg
      class={cn(props.class)}
      width='134'
      height='28'
      viewBox='0 0 134 28'
      fill='none'
      xmlns='http://www.w3.org/2000/svg'
    >
      <path d='M0 3H28' stroke='#06B6D4' stroke-width='6' />
      <path d='M106 3H134' stroke='hsl(var(--logo-fill))' stroke-width='6' />
      <path d='M65 0L65 28' stroke='hsl(var(--logo-fill))' stroke-width='6' />
      <path
        fill-rule='evenodd'
        clip-rule='evenodd'
        d='M101 2.44784e-06L73 0L73 6L84 6L84 28H90V6L101 6V2.44784e-06Z'
        fill='hsl(var(--logo-fill))'
      />
      <path
        fill-rule='evenodd'
        clip-rule='evenodd'
        d='M33 0V23V28H39H57V23L39 23L39 0H33Z'
        fill='hsl(var(--logo-fill))'
      />
      <path
        fill-rule='evenodd'
        clip-rule='evenodd'
        d='M0 17V14V11H3H24V17H6V22H28V28H3H0V25V22V17Z'
        fill='#06B6D4'
      />
      <path
        fill-rule='evenodd'
        clip-rule='evenodd'
        d='M106 17V14V11H109H130V17H112V22H134V28H109H106V25V22V17Z'
        fill='hsl(var(--logo-fill))'
      />
    </svg>
  );
}

export function IconElitesSmall(props: IconProps) {
  return (
    <svg
      width='28'
      height='28'
      class={cn(props.class)}
      viewBox='0 0 28 28'
      xmlns='http://www.w3.org/2000/svg'
    >
      <path d='M0 3H28' stroke='#06B6D4' stroke-width='6' />
      <path
        fill-rule='evenodd'
        clip-rule='evenodd'
        d='M0 17V14V11H3H24V17H6V22H28V28H3H0V25V22V17Z'
        fill='#06B6D4'
      />
    </svg>
  );
}

export default { IconElites, IconElitesSmall };
