export const SERVER = process.env.NEXT_PUBLIC_SERVER
export const CENTRIFUGO = process.env.NEXT_PUBLIC_CENTRIFUGO
export const GQLHTTP = process.env.NEXT_PUBLIC_GQLHTTP
declare let process: {
  env: {
    NODE_ENV: string
    NEXT_PUBLIC_SESSION_PASSWORD: string
    NEXT_PUBLIC_GQLHTTP: string
    NEXT_PUBLIC_CENTRIFUGO: string
    NEXT_PUBLIC_SERVER: string
  }
}
export const sessionCongfig = {
  password: process.env.NEXT_PUBLIC_SESSION_PASSWORD,
  cookieName: 'alfred',
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
}
