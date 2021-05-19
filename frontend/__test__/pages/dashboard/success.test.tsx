import renderer from 'react-test-renderer'
import Success, { handler } from '../../../pages/dashboard/success'

describe('dashboard/success page tests', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Success />).toJSON()
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
})
