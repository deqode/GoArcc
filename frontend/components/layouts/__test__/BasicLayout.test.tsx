import { Button } from '@material-ui/core'
import renderer from 'react-test-renderer'

import BasicLayout from '../BasicLayout'

describe('basic layout component tests', () => {
  it('should match the snapshot', () => {
    const tree = renderer
      .create(
        <BasicLayout
          heading="This is basicLayout headding"
          subHeading="This is basicLayout sub-headding"
          component={
            <Button variant="contained" color="primary" href={'https://github.com/'}>
              SignUP
            </Button>
          }
        />
      )
      .toJSON()
    expect(tree).toMatchSnapshot()
  })
})
