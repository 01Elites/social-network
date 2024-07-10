type Post = {
  post_id: number;
  title: string;
  content: string;
  image?: string;
  likers_usernames?: string[];
  comments_count: number;
  creation_date: number; // unix time
  poster: {
    image?: string;
    first_name: string;
    last_name: string;
    user_name: string;
  };
};

export type { Post };
