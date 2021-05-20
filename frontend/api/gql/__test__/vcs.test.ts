import { ApolloQueryResult } from '@apollo/client'

import { getOwnerName } from '../vcs'

const mockresult: ApolloQueryResult<any> = {
  loading: false,
  error: undefined,
  networkStatus: 7,
  data: {
    response: { data: { vcsConnections: { vcs_connection: [{ user_name: 'user_name' }] } } },
  },
}

jest.mock('../../../apolloClient', () => {
  const mApolloClient = { query: jest.fn().mockImplementation(() => Promise.resolve(mockresult)) }
  return { client: jest.fn().mockImplementation(() => Promise.resolve(() => mApolloClient)) }
})

describe('test no token ', () => {
  it('get Owner Name', async () => {
    const response = await getOwnerName({ idToken: '', accountId: 'accountId' })
    expect(response).toStrictEqual({ error: true, ownerName: '' })
  })
})
