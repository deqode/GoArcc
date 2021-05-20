import mockAxios from 'axios'

import { UserResponse } from '../../../interface'
import { destroyUserSession, setUserSession } from '../session'

describe('call setUserSession', () => {
  const mockUser: UserResponse = {
    accessToken: 'mockAccessToken',
    idToken: 'mockidToken',
    userId: 'mockuserId',
  }
  afterEach(() => {
    jest.clearAllMocks()
  })
  it('setUserSession success', async () => {
    //setup
    mockAxios.post.mockImplementationOnce(() =>
      Promise.resolve({
        data: { success: true },
      })
    )
    //work
    const userAccount = await setUserSession(mockUser)
    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(false)
  })
  it('get empty response login URL', async () => {
    //work
    const userAccount = await setUserSession(mockUser)

    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Server Error')
  })
  it('Network Error ', async () => {
    //setup
    mockAxios.post.mockImplementationOnce(() => Promise.reject({}))
    //work
    const userAccount = await setUserSession(mockUser)

    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Network Error')
  })
})

describe('call destroyUserSession', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })
  it('destroyUserSession success', async () => {
    //setup
    mockAxios.post.mockImplementationOnce(() =>
      Promise.resolve({
        data: { success: true },
      })
    )
    //work
    const userAccount = await destroyUserSession()
    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(false)
  })
  it('get empty response ', async () => {
    //work
    const userAccount = await destroyUserSession()

    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Server Error')
  })
  it('Network Error ', async () => {
    //setup
    mockAxios.post.mockImplementationOnce(() => Promise.reject({}))
    //work
    const userAccount = await destroyUserSession()

    //conclusion
    expect(mockAxios.post).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Network Error')
  })
})
