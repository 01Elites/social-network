import { A, useParams } from '@solidjs/router';
import { createEffect, createSignal, JSXElement, Show } from 'solid-js';
import { buttonVariants } from '~/components/ui/button';

function getCookieValue(name: string) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop()?.split(';').shift() as string;
  return null;
}

function deleteCookie(name: string) {
  // Set the cookie's expiration date to a past date to delete it
  document.cookie = `${name}=; Expires=Thu, 01 Jan 1970 00:00:01 GMT; Path=/; SameSite=Strict; Secure`;
}

export default function PostLogin(): JSXElement {
  const [tokenError, setTokenError] = createSignal<string>(useParams().error);

  // extract token from cookie
  createEffect(() => {
    const token = getCookieValue('SN_SESSION');
    if (!token) {
      setTokenError('Token not found');
      return;
    }
    localStorage.setItem('SN_TOKEN', token);
    deleteCookie('SN_SESSION');

    window.location.href = '/';
  });

  return (
    <section class='mx-auto flex h-full max-w-80 flex-col items-center justify-center gap-4 p-4'>
      <Show when={tokenError()}>
        <h1 class='text-lg font-bold'>Login failed</h1>
        <p class='text-primary/90'>{tokenError()}</p>
        <A class={buttonVariants({ variant: 'default' })} href='/'>
          Go to home
        </A>
      </Show>
    </section>
  );
}
