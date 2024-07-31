import { JSXElement, Show } from "solid-js";
import type {Group} from "~/types/group";
import { A } from '@solidjs/router';
import { AspectRatio } from "~/components/ui/aspect-ratio";
import config from '~/config';
import RequestToJoin from "./request";
import { Avatar, AvatarFallback } from "~/components/ui/avatar";


export default function GroupDetails(props: { targetGroup: () => Group}): JSXElement {
  const numberOfMembers = props.targetGroup().members.length;
  const groupID = String(props.targetGroup().id)
  return (
    <div class='flex flex-col w-[30rem]'> {/* Left div */}
      <div class='flex flex-col justify-center items-center'>
        <AspectRatio ratio={16 / 9}>
          <div class='absolute inset-0 bg-black bg-opacity-50 flex justify-center items-end rounded-lg'>
          <Avatar class='w-[5rem] h-[5rem] mb-2'>
              <AvatarFallback>
                <Show when={props.targetGroup().creator.avatar} fallback={
                  props.targetGroup().title.charAt(0).toUpperCase()
                }><img
                    class='size-full rounded-md rounded-b-none object-cover'
                    loading='lazy'
                    src={`${config.API_URL}/image/${props.targetGroup().creator.avatar}`}
                  /></Show></AvatarFallback>
            </Avatar>
            </div>
        </AspectRatio>
         {/* Profile picture */}
        <div class='flex flex-col items-center w-full'> {/* Username, followers, following */}
          <p class='text-2xl font-bold m-4'>{props.targetGroup().title}</p>
          <div class='grid w-full grid-cols-2 text-sm m-4'>
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