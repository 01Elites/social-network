import { createSignal, onMount } from 'solid-js';
import config from '../../config';
import User, { UserDetailsHook } from '../../types/User';
import { fetchWithAuth } from '~/extensions/fetch';



function useUserDetails(): UserDetailsHook {
  const [userDetails, setUserDetails] = createSignal<User | null>(null);
  const [userDetailsError, setUserDetailsError] = createSignal<string | null>(null);

  async function fetchUserDetails(): Promise<void> {
    try {
      const response = await fetchWithAuth(config.API_URL + '/profile');
      
      if (!response.ok) {
        const error = await response.json();
        if (error.reason) {
          throw new Error(error.reason);
        }
        throw new Error('Failed to fetch user details. Please check your network connection.');
      }

      const data: User = await response.json();

      setUserDetails(data);
      setUserDetailsError(null);
    } catch (err) {
      setUserDetailsError((err as Error).message);
    }
  }

  async function updateUserDetails(partialDetails: Partial<User>) {
    try {
      const response = await fetch(config.API_URL + '/api/profile', {
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
  }

  onMount(fetchUserDetails);

  return {
    userDetails,
    userDetailsError,
    fetchUserDetails,
    updateUserDetails,
  };
}

export {useUserDetails};
