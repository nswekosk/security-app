import { render } from "@testing-library/react";
import App from "./App";
import { AuthProvider } from './AuthProvider';
import { ReactNode } from 'react';
/*
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  BrowserRouter: ({ children }: { children?: ReactNode }) => <div>{children}</div>,
  Routes: ({ children }: { children?: ReactNode }) => <div>{children}</div>,
  Route: ({ children }: { children?: ReactNode }) => <div>{children}</div>,
  useNavigate: () => jest.fn(),
}));

jest.mock('./AuthProvider', () => ({
  __esModule: true,
  AuthProvider: ({children}: { children?: ReactNode }) => children,
}));

test("renders without crashing", () => {
  render(<App />);
});*/