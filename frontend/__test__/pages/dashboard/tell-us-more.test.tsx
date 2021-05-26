import { MockedProvider } from '@apollo/client/testing'
import { act, create, ReactTestRenderer } from 'react-test-renderer'

import * as api1 from '../../../api/rest/user'
import { GET_BRANCHES, GET_REPOSITORIES } from '../../../GraphQL/Query'
import { UserResponse } from '../../../intefaces/interface'
import TellUsMore, { handler } from '../../../pages/dashboard/tell-us-more'
import { redirectToLandingPage } from '../../../utils/redirects'

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

// jest.mock('../../../api/rest/user', () => {
//   return {
//     getAllUserAccounts: jest.fn().mockImplementation(() =>
//       Promise.resolve({
// accounts: [{ id: 'mockId', slug: 'mockslug', userId: 'mockuserId' }],
//       })
//     ),
//   }
// })
jest.mock('../../../api/gql/vcs', () => {
  return {
    getOwnerName: jest.fn().mockImplementation(() =>
      Promise.resolve({
        error: false,
        ownerName: 'mockownerName',
      })
    ),
  }
})
const mockUser: UserResponse = {
  accessToken: 'mockAccessToken',
  idToken: 'mockidToken',
  userId: 'mockuserId',
}
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

  it('should return to landing page if session not available', async () => {
    const req = {
      session: { get: jest.fn(() => undefined), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual(redirectToLandingPage())
  })

  it('should redirect to error page if get all user accounts fails', async () => {
    api1.getAllUserAccounts = jest.fn().mockImplementation(() =>
      Promise.resolve({
        accounts: [],
        error: true,
      })
    )
    const req = {
      session: { get: jest.fn(() => mockUser), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({
      redirect: { permanent: false, destination: '/error?message=Network Error' },
    })
  })

  it('should redirect to dashboard if get owner fails', async () => null)

  it('should match the snapshot if nothing fails', async () => null)
  it('should return props when serverSideProps handler called ', async () => {
    api1.getAllUserAccounts = jest.fn().mockImplementation(() =>
      Promise.resolve({
        accounts: [{ id: 'mockId', slug: 'mockslug', userId: 'mockuserId' }],
        error: false,
      })
    )
    const req = {
      session: { get: jest.fn(() => mockUser), save: jest.fn() },
    }
    const response = await handler({ req })
    expect(response).toStrictEqual({
      props: {
        userID: 'mockuserId',
        accountId: 'mockId',
        user: {
          accessToken: 'mockAccessToken',
          idToken: 'mockidToken',
          userId: 'mockuserId',
        },
        ownerName: 'mockownerName',
      },
    })
  })
})
