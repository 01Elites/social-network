type User = {
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
};

interface UserDetailsHook {
  userDetails: () => User | null;
  setUserDetails: (details: User | null) => void;
  userDetailsError: () => string | null;
  fetchUserDetails: () => Promise<void>;
  updateUserDetails: (partialDetails: Partial<User>) => Promise<void>;
}

export default User;
export type { UserDetailsHook };
