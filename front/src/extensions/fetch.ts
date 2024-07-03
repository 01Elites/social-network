/// <reference lib="dom" />

export async function fetchWithAuth(
  input: RequestInfo | URL,
  init?: RequestInit,
): Promise<Response> {
  const token = localStorage.getItem('SN_TOKEN') || '';

  // Add the Authorization header to the request
  const options: RequestInit = init || {};

  console.log('fetchWithAuth', input, options);

  if (token !== '') {
    options.headers = {
      ...options.headers,
      Authorization: `Bearer ${token}`,
    };
  }


  console.log('fetchWithAuth after', input, options);

  // Perform the fetch
  return fetch(input, options).then((response) => {
    // Check for new token in the response headers
    const newToken = response.headers.get('Authorization');
    if (newToken) {
      localStorage.setItem('SN_TOKEN', newToken.replace('Bearer ', ''));
    }
    return response;
  });
}
