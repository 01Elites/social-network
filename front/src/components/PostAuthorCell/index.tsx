import { A } from '@solidjs/router';
import moment from 'moment';
import { JSXElement, Show } from 'solid-js';
import config from '~/config';
import User from '~/types/User';
import { Avatar, AvatarFallback } from '../ui/avatar';

interface PostAuthorCellProps {
  author: User;
  date: Date;
}

export default function PostAuthorCell(props: PostAuthorCellProps): JSXElement {
  return (
    <div class='flex items-center gap-2'>
      <Avatar>
        <AvatarFallback>
          <Show
            when={props.author.avatar}
            fallback={props.author.first_name.charAt(0).toUpperCase()}
          >
            <img
              alt='avatar'
              class='size-full rounded-md rounded-b-none object-cover'
              loading='lazy'
              src={`${config.API_URL}/image/${props.author.avatar}`}
            />
          </Show>
        </AvatarFallback>
      </Avatar>
      <div>
        <A
          href={`/profile/${props.author.user_name}`}
          class='block text-sm font-bold hover:underline'
        >{`${props.author.first_name} ${props.author.last_name}`}</A>
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
