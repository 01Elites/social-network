import { createContext } from 'solid-js';
import { UserDetailsHook } from '~/hooks/userDetails';

const UserDetailsContext = createContext<UserDetailsHook>();

export default UserDetailsContext;
