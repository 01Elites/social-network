import { createContext } from "solid-js";
import { UserDetailsHook } from "../../types/User";


const UserDetailsContext = createContext<UserDetailsHook>();

export default UserDetailsContext;