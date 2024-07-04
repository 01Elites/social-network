import group from './group.svg';

interface inputs {
  class?: string;
  onClick?: () => void;
}

export default function Group_Icon(props: inputs) {
  return (
    <img
      src={group}
      alt='Group'
      class={props ? props.class : ''}
      onClick={props ? props.onClick : () => {}}
    />
  );
}
