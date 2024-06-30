import config from "~/config";

export default async function Login(email: string, password: string): Promise<boolean> {

  let login_fetch = fetch(config.API_URL + '/api/auth/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  })

  login_fetch.then(response => {
    if (response.status === 200) {
      return true
    } else {
      return false
    }
  })

  return false

}
