import { render, fireEvent } from '@testing-library/react'
import renderer from 'react-test-renderer'

import * as api from '../../api/rest/session'
import { User } from '../../contexts/UserContext'
import Navbar from '../Navbar'

it('should match the snapshot of logged-in user', () => {
  const mockUser: User = {
    idToken: 'mockidToken',
    loggedIn: true,
  }
  const tree = renderer.create(<Navbar user={mockUser} />).toJSON()
  expect(tree).toMatchSnapshot()
})

it('should match the snapshot of logged-out user', () => {
  const mockUser: User = {
    idToken: 'mockidToken',
    loggedIn: false,
  }
  const tree = renderer.create(<Navbar user={mockUser} />).toJSON()
  expect(tree).toMatchSnapshot()
})

it('should click logout ', async () => {
  const mockUser: User = {
    idToken: 'mockidToken',
    loggedIn: true,
  }
  const { getByText } = render(<Navbar user={mockUser} />)
  fireEvent.click(getByText('common:LOGOUT'))
})

it('should get logged out form session  ', async () => {
  api.destroyUserSession = jest.fn().mockReturnValue(() => Promise.resolve({ error: false }))
  const mockUser: User = {
    idToken: 'mockidToken',
    loggedIn: true,
  }
  const { getByText } = render(<Navbar user={mockUser} />)
  fireEvent.click(getByText('common:LOGOUT'))
})
