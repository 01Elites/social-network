type User = {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  date_of_birth: number; // unix timestamp
  avatar_url?: string;
  about?: string;
  nick_name?: string;
  profile_privacy?: 'public' | 'private';
};

export default User;
