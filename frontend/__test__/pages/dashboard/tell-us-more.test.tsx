import { MockedProvider } from '@apollo/client/testing'
import { act, create, ReactTestRenderer } from 'react-test-renderer'
import { GET_BRANCHES, GET_REPOSITORIES } from '../../../GraphQL/Query'
import TellUsMore from '../../../pages/dashboard/tell-us-more'
const props = {
  user: { accessToken: 'accessToken', idToken: 'idToken', userId: 'userId' },
  accountId: 'accountId',
  ownerName: 'accountId',
}

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
  {
    request: {
      query: GET_BRANCHES,
      variables: {
        provider: 'provider',
        ownerName: 'ownerName',
        repoName: 'repoName,',
        accountId: 'accountId',
      },
    },
    result: {
      data: {
        repository: {
          branches: ['Branch1', 'Branch1'],
        },
        clone_url: 'Mockclone_url',
      },
    },
  },
]

//TODO: check if gql queries are mocked
describe('dashboard/tell-us-more page tests', () => {
  let wrapper: ReactTestRenderer

  beforeEach(() => {
    act(() => {
      wrapper = create(
        <MockedProvider mocks={mocks} addTypename={false}>
          <TellUsMore {...props} />
        </MockedProvider>
      )
    })
  })
  it('should match the snapshot', () => {
    expect(wrapper.toJSON()).toMatchSnapshot()
  })
})
