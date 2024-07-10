import 'solid-devtools';
import { JSXElement } from 'solid-js';
import Contacts from '~/components/Contacts';
import Feed from '~/components/Feed';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <section class='flex h-full gap-4 px-3'>
        <div class='w-1/3 max-w-60'>
          <h1 class='text-xl font-bold'>Events</h1>
        </div>
        <Feed class='grow' />
        <Contacts class='w-1/3 max-w-52' />
      </section>
    </Layout>
  );
}
