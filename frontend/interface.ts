export interface UserContextInterface {
  user: {
    accessToken: string
    idToken: string
    userId: string
    provider: string
    accounts: Array<any>
    state: number
    //  loading=0
    //  authenticated=1
    //  unauthenticated=-1
  }
  setUser(value: any): void
  removeUser(): void
}

export interface AppContextInterface {
  app: {
    centrifuge: any
    subscribed: boolean
  }
  setApp(value: any): void
  removeApp(): void
}
