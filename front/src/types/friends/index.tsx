type Friends = {
  user_name: string;
  followers: string[];
  following: string[];
  friend_requests: FriendRequest[] | undefined;
  explore: string[] | undefined;
};

type FriendRequest = {
  requester: string;
  creation_date: string;
};

export default Friends;
