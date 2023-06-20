import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import Login from './Login';

test('renders and updates inputs correctly', () => {
  const { getByPlaceholderText } = render(<Login />);
  
  const usernameInput = getByPlaceholderText('Username') as HTMLInputElement;
  const passwordInput = getByPlaceholderText('Password') as HTMLInputElement;
  
  fireEvent.change(usernameInput, { target: { value: 'testuser' } });
  fireEvent.change(passwordInput, { target: { value: 'testpass' } });
  
  expect(usernameInput.value).toBe('testuser');
  expect(passwordInput.value).toBe('testpass');
});