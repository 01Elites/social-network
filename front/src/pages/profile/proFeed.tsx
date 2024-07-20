import { Tabs } from "@kobalte/core/tabs";
import { JSXElement } from 'solid-js';
import "./style.css";
import { cn } from "~/lib/utils";
import config from "~/config";
import FeedPosts from "~/components/Feed/FeedPosts";
import { useParams } from "@solidjs/router";

type ProfileParams = {
  username: string;
};

export default function ProfileFeed(): JSXElement {

  const params: ProfileParams = useParams();

  return (
    <Tabs aria-label="Main navigation" class="tabs">
      <Tabs.List class="tabs__list">
        <Tabs.Trigger class="tabs__trigger" value="posts">Posts</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill">ما ادري ويش</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill2">وش تمبا</Tabs.Trigger>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="posts">
        <div class={cn('flex flex-col gap-4 p-2')}>
          <FeedPosts path={`/profile/${params.username}/posts`} />
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill2">still NOTHING!!!</Tabs.Content>
    </Tabs>
  )
}




