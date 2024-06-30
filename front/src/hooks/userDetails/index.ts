import config from "~/config";
import User from "~/types/User/User";



export default async function userDetails(): Promise<User|Error> {
    return new Promise((resolve, reject) => {
        fetch(config.API_URL + '/profile')
            .then(async (response) => {
                if (!response.ok) {
                    const body = await response.json();
                    if (body.reason) {
                        throw new Error(body.reason);
                    }
                    throw new Error('Failed to fetch your information.');
                }
                resolve(await response.json());
            })
            .catch((error) => {
                reject(error);
            });
    });
}