import renderer from 'react-test-renderer'
import Landing from '../../pages'

describe('Index page', () => {
  it('should match the snapshot', () => {
    const tree = renderer.create(<Landing url={'moc url'} />).toJSON()
    expect(tree).toMatchSnapshot()
  })
})
