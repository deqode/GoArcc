import renderer from 'react-test-renderer'
import Landing from '../../pages'

describe('Check Landing page snapshot', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Landing url={'moc url'} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})
