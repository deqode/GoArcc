import renderer from 'react-test-renderer'
import Dashboard from '../../pages/dashboard'
import Success from '../../pages/dashboard/success'

describe('dashboard page', () => {
  it('should match the snapshot', () => {
    const tree = renderer
      .create(
        <Dashboard user={{ accessToken: 'mockToken', idToken: 'mockToken', userId: 'mockId' }} />
      )
      .toJSON()
    expect(tree).toMatchSnapshot()
  })
})

describe('dashboard/success page', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Success />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})
