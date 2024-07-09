import 'solid-devtools';
import { JSXElement } from 'solid-js';
import Feed from '~/components/Feed';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <section class='flex h-full gap-4'>
        <div class='w-1/3 max-w-60'>
          <h1 class='text-xl font-bold'>Events</h1>
        </div>
        <Feed class='grow' />
        <div class='w-1/3 max-w-52'>
          <h1 class='text-xl font-bold'>Contacts</h1>
        </div>
      </section>
    </Layout>
  );
}
