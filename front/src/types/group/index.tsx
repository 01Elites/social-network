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

export type GroupInvite = {
  invited_by: requester;
  group_id: number
}

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
  event_time: number;
  creator: creator;
};

type option = {
  option_id: number;
  option_name: string;
  usernames: string[] | null;
  fullnames: string[] | null;
};

export type { Group, GroupEvent }


