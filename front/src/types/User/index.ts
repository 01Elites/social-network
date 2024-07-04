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

interface UserDetailsHook {
  userDetails: () => User | null;
  setUserDetails: (details: User | null) => void;
  userDetailsError: () => string | null;
  fetchUserDetails: () => Promise<void>;
  updateUserDetails: (partialDetails: Partial<User>) => Promise<void>;
}

export default User;
export type { UserDetailsHook };
