import { useNavigate } from '@solidjs/router';
import config from '~/config';
import { LoginProvider } from '~/types/LoginProvider';

import giteaImage from '~/assets/reboot_01_logo.png';
import githubImage from '~/assets/github_logo.png';
import googleImage from '~/assets/google_logo.png';

type LoginProvidersHook = {
  providers: LoginProvider[];
  postLogin: () => void;
};

export default function useLoginProviders(): LoginProvidersHook {
  const providers: LoginProvider[] = [
    {
      name: 'Reboot',
      icon: giteaImage,
      onClick: () => {
        window.location.href = config.API_URL + '/auth/gitea/login';
      }
    },
    {
      name: 'Github',
      icon: githubImage,
      onClick: () => {
        window.location.href = config.API_URL + '/auth/github/login';
      }
    },
    {
      name: 'Google',
      icon: googleImage,
      onClick: () => {
        window.location.href = config.API_URL + '/auth/google/login';
      }
    },
  ];

  function postLogin() {
    // check the last path if it was a login path
    console.log('Post llgin')
  }


  return { providers, postLogin };
}
