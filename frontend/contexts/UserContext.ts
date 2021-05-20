import { createContext, Dispatch, SetStateAction, useCallback, useState } from 'react'

export interface User {
  idToken: string
  loggedIn: boolean
}
export interface UserContext {
  user: User
  setUser: Dispatch<SetStateAction<User>>
  removeUser: () => void
}

export const defaultUser: User = { idToken: '', loggedIn: false }

export const defaultUserContext: UserContext = {
  user: defaultUser,
  setUser: () => null,
  removeUser: () => null,
}

export const useUserContext = (initialState: User): UserContext => {
  const [user, setUser] = useState<User>(initialState)
  const removeUser = useCallback(() => {
    setUser(defaultUser)
  }, [])
  return { user, setUser, removeUser }
}

const UserContext = createContext<UserContext>(defaultUserContext)

export default UserContext
