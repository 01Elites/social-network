import { JSXElement } from "solid-js";
import Layout from "../../Layout";
import { Button } from "../../components/ui/button";

export default function HomePage(): JSXElement {
  return (
    <Layout>
      <h1 class="p-5">Home Page</h1>
      <Button>Click!</Button>
      <Button variant={"outline"}>outline</Button>
      <Button variant={"ghost"}>ghost</Button>
      <Button variant={"secondary"}>DO NOT</Button>
    </Layout>
  );
}
