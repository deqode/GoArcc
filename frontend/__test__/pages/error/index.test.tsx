import renderer from 'react-test-renderer'
import ErrorPage from '../../../pages/error/index'

jest.mock('next/router', () => ({
  useRouter: jest.fn().mockImplementation(() => ({
    query: {
      message: 'Network Error',
    },
  })),
}))

describe('Check Error page snapshot', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<ErrorPage />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})