import { JSXElement, Show } from "solid-js";
import type {Group} from "~/types/group";
import { AspectRatio } from "~/components/ui/aspect-ratio";
import { Avatar, AvatarFallback, AvatarImage } from '~/components/ui/avatar';
import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import Globe_Icon from '~/components/ui/icons/globe_icon';
import Message_Icon from '~/components/ui/icons/message_icon';
import { Switch, Match } from "solid-js"


export default function GroupDetails(props: { targetGroup: () => Group}): JSXElement {
  
  return (
    <div class='flex flex-col'> {/* Left div */}
      <div class='flex flex-col justify-center items-center'>
        <AspectRatio ratio={16 / 9}> {/* Profile picture */}
          <div class='absolute inset-0 bg-black bg-opacity-50 flex justify-center items-end rounded-lg'>
            <Avatar class='w-[5rem] h-[5rem] mb-2'>
              <AvatarFallback>{props.targetGroup().first_name[0]}</AvatarFallback>
            </Avatar>
          </div>
        </AspectRatio>
         {/* Profile picture */}
        <div class='flex flex-col items-center w-full'> {/* Username, followers, following */}
          <p class='text-2xl font-bold m-2'>{props.targetGroup().title}</p>
          <div class='grid w-full grid-cols-2 text-sm m-2'>
            <p class='flex justify-center'>{props.targetGroup().members}</p>
            <p class='flex justify-center'>Following 9999999</p>
          </div>
        </div> {/* Username, followers, following */}
        <div class='m-4'> {/* Bio */}
          <p>{props.targetGroup().description}</p>
        </div> {/* Bio */}
        <Switch>
  <Match when={condition1}>
    <p>Outcome 1</p>
  </Match>
  <Match when={condition2}>
    <p>Outcome 2</p>
  </Match>
</Switch>
          <Button class="flex grow" variant="default">
            <Follow_Icon /> {props.targetGroup().follow_status === 'following' ? 'Unfollow' : 'Follow'}
          </Button>
        <div class='flex flex-row w-full justify-between gap-2 m-4'>
          <div class='flex gap-2'> {/* Follow button */}
            <Button variant="default">
              <Globe_Icon class='w-5 justify-center' />
            </Button>
            <Button variant="default">
              <Message_Icon class='w-5 justify-center' />
            </Button>
          </div> {/* Follow button */}
        </div>
      </div>
    </div>
  )
}
/// create function for pending request 
