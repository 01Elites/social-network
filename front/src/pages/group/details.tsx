import { JSXElement, Show } from "solid-js";
import type {Group} from "~/types/group";
import { A } from '@solidjs/router';
import { AspectRatio } from "~/components/ui/aspect-ratio";
import config from '~/config';
import RequestToJoin from "./request";


export default function GroupDetails(props: { targetGroup: () => Group}): JSXElement {
  const numberOfMembers = props.targetGroup().members.length;
  const groupID = String(props.targetGroup().id)
  return (
    <div class='flex flex-col'> {/* Left div */}
      <div class='flex flex-col justify-center items-center'>
        <Show when={props.targetGroup().creator.avatar}>
        <AspectRatio ratio={16 / 9}>
          <img
            class='size-full rounded-md rounded-b-none object-cover'
            loading='lazy'
            src={`${config.API_URL}/image/${props.targetGroup().creator.avatar}`}
          />
        </AspectRatio>
      </Show>
         {/* Profile picture */}
        <div class='flex flex-col items-center w-full'> {/* Username, followers, following */}
          <p class='text-2xl font-bold m-2'>{props.targetGroup().title}</p>
          <div class='grid w-full grid-cols-2 text-sm m-2'>
          <p class='flex justify-center'>Group Creator:&nbsp<A
              href={"/profile/" + props.targetGroup().creator.user_name} class='block text-sm font-bold hover:underline'>
          {props.targetGroup().creator.first_name}  {props.targetGroup().creator.last_name}</A></p>
            <p class='flex justify-center'>Number of Members: {numberOfMembers}</p>
          </div>
        </div> {/* Username, followers, following */}
        <div class='m-4'> {/* Bio */}
          <p>{props.targetGroup().description}</p>
        </div>
       <RequestToJoin targetGroup={() => props.targetGroup() as Group}/>
      </div>
    </div>
  )
}