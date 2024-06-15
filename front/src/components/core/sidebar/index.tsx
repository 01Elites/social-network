import { JSXElement } from "solid-js";
// import SettingsIcon from '@mui/icons-material/Settings';
import Home from "@suid/icons-material/Home";
import SettingsIcon from "@suid/icons-material/Settings";
import AppsIcon from '@suid/icons-material/Apps';
import PeopleIcon from '@suid/icons-material/People';
import GroupsIcon from '@suid/icons-material/Groups';
import PersonIcon from '@suid/icons-material/Person';
import './styles.css'

type SidebarProps = {
  children: JSXElement;
};

export default function SideBar(props: SidebarProps) {
  return (
    <div class="grid grid-cols-[max-content_1fr] SideHeight">
      <aside class="slef-start sticky top-0 flex flex-col justify-between ml-9 mt-9">
        {/* This div is for Navigation */}
        <div class="flex flex-col gap-10">
          <Home fontSize="large" />
          <AppsIcon fontSize="large" />
          <PersonIcon fontSize="large" />
          <PeopleIcon fontSize="large" />
          <GroupsIcon fontSize="large" />
        </div>

        {/* This one will have the settings */}
        <SettingsIcon fontSize="large" />
      </aside>

      {/* Children will be the contents */}
      <main class="overflow-y-auto w-full text-center justify-self-center">
        {props.children}
      </main>
    </div>
  )
}
