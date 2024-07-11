import { JSXElement, useContext } from 'solid-js';
import { AspectRatio } from "~/components/ui/aspect-ratio"
import { Button } from "~/components/ui/button"
import follow from "~/assets/icons_svgs/follow.svg"
import globe from "~/assets/icons_svgs/globe.svg"
import github from "~/assets/icons_svgs/github.svg"

interface params {
  username: string;
}
export default function ProfileDetails(params: params): JSXElement {


  return (
    <div class='flex flex-col'> {/* Left div */}
      <div class='flex flex-col justify-center items-center'>
        <AspectRatio ratio={16 / 9} class='m-2'> {/* Profile picture */}
          <div class='absolute inset-0 bg-black bg-opacity-50 flex justify-center items-center rounded-lg'>
            IZ NOT HERE
          </div>
        </AspectRatio> {/* Profile picture */}
        <div class='flex flex-col items-center w-full'> {/* Username, followers, following */}
          <p class='text-2xl font-bold m-2'>{params.username}</p>
          <div class='grid w-full grid-cols-2 text-sm m-2'>
            <p class='flex justify-center'>Followers -2</p>
            <p class='flex justify-center'>Following 9999999</p>
          </div>
        </div> {/* Username, followers, following */}
        <div class='m-4'> {/* Bio */}
          <p>I am the go-to one stop and the best lorem ipsum you can find in the not sum of the roughsum in the dorem</p>
        </div> {/* Bio */}
        <div class='flex flex-row w-full justify-between gap-2 m-4'>
          <Button class="flex grow" variant="default">
            <img
              src={follow}
              alt='Apps'
              onClick={() => { }}
              class='m-2'
            />Follow
          </Button>
          <div class='flex gap-2'> {/* Follow button */}
            <Button variant="default">
              <img
                src={globe}
                alt='Apps'
                onClick={() => { }}
                class='m-2'
              />
            </Button>
            <Button variant="default">
              <img
                src={github}
                alt='Apps'
                onClick={() => { }}
                width={20}
                class='m-2'
              />
            </Button>
          </div> {/* Follow button */}
        </div>
      </div>
    </div>
  )
}
