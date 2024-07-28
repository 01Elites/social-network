import User from '~/types/User';

type Friends = {
  user_name: string;
  followers: User[];
  following: User[];
  friend_requests: FriendRequest[];
  explore: User[];
};

type FriendRequest = {
  requester: string;
  creation_date: string;
  user_info: User;
};

export default Friends;
