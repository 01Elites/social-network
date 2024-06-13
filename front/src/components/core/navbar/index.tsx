import { JSXElement } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import logo from '../../../logo.svg';
import { TextField, TextFieldInput } from '../../../components/ui/text-field';

type NavbarProps = {
  variant?: 'loggedin' | 'loggedout';
};

export default function Navbar(prop: NavbarProps): JSXElement {
  // Default variant is 'loggedin'
  if (!prop.variant) {
    prop.variant = 'loggedin';
  }

  const navigate = useNavigate();

  return (
    <header class='mx-5 mt-5 flex gap-10 align-middle'>
      <img
        src={logo}
        alt='Elite Logo'
        onClick={() => {
          navigate('/');
        }}
        class='cursor-pointer'
      />
      <TextField>
        <TextFieldInput type='search' placeholder='Search anything...' />
      </TextField>
    </header>
  );
}
