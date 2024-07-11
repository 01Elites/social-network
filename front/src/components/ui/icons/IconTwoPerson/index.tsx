import { JSXElement } from 'solid-js';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function IconTwoPerson(props?: inputs): JSXElement {
  return (
    <div class={props?.class} onClick={props?.onClick}>
      <svg
        style={{ width: '100%', height: '100%' }}
        viewBox='0 0 41 38'
        fill='none'
        xmlns='http://www.w3.org/2000/svg'
      >
        <mask id='path-1-inside-1_0_1' fill='white'>
          <path d='M0 30C0 26.134 3.13401 23 7 23H24C27.866 23 31 26.134 31 30V38H0V30Z' />
        </mask>
        <path
          d='M-2 30C-2 25.0294 2.02944 21 7 21H24C28.9706 21 33 25.0294 33 30H29C29 27.2386 26.7614 25 24 25H7C4.23858 25 2 27.2386 2 30H-2ZM31 38H0H31ZM-2 38V30C-2 25.0294 2.02944 21 7 21V25C4.23858 25 2 27.2386 2 30V38H-2ZM24 21C28.9706 21 33 25.0294 33 30V38H29V30C29 27.2386 26.7614 25 24 25V21Z'
          fill='hsl(var(--primary))'
          mask='url(#path-1-inside-1_0_1)'
        />
        <path
          d='M24 10C24 14.9706 19.9706 19 15 19C10.0294 19 6 14.9706 6 10C6 5.02944 10.0294 1 15 1C19.9706 1 24 5.02944 24 10Z'
          stroke='hsl(var(--primary))'
          stroke-width='2'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M22.8484 17.7073C22.3036 18.2621 21.7004 18.7593 21.0488 19.1891C22.2609 19.7109 23.5967 20 25 20C30.5229 20 35 15.5228 35 10C35 4.47715 30.5229 0 25 0C23.5967 0 22.2608 0.289075 21.0488 0.810926C21.7004 1.24071 22.3036 1.73793 22.8484 2.29266C23.533 2.10195 24.2546 2 25 2C29.4183 2 33 5.58172 33 10C33 14.4183 29.4183 18 25 18C24.2546 18 23.533 17.8981 22.8484 17.7073Z'
          fill='hsl(var(--primary))'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M30.7453 25H34C36.7614 25 39 27.2386 39 30V38H41V30C41 26.134 37.866 23 34 23H28.6076C29.4531 23.5095 30.1821 24.1926 30.7453 25Z'
          fill='hsl(var(--primary))'
        />
      </svg>
    </div>
  );
}
