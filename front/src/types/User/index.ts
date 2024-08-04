export type User = {
  user_name: string;
  email: string;
  nick_name?: string;
  first_name: string;
  last_name: string;
  gender: 'male' | 'female';
  date_of_birth: number; // unix timestamp
  avatar?: string;
  about?: string;
  profile_privacy: 'public' | 'private';
  follow_status: string;
  follower_count: number;
  following_count: number;
};

export default User;
