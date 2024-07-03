import 'solid-devtools';
import { JSXElement } from 'solid-js';
import Feed from '~/components/Feed';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <section class='flex gap-4'>
        <Feed class='w-3/4' />
        <div class='grow'>
          <h1 class='text-xl font-bold'>Contacts</h1>
        </div>
      </section>
    </Layout>
  );
}
