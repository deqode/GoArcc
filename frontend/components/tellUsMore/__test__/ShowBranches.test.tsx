import { MockedProvider } from '@apollo/client/testing'
import { act, create, ReactTestRenderer } from 'react-test-renderer'
import { GET_BRANCHES } from '../../../GraphQL/Query'
import ShowBranches from '../ShowBranches'

const mocks = [
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

const props = {
  accountId: 'accountId',
  provider: 'GITHUB',
  ownerName: 'MockownerName',
  repoName: 'MockrepoName',
  setBranchName: jest.fn(),
  setCloneUrl: jest.fn(),
}

describe('PSAButton', () => {
  let wrapper: ReactTestRenderer

  beforeEach(() => {
    act(() => {
      wrapper = create(
        <MockedProvider mocks={mocks} addTypename={false}>
          <ShowBranches {...props} />
        </MockedProvider>
      )
    })
  })

  it('renders correctly', async () => {
    expect(wrapper.toJSON()).toMatchSnapshot()
  })
})
