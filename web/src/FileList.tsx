import React, { useState, useEffect, useContext } from 'react';
import {
  Link,
  useParams,
  useNavigate
} from 'react-router-dom';
import Breadcrumbs from './Breadcrumbs';
import AuthContext from './AuthContext';
import Logout from './Logout';

type FileInfo = {
  name: string;
  type: string;
  size: string;
};

type FileResponse = {
    name: string;
    type: string;
    size: string;
    contents: FileInfo[];
};
  
const FileList: React.FC = () => {
    const params = useParams();
    const mainPath = params["path"] || "";
    const extendedPath = params['*'] || "";
    const path = extendedPath ? (mainPath + "/" + extendedPath) : mainPath;
    const [fileData, setFileData] = useState<FileResponse | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [filter, setFilter] = useState<string>('');
    const navigate = useNavigate();
    const authContext = useContext(AuthContext);
    
    if (!authContext) {
      throw new Error("FileList must be used inside an AuthProvider");
    }
    const token = authContext.token || "";

    useEffect(() => {
      const getData = async () => {
          const response = await fetch(`https://localhost:8080/files?path=${path}`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          if (!response.ok) {
            if (response.status === 401) {
              // If unauthorized, log the user out and redirect to login
              authContext.logOut();
              navigate('/login');
            } else {
              // Otherwise, set error message
              const errorData = await response.json();
              setError(errorData.message || 'An error occurred');
            }
            return;
          }
          const data = await response.json() as FileResponse;
          setFileData(data);
      };
  
      getData();
    }, [path, token, authContext, navigate]);
  
    if (!fileData) {
      return <div>Loading...</div>;
    }

    if (error) {
      return <div>Error: {error}</div>;
    }
  
    var contents = (fileData && fileData.contents) || [];
    const filteredFiles = contents.filter((file) =>
      file.name.includes(filter)
    );
  
    return (
      <div>
        <Logout />
        <Breadcrumbs path={path} />
        <h2>Contents in: /{path}</h2>
        <input
          type="text"
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
          placeholder="Filter files"
        />
        {filteredFiles.map((file, index) => (
          <div key={index}>
            {file.type === 'directory' ? (
              <Link to={`/${path}${file.name}/`}>{file.name}</Link>
            ) : (
              <h3>{file.name}</h3>
            )}
            <p>Type: {file.type}</p>
            <p>Size: {file.size}</p>
          </div>
        ))}
      </div>
    );
};

export default FileList;