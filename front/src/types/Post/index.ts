import User from '../User';

type Post = {
  post_id: number;
  title: string;
  content: string;
  image?: string;
  likers_usernames?: string[];
  comments_count: number;
  creation_date: number; // unix time
  poster: User;
};

export type { Post };
