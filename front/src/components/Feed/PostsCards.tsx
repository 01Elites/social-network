import { JSXElement } from "solid-js";
import {
  Card,
  CardContent,
  // CardDescription,
  CardFooter,
  // CardHeader,
  // CardTitle
} from "~/components/ui/card"
import sample_img from "../../assets/sample.avif"

export default function PostsCards(): JSXElement {
  return (
    <>
      <Card>
        <div class="flex items-center justify-center flex-col bg-green">
          <CardContent>
            <img src={sample_img} alt="Example Image" width={400} class="flex align-center text-center" />
            <p class="border">Author</p>
            <p class="border">Text</p>

          </CardContent>
          <CardFooter>
            <p>Likes and comments</p>
          </CardFooter>
        </div>

      </Card>
    </>
  );
}
