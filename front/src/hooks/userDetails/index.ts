import { createSignal, onMount } from 'solid-js';
import config from '../../config';
import User from '../../types/User';

interface UseUserDetails {
  userDetailsData: () => User | null;
  error: () => string | null;
  fetchUserDetails: () => Promise<void>;
  updateUserDetails: (partialDetails: Partial<User>) => Promise<void>;
}

function useUserDetails(): UseUserDetails {
  const [userDetailsData, setUserDetailsData] = createSignal<User | null>(null);
  const [error, setError] = createSignal<string | null>(null);

  async function fetchUserDetails() {
    try {
      const response = await fetch(config.API_URL + '/api/profile');
      if (!response.ok) {
        throw new Error('Failed to fetch user details.');
      }
      const data: User = await response.json();
      setUserDetailsData(data);
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
      setUserDetailsData(updatedData);
    } catch (err) {
      setError((err as Error).message);
    }
  }

  onMount(fetchUserDetails);

  return {
    userDetailsData,
    error,
    fetchUserDetails,
    updateUserDetails,
  };
}

export default useUserDetails;
