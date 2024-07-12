import { useParams } from '@solidjs/router';
import { JSXElement } from 'solid-js';
import ProfileDetails from './profileDetails';
import Layout from '~/Layout';

type ProfileParams = {
  username: string;
};

export default function Profile(): JSXElement {


  const params: ProfileParams = useParams();
  return (
    <Layout>
      <div class='grid grid-cols-1 md:grid-cols-6 m-4 '> {/* Main grid */}
        <div class='col-span-2'>
          <ProfileDetails username={params.username} />
        </div>
        <div class='col-span-4'>
          div 2
        </div>
      </div> {/* Main grid */}
    </Layout>
  )
}
