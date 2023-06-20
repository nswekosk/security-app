import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react';
import { MemoryRouter as Router, Route } from 'react-router-dom';
import { AuthProvider } from './AuthProvider';
import Logout from './Logout';
/*
describe('Logout', () => {
  it('clears auth state and redirects', async () => {
    const utils = render(
      <AuthProvider>
        <Router>
          <Route path="/">
            <Logout />
          </Route>
          <Route path="/login">
            <div>Login page</div>
          </Route>
        </Router>
      </AuthProvider>
    );

    const buttonLogout = utils.getByText('Logout');
    fireEvent.click(buttonLogout);

    await waitFor(() => {
      expect(utils.getByText('Login page')).toBeInTheDocument();
    });
  });
});*/