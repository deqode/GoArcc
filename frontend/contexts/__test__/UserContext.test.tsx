import { render, cleanup, act } from '@testing-library/react'

import { useUserContext, defaultUser, UserContext, defaultUserContext } from '../UserContext'

const User = ({ children }) => children(useUserContext(defaultUser))
const setup = (): UserContext => {
  const returnVal: UserContext = {
    removeUser: () => null,
    setUser: () => null,
    user: defaultUser,
  }
  render(
    <User>
      {(val: any) => {
        Object.assign(returnVal, val)
        return null
      }}
    </User>
  )
  return returnVal
}

describe('Test User Context', () => {
  afterEach(cleanup)
  it('should set default User context', () => {
    const context = setup()
    expect(context.user).toStrictEqual({ idToken: '', loggedIn: false })
  })
  it('should set new  User context', () => {
    const context = setup()
    act(() => {
      context.setUser({ idToken: 'MockId', loggedIn: true })
    })
    expect(context.user).toStrictEqual({ idToken: 'MockId', loggedIn: true })
  })
  it('should remove User context', () => {
    const context = setup()
    act(() => {
      context.removeUser()
    })
    expect(context.user).toStrictEqual({ idToken: '', loggedIn: false })
  })
  it('check default User context ', () => {
    expect(defaultUserContext.removeUser()).toBe(null)
  })
})
