import { JSXElement } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import logo from '../../../logo.svg';
import { TextField, TextFieldInput } from '../../../components/ui/text-field';
import Logout from '../../../lib/logout';

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
    <header class='mx-5 mt-5 flex gap-10 align-middle justify-between'>
      <img
        src={logo}
        alt='Elite Logo'
        onClick={() => {
          navigate('/');
        }}
        class='cursor-pointer w-20'
      />
      <TextField class='basis-1/2'>
        <TextFieldInput type='search' placeholder='Search anything...' />
      </TextField>

      <TextField
        class='cursor-pointer'
        onClick={Logout}>
        Logout
      </TextField>
    </header>
  );
}
