export interface UserResponse {
  accessToken: string
  idToken: string
  userId: string
}

export interface ResponseError {
  error: boolean
  message: string
}

export interface UserAccount {
  id: string
  slug: string
  userId: string
}
