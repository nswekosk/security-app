import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import jwt_decode from 'jwt-decode';
import { AuthProvider } from './AuthProvider';
import AuthContext from './AuthContext';

jest.mock('jwt-decode');

// Mock component to simulate child components within AuthProvider
const MockComponent = () => {
  const authContext = React.useContext(AuthContext);
  
  if (!authContext) {
    return null;
  }

  const { token, logIn, logOut } = authContext;

  return (
    <div>
      <div data-testid="token-value">{token}</div>
      <button onClick={() => logIn('mockToken')}>Log in</button>
      <button onClick={logOut}>Log out</button>
    </div>
  );
};

// Test for AuthProvider
test('AuthProvider manages token', () => {
  // Mock jwt_decode to return a decoded token with a far future expiration time
  (jwt_decode as jest.MockedFunction<typeof jwt_decode>).mockReturnValue({
    exp: Math.floor(Date.now() / 1000) + 10000,
  });

  const { getByText, getByTestId } = render(
    <MemoryRouter>
      <AuthProvider>
        <MockComponent />
      </AuthProvider>
    </MemoryRouter>
  );

  // Test initial state (no token)
  expect(getByTestId('token-value').textContent).toBe('');

  // Simulate log in
  fireEvent.click(getByText('Log in'));

  // Test that token is stored
  expect(getByTestId('token-value').textContent).toBe('mockToken');

  // Simulate log out
  fireEvent.click(getByText('Log out'));

  // Test that token is cleared
  expect(getByTestId('token-value').textContent).toBe('');
});
