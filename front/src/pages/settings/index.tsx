import { createSignal, JSXElement } from 'solid-js';
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from '~/components/ui/sheet';

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
  return (
    <Sheet open={settingsOpen()} onOpenChange={setSettingsOpen}>
      <SheetContent position='left'>
        <SheetHeader>
          <SheetTitle>Settings</SheetTitle>
        </SheetHeader>
      </SheetContent>
    </Sheet>
  );
}

export { SettingsPage, showSettings };
