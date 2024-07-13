import { useColorMode } from '@kobalte/core/color-mode';
import { createSignal, JSXElement } from 'solid-js';
import { Label } from '~/components/ui/label';
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from '~/components/ui/sheet';
import { Switch, SwitchControl, SwitchThumb } from '~/components/ui/switch';

/**
 * Opens the settings sheet.
 */
const [settingsOpen, setSettingsOpen] = createSignal(false);

/**
 * Shows the settings sheet.
 */
function showSettings() {
  setSettingsOpen(true);
}

/**
 * Renders the SettingsPage component.
 * It should be added once in the app, and can be controlled using the `showSettings()`.
 * @returns JSXElement representing the SettingsPage component.
 */
function SettingsPage(): JSXElement {
  const { colorMode, setColorMode } = useColorMode();

  return (
    <Sheet open={settingsOpen()} onOpenChange={setSettingsOpen}>
      <SheetContent position='left'>
        <SheetHeader class='mb-4'>
          <SheetTitle>Settings</SheetTitle>
        </SheetHeader>
        <h3 class='text-base font-semibold'>Theme</h3>
        <section>
          <div class='flex items-center justify-between'>
            <Label for='dark-swtich'>Dark Mode</Label>
            <Switch
              checked={colorMode() === 'dark'}
              onChange={(isChecked) => {
                if (isChecked) {
                  setColorMode('dark');
                } else {
                  setColorMode('light');
                }
              }}
            >
              <SwitchControl>
                <SwitchThumb />
              </SwitchControl>
            </Switch>
          </div>
        </section>

        {/* <Separator /> */}
      </SheetContent>
    </Sheet>
  );
}

export { SettingsPage, showSettings };
