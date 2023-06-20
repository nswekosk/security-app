import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './AuthProvider';
import PrivateRoute from './PrivateRoute';
import Login from './Login';
import FileList from './FileList';

export default function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" Component={Login} />
          <PrivateRoute path="/:path*">
            <FileList />
          </PrivateRoute>
          <PrivateRoute path="/">
            <FileList />
          </PrivateRoute>
        </Routes>  
      </Router>
    </AuthProvider>
  );
};