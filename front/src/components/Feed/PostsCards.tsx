import { JSXElement } from "solid-js";
import {
  Card,
  CardContent,
  // CardDescription,
  CardFooter,
  // CardHeader,
  // CardTitle
} from "~/components/ui/card"

export default function PostsCards(): JSXElement {
  return (
    <>
      <Card>
        <CardContent>
          <p>Image</p>
          <p>Author</p>
          <p>Text</p>
        </CardContent>
        <CardFooter>
          <p>Likes and comments</p>
        </CardFooter>
      </Card>
    </>
  );
}
