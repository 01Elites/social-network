type Friends = {
  user_name: string;
  followers: string[];
  following: string[];
  friend_requests: FriendRequest[];
  explore: string[];
};

type FriendRequest = {
  requester: string;
  creation_date: string;
};

export default Friends;
