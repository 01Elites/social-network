type Group = {
  id: number,
  title: string;
  description: string;
  members: string[];
  ismember: boolean;
  events: groupEvent[];
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