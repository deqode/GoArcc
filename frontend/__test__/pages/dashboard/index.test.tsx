import renderer from 'react-test-renderer'
import { UserResponse } from '../../../interface'
import Dashboard, { handler } from '../../../pages/dashboard'

describe('dashboard/index page tests', () => {
  const mockUser: UserResponse = {
    accessToken: 'mockAccessToken',
    idToken: 'mockidToken',
    userId: 'mockuserId',
  }
  it('should match the snapshot', () => {
    const tree = renderer.create(<Dashboard user={mockUser} />).toJSON()
    expect(tree).toMatchSnapshot()
  })

  it('should return to landing page if session not available', async () => {
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
  it('should return user if session is available', async () => {
    const req = {
      session: { get: jest.fn(() => 'MockUser'), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({ props: { user: 'MockUser' } })
  })

  it('should show circular progress if url not set yet', async () => null)

  it('should show "Connect with github" button if url set', async () => null)

  it('should redirect to github connect page on "Connect with github" button click', async () =>
    null)
})
