import User from '../User';

type Comment = {
  comment_id: number;
  image?: string;
  body: string;
  creation_date: Date; // unix timestamp
  commenter: User;
};

export type { Comment };
