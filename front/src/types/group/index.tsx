import { requester } from "~/pages/group/groupfeed";
import { User } from "~/types/User";

type Group = {
  id: number,
  title: string;
  description: string;
  members: User[];
  ismember: boolean;
  iscreator: boolean;
  request_made: boolean;
  events: GroupEvent[];
  creator: creator;
  invited_by: requester;
  requesters: requester[];
  explore: User[];
};

type creator = {
  user_name: string;
  first_name: string;
  last_name: string;
  avatar?: string;
};

type GroupEvent = {
  id: number;
  title: string;
  description: string;
  options: option[];
  event_time: number
  responded_users: string[] | undefined;
  choices: string[] | undefined;
  full_names: string[] | undefined;
};

type option = {
  option_id: number;
  option_name: string;
};

export type { Group, GroupEvent }


