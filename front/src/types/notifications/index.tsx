import { FriendRequest } from '~/types/friends';
import { GroupEvent, GroupInvite } from '~/types/group';

export type Notifications = {
  FollowRequest: FriendRequest[],
  Events: GroupEvent[];
  GroupInvite: GroupInvite[];
  GroupRequests: GroupInvite[];
}