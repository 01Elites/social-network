import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement, Show, useContext } from 'solid-js';
import { createSignal } from 'solid-js';
import UserDetailsContext from '~/contexts/UserDetailsContext';
import User, { UserDetailsHook } from '~/types/User';
import Unfollow_Icon from "~/components/ui/icons/unfollow_icon";

type FollowRequestParams = {
  username: string | undefined;
  status: string;
  privacy: string;
}
export default function FollowRequest(props: FollowRequestParams): JSXElement {
  const { userDetails } = useContext(UserDetailsContext) as UserDetailsHook;

  var [buttonData, setButtonData] = createSignal("");
  if (props.status == "following") {
    setButtonData("Unfollow")
  } else if (props.status == "pending") {
    setButtonData("Requested")
  } else if (props.status == "not_following") {
    setButtonData("Follow")
  }
  function sendRequestApi() {
    console.log(buttonData())
    if (buttonData() === "") {
      return
    }
    fetchWithAuth(config.API_URL + "/follow_request", {
      method: 'POST',
      body: JSON.stringify({
        receiver: props.username
      })
    }).then(async (res) => {
      if (res.status === 200) {
        if (buttonData() === "Follow") {
          if (props.privacy === "private") {
            setButtonData("Requested")
          } else {
            setButtonData("Unfollow")
          }
        } else if (buttonData() === "Unfollow") {
          setButtonData("Follow")
        } else if (buttonData() === "Requested") {
          setButtonData("Follow")
        }
        return;
      } else {
        console.log(res.body);
        console.log('Error making request');
        return
      }
    })
  }
  return (<>
    <Show when={userDetails()?.user_name != props.username}>
      <Button class="flex flex-row grow" variant="default" onClick={sendRequestApi}>
        <Show when={props.status === "not_following"} fallback={<Unfollow_Icon />}>
        <Follow_Icon />
        </Show>{buttonData()}
      </Button>
    </Show>
  </>)
}
