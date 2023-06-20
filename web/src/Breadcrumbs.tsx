import React from 'react';
import { Link } from 'react-router-dom';

type BreadcrumbsProps = {
  path: string;
};

const Breadcrumbs: React.FC<BreadcrumbsProps> = ({ path }) => {
  // Split the path into its components
  const pathParts = path.split('/').filter(Boolean);

  // This will be our list of breadcrumb links
  let breadcrumbLinks = [];

  // Iterate over the path parts and create a link for each one
  for (let i = 0; i < pathParts.length; i++) {
    // The link's path is the joined array of path parts up to and including this one
    const linkPath = '/' + pathParts.slice(0, i + 1).join('/');

    breadcrumbLinks.push(
      <span key={i}>
        <Link to={linkPath}>{pathParts[i]}</Link>
        {i < pathParts.length - 1 && ' / '}
      </span>
    );
  }

  return (
    <div>
      <Link to="/">Root</Link>
      {breadcrumbLinks.length > 0 && ' / '}
      {breadcrumbLinks}
    </div>
  );
};

export default Breadcrumbs;