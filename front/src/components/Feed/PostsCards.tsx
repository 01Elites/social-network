import { JSXElement, useContext } from "solid-js";
import {
  Card,
  CardContent,
  // CardDescription,
  CardFooter,
  // CardHeader,
  // CardTitle
} from "~/components/ui/card"
import sample_img from "../../assets/sample.avif"
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import UserDetailsContext from "~/contexts/UserDetailsContext";
import { UserDetailsHook } from "~/types/User";

export default function PostsCards(): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  return (
    <Card class="flex items-center justify-center flex-col bg-green">
      <div>
        <CardContent>
          <div class="flex grow justify-center">
            <img src={sample_img} alt="Example Image" width={400} class="flex align-center text-center" />
          </div>
          <div class="flex justify-center">
            <div class="flex flex-row items-center justify-between w-[400px]">
              <Avatar>
                <AvatarImage src={userDetails()?.avatar_url} />
                <AvatarFallback>
                  {userDetails()?.first_name.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <div class="flex flex-col">
                <p class="border">Mohammad AlSammak</p>
                <p class="border">1y ago</p>
              </div>
            </div>
          </div>
          <p class="border">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean in ipsum at ante blandit ornare. Duis est nisl, lacinia vitae tempor eget, hendrerit non libero. Nunc ullamcorper orci mauris, eu vehicula nunc aliquam sit amet. Morbi luctus, ipsum eu euismod convallis, orci justo pellentesque lectus, vitae mollis ante sem eu nunc. Duis blandit cursus scelerisque. Donec id dui quam. Nunc pulvinar massa quis nunc blandit posuere. Vivamus et augue nisl. Nunc quis euismod neque. Suspendisse ac consequat nunc. Aenean volutpat sapien ut risus maximus pellentesque eget fermentum metus. Sed dignissim elementum iaculis. Curabitur ac erat ut tellus iaculis viverra. Sed efficitur at erat id vestibulum. Curabitur scelerisque, nisl id sollicitudin fringilla, dolor purus varius neque, quis imperdiet nunc nulla a metus.
          </p>
        </CardContent>
        <CardFooter>
          <p>Likes and comments</p>
        </CardFooter>
      </div>
    </Card >
  );
}
