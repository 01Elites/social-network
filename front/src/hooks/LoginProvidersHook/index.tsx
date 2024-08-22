import config from '~/config';
import { LoginProvider } from '~/types/LoginProvider';

import githubImage from '~/assets/github_logo.png';
import googleImage from '~/assets/google_logo.png';
import giteaImage from '~/assets/reboot_01_logo.png';

type LoginProvidersHook = {
  providers: LoginProvider[];
};

export default function useLoginProviders(): LoginProvidersHook {
  const providers: LoginProvider[] = [
    {
      name: 'Reboot',
      icon: giteaImage,
      onClick: () => {
        window.location.replace(config.API_URL + '/auth/gitea/login');
      },
    },
    {
      name: 'Github',
      icon: githubImage,
      onClick: () => {
        window.location.replace(config.API_URL + '/auth/github/login');
      },
    },
    {
      name: 'Google',
      icon: googleImage,
      onClick: () => {
        window.location.replace(config.API_URL + '/auth/google/login');
      },
    },
  ];

  return { providers };
}
