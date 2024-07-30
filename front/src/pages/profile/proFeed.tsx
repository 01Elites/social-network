import { Tabs } from "@kobalte/core/tabs";
import { JSXElement, Show } from 'solid-js';
import "./style.css";
import { cn } from "~/lib/utils";
import config from "~/config";
import FeedPosts from "~/components/Feed/FeedPosts";
import { useParams } from "@solidjs/router";
import User from "~/types/User";

type ProfileParams = {
  username: string;
};

export default function ProfileFeed(props: {
  targetUser: () => User;
}): JSXElement {

  const params: ProfileParams = useParams();

  return (
    <Tabs aria-label="Main navigation" class="tabs">
      <Tabs.List class="tabs__list">
        <Tabs.Trigger class="tabs__trigger" value="posts">Posts</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill">Followers</Tabs.Trigger>
        <Tabs.Trigger class="tabs__trigger" value="nill2">Following</Tabs.Trigger>
        <Tabs.Indicator class="tabs__indicator" />
      </Tabs.List>
      <Tabs.Content class="tabs__content overflow-scroll h-[80vh]" value="posts">
        <div class={cn('flex flex-col gap-4 p-2')}>
          <Show when={props.targetUser()?.follow_status === "following" || props.targetUser()?.profile_privacy === "public"}
            fallback={
              <div class="flex flex-col items-center justify-center h-full">
                <p class="text-2xl font-bold">This user's profile is private</p>
                <p class="text-lg">Follow this user to see their posts</p>
              </div>
            }
          >
            <FeedPosts path={`/profile/${params.username}/posts`} />
          </Show>
        </div>
      </Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill">NOTHING!!!</Tabs.Content>
      <Tabs.Content class="tabs__content" value="nill2">still NOTHING!!!</Tabs.Content>
    </Tabs>
  )
}




