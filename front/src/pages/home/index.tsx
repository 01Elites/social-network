import 'solid-devtools';
import { JSXElement } from 'solid-js';
import Feed from '~/components/Feed';
import HomeContacts from '~/components/HomeContacts';
import HomeEvents from '~/components/HomeEvents';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <section class='flex h-full gap-4 px-3'>
        <HomeEvents class='w-1/3 max-w-60' />
        <Feed class='grow' />
        <HomeContacts class='w-1/3 max-w-52' />
      </section>
    </Layout>
  );
}
