import { useParams } from '@solidjs/router';
import { JSXElement, useContext } from 'solid-js';
import Layout from '~/Layout';
import { AspectRatio } from "~/components/ui/aspect-ratio"
import { Button } from "~/components/ui/button"
import { Col, Grid } from "~/components/ui/grid"
import ProfileDetails from './profileDetails';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import { UserDetailsHook } from '~/types/User';

type ProfileParams = {
  username: string;
};

export default function Profile(): JSXElement {


  const params: ProfileParams = useParams();
  return (
    <div class='grid grid-cols-1 md:grid-cols-6 m-4 '> {/* Main grid */}
      <div class='col-span-2'>
        <ProfileDetails username={params.username} />
      </div>
      <div class='col-span-4'>
        div 2
      </div>
    </div> /* Main grid */
  )
}
