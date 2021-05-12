import axios from 'axios'
import { ResponseError, UserAccount } from '../../interface'
import { SERVER } from '../../utils/constants'

export const getAllUserAccounts = async (
  userId: string,
  idToken: string
): Promise<AllUserAccounts> => {
  try {
    const response = await axios.get(`${SERVER}/account/get-user-all-account/${userId}`, {
      headers: { Authorization: `Bearer ${idToken}` },
    })
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
  } catch (e) {
    return {
      error: true,
      message: 'Network Error',
      accounts: [],

      // TODO : integrate with error message from backend
    }
  }
}

interface AllUserAccounts extends ResponseError {
  accounts: Array<UserAccount>
}
