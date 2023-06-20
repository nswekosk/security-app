import React from 'react';
import { render, screen } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import '@testing-library/jest-dom';
import FileList from './FileList';
import { AuthProvider } from './AuthProvider';
/*
// A mock for the useParams hook
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useParams: () => ({
    path: '',
  }),
}));

test('displays loading state', () => {
  render(
    <AuthProvider>
      <Router>
        <FileList />
      </Router>
    </AuthProvider>
  );

  expect(screen.getByText(/loading.../i)).toBeInTheDocument();
});*/