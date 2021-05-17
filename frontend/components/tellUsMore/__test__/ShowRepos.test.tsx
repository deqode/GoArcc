import { MockedProvider } from '@apollo/client/testing'
import { act, create, ReactTestRenderer } from 'react-test-renderer'
import { GET_REPOSITORIES } from '../../../GraphQL/Query'
import ShowRepos from '../ShowRepos'
import waitForExpect from 'wait-for-expect'

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
    wrapper = create(
      <MockedProvider mocks={mocks} addTypename={false}>
        <ShowRepos {...props} />
      </MockedProvider>
    )
  })

  it('renders correctly', async () => {
    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 0))

      await waitForExpect(() => {
        expect(wrapper.toJSON()).toMatchSnapshot()
      })
    })
  })
})
