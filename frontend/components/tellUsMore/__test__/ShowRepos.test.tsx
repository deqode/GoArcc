import { MockedProvider } from '@apollo/client/testing'
import { act, create, ReactTestRenderer } from 'react-test-renderer'
import { GET_REPOSITORIES } from '../../../GraphQL/Query'
import ShowRepos from '../ShowRepos'

const mocks = [
  {
    request: {
      query: GET_REPOSITORIES,
      variables: {
        provider: 'provider',
        userId: 'userId',
        accountId: 'accountId',
      },
    },
    result: {
      data: {
        repositories: {
          repositories: ['REPO1', 'REPO2'],
        },
      },
    },
  },
]

const props = {
  userId: 'userId',
  accountId: 'accountId',
  setCurrentRepo: jest.fn(),
  provider: 'provider',
}
// TODO need to wait for load

describe('PSAButton', () => {
  let wrapper: ReactTestRenderer

  beforeEach(async () => {
    await act(async () => {
      wrapper = create(
        <MockedProvider mocks={mocks} addTypename={false}>
          <ShowRepos {...props} />
        </MockedProvider>
      )
    })
  })

  it('renders correctly', async () => {
    expect(wrapper.toJSON()).toMatchSnapshot()
  })
})
