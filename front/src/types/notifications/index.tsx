import { FriendRequest } from '~/types/friends';
import { GroupEvent, GroupInvite } from '~/types/group';

export type Notifications = {
  event: string;
  id: number;
  payload: any;
}