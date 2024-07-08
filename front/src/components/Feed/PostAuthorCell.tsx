import moment from 'moment';
import { JSXElement } from 'solid-js';
import User from '~/types/User';
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';

interface PostAuthorCellProps {
  author: User;
  date: Date;
}

export default function PostAuthorCell(props: PostAuthorCellProps): JSXElement {
  return (
    <div class='flex items-center gap-2'>
      <Avatar>
        <AvatarImage src={props.author.avatar_url}></AvatarImage>
        <AvatarFallback>
          {props.author.first_name[0].toUpperCase()}
        </AvatarFallback>
      </Avatar>
      <div>
        <h3 class='text-sm font-bold'>{`${props.author.first_name} ${props.author.last_name}`}</h3>
        <time
          class='text-xs font-light text-muted-foreground'
          dateTime={moment(props.date).calendar()}
          title={moment(props.date).calendar()}
        >
          {moment(props.date).fromNow()}
        </time>
      </div>
    </div>
  );
}
