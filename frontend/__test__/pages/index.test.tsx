import renderer from 'react-test-renderer'

import * as api from '../../api/rest/fetchUrls'
import Landing, { handler } from '../../pages'
import { redirectToErrorPage } from '../../utils/redirects'

// jest.mock('../../api/rest/fetchUrls', () => {
//   return {
// getLoginURL: jest.fn().mockImplementation(() =>
//   Promise.resolve({
//     url: 'mockUrl',
//     error: false,
//   })
//     ),
//   }
// })

describe('Landing page tests', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Landing url={'moc url'} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
  it('should get login url if user session not available', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
    api.getLoginURL = jest.fn().mockImplementation(() =>
      Promise.resolve({
        url: 'mockUrl',
        error: false,
      })
    )
    const response = await handler({ req })
    expect(response).toStrictEqual({ props: { url: 'mockUrl' } })
  })
  it('should redirect to dashboard if user session available', async () => {
    const req = {
      session: { get: jest.fn().mockImplementation(() => true), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({ redirect: { permanent: false, destination: '/dashboard' } })
  })

  it('should redirect to error page if network error', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
    api.getLoginURL = jest.fn().mockImplementation(() =>
      Promise.resolve({
        url: '',
        error: true,
      })
    )
    const response = await handler({ req })
    expect(response).toStrictEqual(redirectToErrorPage('Network Error'))
  })

  it('should redirect to github sign up page on Sing Up button click', async () => null)
})
