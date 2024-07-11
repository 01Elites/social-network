interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function IconGroup(props: inputs) {
  return (
    <div class={props.class} onClick={props.onClick}>
      <svg
        style={{ width: '100%', height: '100%' }}
        viewBox='0 0 51 38'
        fill='none'
        xmlns='http://www.w3.org/2000/svg'
      >
        <mask id='path-1-inside-1_27_4' fill='white'>
          <path d='M10 30C10 26.134 13.134 23 17 23H34C37.866 23 41 26.134 41 30V38H10V30Z' />
        </mask>
        <path
          d='M8 30C8 25.0294 12.0294 21 17 21H34C38.9706 21 43 25.0294 43 30H39C39 27.2386 36.7614 25 34 25H17C14.2386 25 12 27.2386 12 30H8ZM41 38H10H41ZM8 38V30C8 25.0294 12.0294 21 17 21V25C14.2386 25 12 27.2386 12 30V38H8ZM34 21C38.9706 21 43 25.0294 43 30V38H39V30C39 27.2386 36.7614 25 34 25V21Z'
          fill='hsl(var(--primary))'
          mask='url(#path-1-inside-1_27_4)'
        />
        <path
          d='M34 10C34 14.9706 29.9706 19 25 19C20.0294 19 16 14.9706 16 10C16 5.02944 20.0294 1 25 1C29.9706 1 34 5.02944 34 10Z'
          stroke='hsl(var(--primary))'
          stroke-width='2'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M32.8484 17.7073C32.3036 18.2621 31.7004 18.7593 31.0488 19.1891C32.2609 19.7109 33.5967 20 35 20C40.5229 20 45 15.5228 45 10C45 4.47715 40.5229 0 35 0C33.5967 0 32.2608 0.289075 31.0488 0.810926C31.7004 1.24071 32.3036 1.73793 32.8484 2.29266C33.533 2.10195 34.2546 2 35 2C39.4183 2 43 5.58172 43 10C43 14.4183 39.4183 18 35 18C34.2546 18 33.533 17.8981 32.8484 17.7073Z'
          fill='hsl(var(--primary))'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M40.7453 25H44C46.7614 25 49 27.2386 49 30V38H51V30C51 26.134 47.866 23 44 23H38.6076C39.4531 23.5095 40.1821 24.1926 40.7453 25Z'
          fill='hsl(var(--primary))'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M17.1516 17.7073C17.6964 18.2621 18.2996 18.7593 18.9512 19.1891C17.7391 19.7109 16.4033 20 15 20C9.47711 20 4.99996 15.5228 4.99996 10C4.99996 4.47715 9.47711 0 15 0C16.4033 0 17.7392 0.289075 18.9512 0.810926C18.2996 1.24071 17.6964 1.73793 17.1516 2.29266C16.467 2.10195 15.7454 2 15 2C10.5817 2 6.99996 5.58172 6.99996 10C6.99996 14.4183 10.5817 18 15 18C15.7454 18 16.467 17.8981 17.1516 17.7073Z'
          fill='hsl(var(--primary))'
        />
        <path
          fill-rule='evenodd'
          clip-rule='evenodd'
          d='M10.2547 25H6.99998C4.23856 25 1.99998 27.2386 1.99998 30V38H-1.79546e-05V30C-1.79546e-05 26.134 3.13399 23 6.99998 23H12.3924C11.5469 23.5095 10.8179 24.1926 10.2547 25Z'
          fill='hsl(var(--primary))'
        />
      </svg>
    </div>
  );
}
