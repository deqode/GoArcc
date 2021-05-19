import renderer from 'react-test-renderer'
import { UserResponse } from '../../../interface'
import Dashboard, { handler } from '../../../pages/dashboard'

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

  it('should return to /', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({
      redirect: {
        permanent: false,
        destination: '/',
      },
    })
  })
  it('should return user', async () => {
    const req = {
      session: { get: jest.fn(() => 'MockUser'), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({ props: { user: 'MockUser' } })
  })
})
