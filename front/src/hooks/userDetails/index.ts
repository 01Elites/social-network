import { Accessor, createSignal, onMount } from 'solid-js';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '../../config';
import User from '../../types/User';

function useUserDetails(): UserDetailsHook {
  const [userDetails, setUserDetails] = createSignal<User | null>(null);
  const [userDetailsError, setUserDetailsError] = createSignal<string | null>(
    null,
  );
  const [proccessing, setProccessing] = createSignal(false);

  async function fetchUserDetails(): Promise<void> {
    setProccessing(true);
    try {
      const response = await fetchWithAuth(config.API_URL + '/profile');

      if (!response.ok) {
        const error = await response.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error(
          'Failed to fetch user details. Please check your network connection.',
        );
      }

      const data: User = await response.json();

      setUserDetails(data);
      setUserDetailsError(null);
    } catch (err) {
      setUserDetailsError((err as Error).message);
    }
    setProccessing(false);
  }

  async function updateUserDetails(partialDetails: Partial<User>) {
    setProccessing(true);
    try {
      const response = await fetchWithAuth(config.API_URL + '/api/profile', {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(partialDetails),
      });
      if (!response.ok) {
        throw new Error('Failed to update user details.');
      }
      const updatedData: User = await response.json();
      setUserDetails(updatedData);
    } catch (err) {
      setUserDetailsError((err as Error).message);
    }
    setProccessing(false);
  }

  onMount(fetchUserDetails);

  return {
    userDetails,
    setUserDetails,
    userDetailsError,
    fetchUserDetails,
    updateUserDetails,
    proccessing,
  };
}

interface UserDetailsHook {
  userDetails: Accessor<User | null>;
  setUserDetails: (details: User | null) => void;
  userDetailsError: Accessor<string | null>;
  fetchUserDetails: () => Promise<void>;
  updateUserDetails: (partialDetails: Partial<User>) => Promise<void>;
  proccessing: Accessor<boolean>;
}

export type { UserDetailsHook };

export { useUserDetails };
