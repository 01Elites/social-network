// type `Notification` is already defined in the global scope, we prefix with SN (Social-Network) to avoid conflicts
export type SNNotification = {
  type: string;
  message: string;
  notification_id: string;
  read: boolean;
  /**
   * Different for each type of notification
   */
  metadata: any;

  /**
   *  Wrapper to ease the process of marking a notification as read
   */
  // markAsRead: Promise<Error | null>;
};
