import { createSignal, onMount } from 'solid-js';
import config from '../../config';
import User from '../../types/User';


function useUserDetails(): {
  userDetails: () => User | null;
  error: () => string | null;
  fetchUserDetails: () => Promise<void>;
  updateUserDetails: (partialDetails: Partial<User>) => Promise<void>;
} {
  const [userDetails, setUserDetails] = createSignal<User | null>(null);
  const [userDetailsError, setError] = createSignal<string | null>(null);

  async function fetchUserDetails() {
    try {
      const response = await fetch(config.API_URL + '/api/profile');
      if (!response.ok) {
        throw new Error('Failed to fetch user details.');
      }
      const data: User = await response.json();
      setUserDetails(data);
    } catch (err) {
      setError((err as Error).message);
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
      setError((err as Error).message);
    }
  }

  onMount(fetchUserDetails);

  return {
    userDetails,
    error: userDetailsError,
    fetchUserDetails,
    updateUserDetails,
  };
}

export default useUserDetails;
