import { createSignal, JSXElement } from 'solid-js';
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from '~/components/ui/sheet';
import NotificationsFeed from './notificationsfeed';

/**
 * Opens the settings sheet.
 */
const [settingsOpen, setSettingsOpen] = createSignal(false);

/**
 * Shows the settings sheet.
 */
function showNotifications() {
  setSettingsOpen(true);
}

/**
 * Renders the SettingsPage component.
 * It should be added once in the app, and can be controlled using the `showSettings()`.
 * @returns JSXElement representing the SettingsPage component.
 */
function NotificationsPage(): JSXElement {

  return (
    <Sheet open={settingsOpen()} onOpenChange={setSettingsOpen}>
      <SheetContent position='left'>
        <SheetHeader class='mb-4'>
          <SheetTitle>Notifications</SheetTitle>
        </SheetHeader>
        <h3 class='text-base font-semibold'>Theme</h3>
        <section>
          <div class='flex items-center justify-between'>
            <NotificationsFeed/>
          </div>
        </section>
        {/* <Separator /> */}
      </SheetContent>
    </Sheet>
  );
}

export { showNotifications, NotificationsPage };
