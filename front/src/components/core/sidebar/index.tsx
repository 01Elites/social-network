import { JSXElement } from "solid-js";
// import SettingsIcon from '@mui/icons-material/Settings';
import SettingsIcon from "@suid/icons-material/Settings";
import AppsIcon from '@suid/icons-material/Apps';
// import PeopleIcon from '@suid/icons-material/People';
// import GroupsIcon from '@suid/icons-material/Groups';
import FlagIcon from '@suid/icons-material/Flag';
import './styles.css'
import Two_Persons_Icon from '../../../components/ui/icons/two_person/two_person';
import Group_Icon from "../../../components/ui/icons/group/group_icon";
import Settings_Icon from "../../../components/ui/icons/settings/settings";

type SidebarProps = {
  children: JSXElement;
};

export default function SideBar(props: SidebarProps) {
  return (
    <div class="grid grid-cols-[max-content_1fr] SideHeight">
      <aside class="slef-start sticky top-0 flex flex-col justify-between ml-9 mt-9">
        {/* This div is for Navigation */}
        <div class="flex flex-col gap-10">
          <AppsIcon class="cursor-pointer" onClick={() => { console.log("apps clicked") }} fontSize="large" />
          {/* <PeopleIcon class="cursor-pointer" onClick={() => { console.log("two person clicked") }} fontSize="large" /> */}
          <Two_Persons_Icon class="w-7 h-7 cursor-pointer" onClick={() => { console.log("two person clicked") }} />
          {/* <GroupsIcon class="cursor-pointer" onClick={() => { console.log("group clicked") }} fontSize="large" /> */}
          <Group_Icon class="cursor-pointer w-8 h-8" onClick={() => { console.log("group clicked") }} />
          <FlagIcon class="cursor-pointer" onClick={() => { console.log("flag clicked") }} fontSize="large" />
        </div>

        {/* This one will have the settings */}
        {/* <SettingsIcon class="cursor-pointer" onClick={() => { console.log("settings clicked") }} fontSize="large" /> */}
        <Settings_Icon class="w-9 h-9 cursor-pointer" onClick={() => { console.log("settings clicked") }} />
      </aside>

      {/* Children will be the contents */}
      <main class="overflow-y-auto w-full text-center justify-self-center">
        {props.children}
      </main>
    </div>
  )
}
