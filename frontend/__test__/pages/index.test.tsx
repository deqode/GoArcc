import renderer from 'react-test-renderer'
import Landing, { handler } from '../../pages'

jest.mock('../../api/rest/fetchUrls', () => {
  return {
    getLoginURL: jest.fn().mockImplementation(() =>
      Promise.resolve({
        url: 'mockUrl',
        error: false,
      })
    ),
  }
})

describe('Check Landing page snapshot', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Landing url={'moc url'} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
  it('should match url', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({ props: { url: 'mockUrl' } })
  })
  it('should not have session', async () => {
    const req = {
      session: { get: jest.fn().mockImplementation(() => true), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({ redirect: { permanent: false, destination: '/dashboard' } })
  })
})
