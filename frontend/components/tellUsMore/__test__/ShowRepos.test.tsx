import TestRenderer from 'react-test-renderer'
import { MockedProvider } from '@apollo/client/testing'
import { waitFor } from '@testing-library/react'
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
    error: new Error('An error occurred'),
  },
]

const props = {
  userId: 'userId',
  accountId: 'accountId',
  setCurrentRepo: jest.fn(),
  provider: 'provider',
}
// TODO need to wait for load

it('should show circular progress and match snapshot if repo list is not available yet', async () =>
  null)

it('should render repo list and match the snapshot if repo list is available', async () => {
  let component: any
  const { act } = TestRenderer

  await new Promise((resolve) => setTimeout(resolve, 0))

  await act(async () => {
    component = TestRenderer.create(
      <MockedProvider mocks={mocks} addTypename={false}>
        <ShowRepos {...props} />
      </MockedProvider>
    )
  })

  await waitFor(() => {
    expect(component.toJSON()).toMatchSnapshot()
  })
})

it('should set current repo on select', async () => null)
