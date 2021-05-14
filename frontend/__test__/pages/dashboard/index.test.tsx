import renderer from 'react-test-renderer'
import { UserResponse } from '../../../interface'
import Dashboard from '../../../pages/dashboard'

describe('Check Dashboard indexPage page snapshot', () => {
  const mockUser: UserResponse = {
    accessToken: 'mockAccessToken',
    idToken: 'mockidToken',
    userId: 'mockuserId',
  }
  it('should match the snapshot', () => {
    const tree = renderer.create(<Dashboard user={mockUser} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})
