import { Button } from "~/components/ui/button";
import Follow_Icon from '~/components/ui/icons/follow_icon';
import { fetchWithAuth } from '~/extensions/fetch';
import config from '~/config';
import { JSXElement, Show } from 'solid-js';
import { createSignal } from 'solid-js'

type FollowRequestParams = {
  username: string | undefined;
  status: string;
  privacy: string;
}
export default function FollowRequest(props: FollowRequestParams): JSXElement {
  console.log(props.username, props.status)
  var [buttonData, setButtonData] = createSignal("");
  if (props.status == "following") {
    setButtonData("Unfollow")
  } else if (props.status == "pending") {
    setButtonData("Cancel Follow Request")
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
            setButtonData("Cancel Follow Request")
          } else {
            setButtonData("Unfollow")
          }
        } else if (buttonData() === "Unfollow") {
          setButtonData("Follow")
        } else if (buttonData() === "Cancel Follow Request") {
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
    <Button class="flex grow" variant="default" onClick={sendRequestApi}>
      <Show when={props.status == "not followed"}>
        <Follow_Icon />
      </Show>{buttonData()}
    </Button>
  </>)
}
