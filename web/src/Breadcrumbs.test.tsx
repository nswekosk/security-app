import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Breadcrumbs from './Breadcrumbs';

test('renders breadcrumbs correctly', () => {
  const path = "/folder1/folder2/file";

  render(
    <MemoryRouter>
      <Breadcrumbs path={path} />
    </MemoryRouter>
  );

  // Check if 'Root' link exists
  expect(screen.getByText('Root')).toBeInTheDocument();

  // Check if each part of the path is rendered as a link
  const pathParts = path.split('/').filter(Boolean);
  pathParts.forEach(part => {
    expect(screen.getByText(part)).toBeInTheDocument();
  });
});
