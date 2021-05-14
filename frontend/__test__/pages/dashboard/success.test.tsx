import renderer from 'react-test-renderer'
import Success from '../../../pages/dashboard/success'

describe('Check Success  page snapshot', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Success />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})
