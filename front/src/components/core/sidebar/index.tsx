import { JSXElement } from 'solid-js';
// import SettingsIcon from "@suid/icons-material/Settings";
// import AppsIcon from '@suid/icons-material/Apps';
// import PeopleIcon from '@suid/icons-material/People';
// import GroupsIcon from '@suid/icons-material/Groups';
// import FlagIcon from '@suid/icons-material/Flag';
import './styles.css';
import Two_Persons_Icon from '../../../components/ui/icons/two_person/two_person';
import Group_Icon from '../../../components/ui/icons/group/group_icon';
import Settings_Icon from '../../../components/ui/icons/settings/settings';
import Apps_Icon from '../../../components/ui/icons/apps/apps_icon';
import Flag_Icon from '../../../components/ui/icons/flag/flag_icon';

type SidebarProps = {
  // children: JSXElement;
};

export default function SideBar(props: SidebarProps) {
  return (
    <div class='grid grid-cols-[max-content_1fr]'>
      <aside class='slef-start sticky top-0 flex flex-col justify-between p-5'>
        {/* This div is for Navigation */}
        <div class='flex flex-col gap-10'>
          <Apps_Icon
            class='h-8 w-8 cursor-pointer self-center'
            onClick={() => {
              console.log('apps clicked');
            }}
          />
          <Two_Persons_Icon
            class='h-7 w-7 cursor-pointer self-center'
            onClick={() => {
              console.log('two person clicked');
            }}
          />
          <Group_Icon
            class='h-8 w-8 cursor-pointer self-center'
            onClick={() => {
              console.log('group clicked');
            }}
          />
          <Flag_Icon
            class='flex h-8 w-8 cursor-pointer self-center'
            onClick={() => {
              console.log('flag clicked');
            }}
          />
        </div>

        {/* This one will have the settings */}
        <Settings_Icon
          class='h-9 w-9 cursor-pointer self-center'
          onClick={() => {
            console.log('settings clicked');
          }}
        />
      </aside>
    </div>
  );
}
