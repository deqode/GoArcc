import { GET_OWNER_NAME } from '../../GraphQL/Query'
import client from '../../services/apollo/apolloClient'

export const getOwnerName = async ({
  idToken,
  accountId,
}: {
  idToken: string
  accountId: string
}): Promise<{ error: boolean; ownerName: string }> => {
  if (idToken !== '') {
    const response = await client({ idToken, loggedIn: false }).query({
      query: GET_OWNER_NAME,
      variables: {
        accountId: accountId,
      },
    })
    if (!response.error)
      return {
        error: false,
        ownerName: response.data.vcsConnections.vcs_connection[0].user_name,
      }
    return {
      error: true,
      ownerName: '',
    }
  }
  return {
    error: true,
    ownerName: '',
  }
}
