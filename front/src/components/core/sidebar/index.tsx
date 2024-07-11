import { showSettings } from '~/pages/settings';
import IconApps from '../../ui/icons/IconApps';
import IconFlag from '../../ui/icons/IconFlag';
import IconGroup from '../../ui/icons/IconGroup';
import IconSettings from '../../ui/icons/IconSettings';
import IconTwoPerson from '../../ui/icons/IconTwoPerson';

type SidebarProps = {
  // children: JSXElement;
};

export default function SideBar(props: SidebarProps) {
  return (
    <div class='flex h-full flex-col'>
      {/* This div is for Navigation */}
      <div class='flex h-full flex-col justify-center gap-10'>
        <IconApps
          class='h-8 w-8 cursor-pointer self-center'
          onClick={() => {
            console.log('apps clicked');
          }}
        />
        <IconTwoPerson
          class='h-7 w-7 cursor-pointer self-center'
          onClick={() => {
            console.log('two person clicked');
          }}
        />
        <IconGroup
          class='h-8 w-8 cursor-pointer self-center'
          onClick={() => {
            console.log('group clicked');
          }}
        />
        <IconFlag
          class='flex h-8 w-8 cursor-pointer self-center'
          onClick={() => {
            console.log('flag clicked');
          }}
        />
      </div>
      {/* This one will have the settings */}
      <div class='flex flex-col justify-end'>
        <IconSettings
          class='h-9 w-9 cursor-pointer self-center'
          onClick={() => {
            showSettings();
          }}
        />
      </div>
    </div>
  );
}
