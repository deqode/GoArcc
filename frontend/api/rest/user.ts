import { ResponseError, UserAccount } from '../../intefaces/interface'
import { get } from '../../services/rest'
import { SERVER } from '../../utils/constants'

export interface AllUserAccounts extends ResponseError<AllUserAccounts> {
  accounts: Array<UserAccount>
}

export const getAllUserAccounts = async (
  userId: string,
  idToken: string
): Promise<AllUserAccounts> => {
  // TODO:valide inputs
  const response = await get<AllUserAccounts>(`${SERVER}/account/get-user-accounts/${userId}`, {
    Authorization: `Bearer ${idToken}`,
  })
  if (response.error) {
    return { accounts: [], ...response }
  }
  if (response.data && response.data.accounts) {
    return {
      error: false,
      message: '',
      accounts: response.data.accounts,
    }
  } else
    return {
      error: true,
      message: 'Data Not found',
      accounts: [],
    }
}
