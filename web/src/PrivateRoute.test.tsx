import React from 'react';
import { render } from '@testing-library/react';
import { MemoryRouter as Router, Route } from 'react-router-dom';
import { AuthProvider } from './AuthProvider';
import PrivateRoute from './PrivateRoute';
/*
describe('PrivateRoute', () => {
  it('redirects unauthenticated users to login', () => {
    const utils = render(
      <AuthProvider>
        <Router initialEntries={['/private']}>
          <PrivateRoute path="/private">
            <div>Private page</div>
          </PrivateRoute>
          <Route path="/login">
            <div>Login page</div>
          </Route>
        </Router>
      </AuthProvider>
    );

    expect(utils.getByText('Login page')).toBeInTheDocument();
  });
});*/