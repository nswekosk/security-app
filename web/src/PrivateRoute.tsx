import React, { useContext } from 'react';
import { Route, Navigate } from 'react-router-dom';
import AuthContext from './AuthContext';

interface PrivateRouteProps {
  children: React.ReactNode;
  path: string;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ children, ...rest }) => {
  const authContext = useContext(AuthContext);

  if (!authContext) {
    throw new Error("PrivateRoute must be used inside an AuthProvider");
  }

  const { token } = authContext;

  return (
    <Route {...rest}>
      {token ? children : <Navigate to="/login" />}
    </Route>
  );
};

export default PrivateRoute;