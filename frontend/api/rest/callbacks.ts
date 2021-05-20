import { ResponseError, UserResponse } from '../../intefaces/interface'
import { get } from '../../services/rest/get'
import { SERVER } from '../../utils/constants'

interface UserAuthResponse extends ResponseError<UserAuthResponse>, UserResponse {}

export const getAuth0Callback = async (code: string, state: string): Promise<UserAuthResponse> => {
  if (code !== '' && state !== '') {
    const response = await get<UserAuthResponse>(
      `${SERVER}/authentication/callback?code=${code}&state=${state}`
    )
    if (response.error) {
      return { accessToken: '', idToken: '', userId: 'string', ...response }
    }
    if (response.data && response.data.userId) {
      return {
        userId: response.data.userId,
        idToken: response.data.idToken,
        accessToken: response.data.idToken,
        error: false,
        message: '',
      }
    } else
      return {
        userId: '',
        idToken: '',
        accessToken: '',
        error: true,
        message: 'Data Not found',
        // TODO : integrate with error message from backend
      }
  } else
    return {
      userId: '',
      idToken: '',
      accessToken: '',
      error: true,
      message: 'Invalid types',
      // TODO : integrate with error message from backend
    }
}

export const getVCSConnectionGitHubCallback = async (
  code: string,
  accountId: string,
  idToken: string
): Promise<ResponseError<any>> => {
  //TODO:any?
  if (code !== '' && accountId !== '') {
    const response = await get<ResponseError<any>>(
      `${SERVER}/vcs-connection/GITHUB/callback?code=${code}&account_id=${accountId}`,
      { Authorization: `Bearer ${idToken}` }
    )
    if (response.error) return response
    if (response.data?.data) {
      return {
        error: false,
        message: '',
      }
    } else
      return {
        error: true,
        message: 'Data Not found',
        // TODO : integrate with error message from backend
      }
  } else
    return {
      error: true,
      message: 'Invalid types',
      // TODO : integrate with error message from backend
    }
}
