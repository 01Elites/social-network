import { JSXElement } from 'solid-js';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Settings_Icon(props?: inputs): JSXElement {
  return (
    <div class={props?.class} onClick={props?.onClick}>
      <svg style={{ width: "100%", height: "100%" }} viewBox="0 0 100 100" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M60.1 88.7C60 88.7 59.8 88.7 59.7 88.7C57.6 88.3 55.7 86.9 54.7 85C54.6 84.9 54.6 84.8 54.5 84.7C54.4 84.5 53.6 83 51.3 82.9H51.2C50.4 82.9 49.7 82.9 48.9 82.9H48.8C46.4 82.9 45.6 84.6 45.6 84.7C45.5 84.8 45.5 84.9 45.4 85C44.4 86.9 42.5 88.3 40.4 88.7C40.1 88.8 39.8 88.8 39.5 88.7C37.6 88.2 35.6 87.5 33.8 86.6C33.5 86.5 33.3 86.3 33.1 86.1C31.8 84.4 31.2 82.1 31.7 80C31.7 79.9 31.7 79.8 31.8 79.6C31.8 79.4 32.2 77.7 30.5 76.2L30.4 76.1C29.8 75.6 29.2 75.1 28.6 74.6C27.8 74 27 73.7 26.1 73.7C25.4 73.7 24.9 73.9 24.9 73.9C24.8 74 24.7 74 24.5 74C22.5 74.8 20.1 74.7 18.2 73.6C17.9 73.5 17.7 73.3 17.6 73C16.5 71.4 15.5 69.7 14.6 67.9C14.5 67.6 14.4 67.4 14.4 67.1C14.4 64.9 15.5 62.7 17.2 61.4C17.3 61.3 17.4 61.2 17.5 61.2C17.6 61.1 19 60 18.7 57.7C18.7 57.7 18.7 57.7 18.7 57.6C18.5 56.9 18.4 56.1 18.3 55.3C18.3 55.2 18.3 55.2 18.3 55.2C17.9 52.9 16 52.4 16 52.4C15.9 52.4 15.8 52.3 15.6 52.3C13.5 51.6 11.7 49.9 11 47.8C10.9 47.5 10.9 47.3 10.9 47C11.1 45.2 11.4 43.3 11.9 41.4C12 41.1 12.1 40.9 12.3 40.7C13.8 38.9 16 37.9 18.2 38C18.3 38 18.4 38 18.6 38C19.1 38 20.7 37.9 21.7 36.1C21.7 36.1 21.7 36 21.8 36C22.1 35.4 22.5 34.7 23 34C23 34 23 33.9 23.1 33.9C24.2 31.9 23.3 30.4 23.2 30.2C23.1 30.1 23.1 30 23 29.9C21.8 27.9 21.5 25.5 22.4 23.3C22.5 23.1 22.6 22.8 22.8 22.6C24.1 21.3 25.5 20.1 27 19.1C27.2 18.9 27.5 18.8 27.7 18.8C30 18.3 32.4 19 34.1 20.5C34.2 20.6 34.3 20.6 34.4 20.7C34.4 20.7 35.3 21.5 36.6 21.5C37 21.5 37.5 21.4 38 21.3C38.8 21 39.6 20.7 40.3 20.5C42.4 19.7 42.7 17.9 42.7 17.7C42.7 17.6 42.7 17.5 42.8 17.4C43.1 15.1 44.5 13 46.6 11.9C46.8 11.8 47.1 11.7 47.3 11.7C49.3 11.6 50.7 11.6 52.7 11.7C53 11.7 53.2 11.8 53.4 11.9C55.5 13 56.9 15 57.2 17.3C57.2 17.4 57.3 17.5 57.3 17.7C57.3 17.9 57.6 19.6 59.7 20.5C60.5 20.7 61.3 21 62 21.3C62.5 21.5 63 21.6 63.4 21.6C64.8 21.6 65.6 20.8 65.6 20.8C65.7 20.7 65.8 20.6 65.9 20.6C67.6 19.1 70 18.4 72.3 18.9C72.6 19 72.8 19.1 73 19.2C74.5 20.3 75.9 21.5 77.2 22.7C77.4 22.9 77.5 23.1 77.6 23.4C78.4 25.6 78.2 28 77 30C77 30.1 76.9 30.2 76.8 30.3C76.7 30.4 75.8 32 76.9 34C76.9 34 76.9 34.1 77 34.1C77.4 34.8 77.8 35.4 78.2 36.1C78.2 36.1 78.2 36.2 78.3 36.2C79.3 37.9 80.9 38.1 81.4 38.1C81.5 38.1 81.7 38.1 81.8 38.1C84 38 86.3 39 87.7 40.8C87.9 41 88 41.3 88.1 41.5C88.6 43.3 88.9 45.2 89.1 47.1C89.1 47.4 89.1 47.6 89 47.9C88.2 50.1 86.5 51.7 84.4 52.4C84.3 52.5 84.2 52.5 84 52.5C83.8 52.5 82.1 53.1 81.7 55.4V55.5C81.6 56.3 81.5 57.1 81.3 57.8V57.9C81 59.9 82.5 61 82.5 61C82.6 61.1 82.7 61.2 82.8 61.3C84.5 62.7 85.6 64.8 85.6 67C85.6 67.3 85.5 67.6 85.4 67.8C84.6 69.5 83.6 71.3 82.4 72.9C82.2 73.1 82 73.3 81.8 73.5C80.7 74.1 79.5 74.4 78.2 74.4C77.3 74.4 76.3 74.2 75.5 73.9C75.4 73.9 75.3 73.8 75.1 73.8C75.1 73.8 74.6 73.6 73.9 73.6C73 73.6 72.2 73.9 71.5 74.5C71.5 74.5 71.5 74.5 71.4 74.5C70.9 75 70.3 75.5 69.6 76C69.6 76 69.5 76 69.5 76.1C68 77.4 68 78.8 68.1 79.4C68.2 79.6 68.2 79.8 68.2 80C68.6 82.1 68.1 84.3 66.7 85.9C66.5 86.1 66.3 86.3 66 86.4C64.2 87.2 62.2 87.9 60.3 88.5C60.4 88.7 60.2 88.7 60.1 88.7ZM35.7 83.4C37.1 84 38.5 84.5 39.9 84.9C40.9 84.6 41.7 83.9 42.1 83C42.2 82.9 42.2 82.8 42.3 82.6C43.1 81.2 45.1 79.2 48.5 79.1C48.6 79.1 48.6 79.1 48.6 79.1H48.7C49.5 79.1 50.2 79.1 51 79.1H51.1H51.2C54.5 79.2 56.6 81.2 57.4 82.6C57.5 82.7 57.6 82.8 57.6 83C58 83.9 58.8 84.6 59.8 84.9C61.2 84.5 62.7 84 64 83.4C64.5 82.5 64.7 81.5 64.4 80.5C64.4 80.4 64.3 80.2 64.3 80.1C64 78.5 64.3 75.7 66.8 73.4L66.9 73.3L67 73.2C67.7 72.7 68.2 72.2 68.8 71.7L68.9 71.6C68.9 71.6 69 71.6 69 71.5C70.4 70.4 72 69.8 73.7 69.8C74.7 69.8 75.5 70 76 70.2C76.1 70.2 76.3 70.3 76.4 70.3C76.9 70.5 77.4 70.6 77.9 70.6C78.4 70.6 78.9 70.5 79.3 70.3C80.1 69.1 80.9 67.8 81.5 66.5C81.4 65.5 80.8 64.5 80 63.9C79.9 63.8 79.8 63.7 79.7 63.6C78.4 62.5 76.8 60.2 77.3 56.9V56.8V56.7C77.5 56 77.6 55.2 77.7 54.4C77.7 54.3 77.7 54.3 77.7 54.2V54.1C78.4 50.8 80.7 49.2 82.2 48.6C82.3 48.5 82.4 48.5 82.6 48.4C83.6 48.1 84.5 47.4 85 46.4C84.8 45 84.6 43.6 84.3 42.3C83.5 41.5 82.5 41.1 81.4 41.2C81.2 41.2 81.1 41.2 80.9 41.2C79.8 41.2 76.8 40.8 74.8 37.6L74.7 37.5C74.3 36.7 73.9 36.1 73.5 35.4L73.4 35.3C73.4 35.3 73.4 35.2 73.3 35.2C71.7 32.2 72.4 29.5 73.2 28.1C73.3 28 73.3 27.9 73.4 27.7C74 26.8 74.2 25.7 73.9 24.6C72.9 23.7 71.9 22.8 70.8 22C69.7 21.9 68.6 22.3 67.8 23C67.7 23.1 67.6 23.2 67.4 23.3C67.3 24 65.8 25 63.5 25C62.7 25 61.8 24.9 61 24.6H60.9H60.8C60.1 24.3 59.3 24 58.6 23.8H58.5C58.5 23.8 58.5 23.8 58.4 23.8C55.3 22.6 54 20 53.7 18.4C53.7 18.3 53.6 18.1 53.6 18C53.5 16.9 52.9 15.9 52 15.3C50.5 15.2 49.5 15.2 48 15.3C47.1 15.9 46.5 16.9 46.4 18C46.4 18.1 46.4 18.3 46.3 18.4C46 20 44.8 22.6 41.6 23.8C41.6 23.8 41.6 23.8 41.5 23.8H41.4C40.7 24 39.9 24.3 39.2 24.6H39.1H39C38.2 24.9 37.4 25 36.5 25C34.2 25 32.7 24 32 23.4C31.9 23.3 31.8 23.2 31.6 23.1C30.8 22.3 29.7 21.9 28.6 22.1C27.5 22.9 26.5 23.8 25.5 24.7C25.2 25.8 25.4 26.9 26 27.8C26.1 27.9 26.1 28 26.2 28.2C27 29.6 27.7 32.4 26.1 35.3V35.4L26 35.5C25.6 36.2 25.2 36.8 24.8 37.5C24.8 37.6 24.7 37.6 24.7 37.7C22.7 40.9 19.7 41.3 18.6 41.3C18.5 41.3 18.3 41.3 18.2 41.3C17.1 41.2 16.1 41.6 15.3 42.4C15 43.8 14.7 45.2 14.6 46.5C15.1 47.5 15.9 48.2 17 48.5C17.1 48.5 17.3 48.6 17.4 48.7C18.9 49.3 21.2 50.9 21.9 54.2V54.3C21.9 54.4 21.9 54.4 21.9 54.5C22 55.3 22.1 56 22.3 56.8V56.9V57C22.8 60.3 21.2 62.6 19.9 63.7C19.8 63.8 19.7 63.9 19.6 64C18.7 64.6 18.2 65.6 18.1 66.6C18.8 67.9 19.5 69.2 20.3 70.4C21.2 70.8 22.3 70.8 23.3 70.3C23.4 70.2 23.6 70.2 23.7 70.2C24.2 70 25 69.8 26 69.8C27.7 69.8 29.3 70.4 30.7 71.5C30.7 71.5 30.7 71.5 30.8 71.5C30.9 71.5 30.9 71.6 31 71.6C31.5 72.1 32.1 72.6 32.8 73.1L32.9 73.2L33 73.3C35.5 75.5 35.8 78.4 35.5 80C35.5 80.1 35.5 80.3 35.4 80.4C35 81.5 35.2 82.5 35.7 83.4ZM50 70.2C39.2 70.2 30.5 61.4 30.5 50.7C30.5 40 39.2 31.1 50 31.1C60.8 31.1 69.5 39.9 69.5 50.6C69.5 61.3 60.8 70.2 50 70.2ZM50 34.8C41.3 34.8 34.1 41.9 34.1 50.7C34.1 59.5 41.3 66.5 50 66.5C58.7 66.5 65.9 59.4 65.9 50.6C65.9 41.8 58.7 34.8 50 34.8Z" fill="hsl(var(--primary))" />
      </svg>
    </div>
  );
}
