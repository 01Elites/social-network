import 'solid-devtools';
import { JSXElement } from 'solid-js';
import Feed from '~/components/Feed';
import HomeContacts from '~/components/HomeContacts';
import HomeEvents from '~/components/HomeEvents';
import Layout from '../../Layout';

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <section class='flex h-full gap-4'>
        <HomeEvents class='hidden w-5/12 max-w-60 md:block' />
        <Feed class='grow overflow-hidden' />
        <HomeContacts class='hidden w-1/3 max-w-52 md:block' />
      </section>
    </Layout>
  );
}
