import mockAxios from 'axios'

import { getGithubVCSConnection, getLoginURL } from '../fetchUrls'

describe('call getLoginURL', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })
  it('get Auth0 login URL', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() =>
      Promise.resolve({
        data: { url: 'mockURL' },
      })
    )
    //work
    const loginURL = await getLoginURL()
    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(false)
  })
  it('get empty response', async () => {
    //work
    const loginURL = await getLoginURL()

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(true)
    expect(loginURL.message).toEqual('Data Not found')
  })
  it('Network Error ', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() => Promise.reject({}))
    //work
    const loginURL = await getLoginURL()

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(true)
    expect(loginURL.message).toEqual('Network Error')
  })
})

describe('call getGithubVCSConnection', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('get VCS login URL', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() =>
      Promise.resolve({
        data: { redirectUrl: 'mockURL' },
      })
    )
    //work
    const loginURL = await getGithubVCSConnection('mockIDtoken')
    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(false)
  })

  it('get empty response VCS URL', async () => {
    //work
    const loginURL = await getGithubVCSConnection('mockIDtoken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(true)
    expect(loginURL.message).toEqual('Data Not found')
  })

  it('Network Error ', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() => Promise.reject({}))
    //work
    const loginURL = await getGithubVCSConnection('mockIDtoken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(loginURL.error).toEqual(true)
    expect(loginURL.message).toEqual('Network Error')
  })
})
