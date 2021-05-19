import TestRenderer from 'react-test-renderer'
import { MockedProvider } from '@apollo/client/testing'
import { waitFor } from '@testing-library/react'
import { GET_BRANCHES } from '../../../GraphQL/Query'
import ShowBranches from '../ShowBranches'

const mocks = [
  {
    request: {
      query: GET_BRANCHES,
      variables: {
        provider: 'GITHUB',
        ownerName: 'ownerName',
        repoName: 'repoName,',
        accountId: 'accountId',
      },
    },
    result: {
      data: {
        repository: {
          branches: ['Branch1', 'Branch2'],
          clone_url: 'Mockclone_url',
        },
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

it('should render branches list and match snapshot', async () => {
  let component: any
  const { act } = TestRenderer

  await new Promise((resolve) => setTimeout(resolve, 0))

  await act(async () => {
    component = TestRenderer.create(
      <MockedProvider mocks={mocks} addTypename={false}>
        <ShowBranches {...props} />
      </MockedProvider>
    )
  })

  await waitFor(() => {
    expect(component.toJSON()).toMatchSnapshot()
  })
})
