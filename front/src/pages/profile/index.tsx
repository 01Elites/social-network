import { useParams } from '@solidjs/router';
import { JSXElement } from 'solid-js';

type ProfileParams = {
  username: string;
};

export default function Profile(): JSXElement {
  const params: ProfileParams = useParams();
  return <h1>{params.username} Profile</h1>;
}
