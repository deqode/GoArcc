import renderer from 'react-test-renderer'
import Landing, { handler } from '../../pages'

jest.mock('../../api/rest/fetchUrls', () => {
  return {
    getLoginURL: jest.fn().mockImplementation(() =>
      Promise.resolve({
        url: 'mockUrl',
        error: false,
      })
    )
  }
})



describe('Landing page tests', () => {

  it('should match the snapshot', () => {
    const tree = renderer.create(<Landing url={'moc url'} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
  it('should get login url if user session not available', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
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

  it('should redirect to error page if network error', async() => {
    
  })

  it('should redirect to github sign up page on Sing Up button click', async() => {
    
  })
  
})
