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
        provider: 'provider',
        ownerName: 'ownerName',
        repoName: 'repoName',
        accountId: 'accountId',
      },
    },
    error: new Error('An error occurred'),
  },
]

const props = {
  accountId: 'accountId',
  provider: 'GITHUB',
  ownerName: 'ownerName',
  repoName: 'repoName',
  setBranchName: jest.fn(),
  setCloneUrl: jest.fn(),
}

it('should show circular progress and match snapshot, if branch list is not available yet', async() =>{

})

it('should not show any branch to select if repo name not yet set', async() =>{

})

it('should render branches list and match snapshot if repo name is set', async () => {
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

it('should set current branch on select', async() => {
  
})
