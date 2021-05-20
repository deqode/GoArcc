import mockAxios from 'axios'

import { getAuth0Callback, getVCSConnectionGitHubCallback } from '../callbacks'

describe('call getAuth0Callback', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('get Auth0 callback', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() =>
      Promise.resolve({
        data: {
          userId: 'MockUserID',
          idToken: 'MockIdToken',
          accessToken: 'MockAccessToken',
        },
      })
    )
    //work
    const cb = await getAuth0Callback('mockCode', 'mockState')
    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(false)
  })

  it('get empty response AUTH0 cb', async () => {
    //work
    const cb = await getAuth0Callback('mockCode', 'mockState')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Data Not found')
  })

  it('Invalid types', async () => {
    //work
    const cb = await getAuth0Callback('', 'mockState')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(0)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Invalid types')
  })

  it('Network Error ', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() => Promise.reject({}))
    //work
    const cb = await getAuth0Callback('mockCode', 'mockState')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Network Error')
  })
})

describe('call getVCSConnectionGitHubCallback', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('get VCS callback', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() =>
      Promise.resolve({
        data: {
          accountId: 'MockAccountId',
        },
      })
    )
    //work
    const cb = await getVCSConnectionGitHubCallback('mockCode', 'mockAccountId', 'mockIdToken')
    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(false)
  })

  it('get empty response VCS cb', async () => {
    //work
    const cb = await getVCSConnectionGitHubCallback('mockCode', 'mockAccountId', 'mockIdToken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Data Not found')
  })

  it('Invalid types', async () => {
    //work
    const cb = await getVCSConnectionGitHubCallback('', 'mockAccountId', 'mockIdToken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(0)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Invalid types')
  })

  it('Network Error ', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() => Promise.reject({}))
    //work
    const cb = await getVCSConnectionGitHubCallback('mockCode', 'mockAccountId', 'mockIdToken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(cb.error).toEqual(true)
    expect(cb.message).toEqual('Network Error')
  })
})
