import mockAxios from 'axios'
import { getAllUserAccounts } from '../user'
describe('call getAllUserAccounts', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })
  it('get all User Accounts', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() =>
      Promise.resolve({
        data: { accounts: ['Account1', 'Account2'] },
      })
    )
    //work
    const userAccount = await getAllUserAccounts('MockUserId', 'MockIdToken')
    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(false)
  })
  it('get empty response login URL', async () => {
    //work
    const userAccount = await getAllUserAccounts('MockUserId', 'MockIdToken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Data Not found')
  })
  it('Network Error ', async () => {
    //setup
    mockAxios.get.mockImplementationOnce(() => Promise.reject({}))
    //work
    const userAccount = await getAllUserAccounts('MockUserId', 'MockIdToken')

    //conclusion
    expect(mockAxios.get).toHaveBeenCalledTimes(1)
    expect(userAccount.error).toEqual(true)
    expect(userAccount.message).toEqual('Network Error')
  })
})
