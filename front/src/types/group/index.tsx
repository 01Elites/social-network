type Group = {
  id: number,
  title: string;
  description: string;
  members: string[];
  ismember: boolean;
  request_made: boolean;
  events: groupEvent[];
  creator: creator;
};

type creator = {
  user_name: string;
  first_name: string;
  last_name: string;
  avatar?: string;
};

type groupEvent = {
  id: number;
  title: string;
  description: string;
  options: string[];
  event_time: number
  responded_users: string[];
};

export type {Group}