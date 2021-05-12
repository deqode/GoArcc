import client from '../../apolloClient'
import { GET_OWNER_NAME } from '../../GraphQL/Query'

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
        ownerName: response.data.VCSConnections.vcs_connection[0].user_name,
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
