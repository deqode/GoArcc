export const SERVER = process.env.NEXT_PUBLIC_SERVER
export const CENTRIFUGO = process.env.NEXT_PUBLIC_CENTRIFUGO
export const GQLHTTP = process.env.NEXT_PUBLIC_GQLHTTP
declare let process: {
  env: {
    NODE_ENV: string
    SESSION_PASSWORD: string
    NEXT_PUBLIC_GQLHTTP: string
    NEXT_PUBLIC_CENTRIFUGO: string
    NEXT_PUBLIC_SERVER: string
    COOKIE_NAME: string
  }
}
export const sessionCongfig = {
  password: process.env.SESSION_PASSWORD,
  cookieName: process.env.COOKIE_NAME,
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
}
