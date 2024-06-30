import { Component, JSX } from 'solid-js';
import { Router, Route, useNavigate } from '@solidjs/router';

function isAuthenticated(): boolean {
  // Replace with your actual authentication logic
  return false;
}

interface AuthGuardProps {
  children: JSX.Element;
}

function AuthGuard(props: AuthGuardProps): JSX.Element | null {
  const navigate = useNavigate();

  if (!isAuthenticated()) {
    navigate('/login');
    return null;
  }

  return props.children;
}

export default AuthGuard;
