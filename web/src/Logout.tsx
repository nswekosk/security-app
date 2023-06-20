import React, { useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import AuthContext from './AuthContext';

const Logout: React.FC = () => {
  const navigate = useNavigate();
  const authContext = useContext(AuthContext);

  if (!authContext) {
    throw new Error("Logout must be used inside an AuthProvider");
  }

  const { logOut } = authContext;

  const handleLogout = () => {
    logOut();
    navigate('/login');
  };

  return (
    <button onClick={handleLogout}>
      Log Out
    </button>
  );
};

export default Logout;