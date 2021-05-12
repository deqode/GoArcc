export const SERVER = process.env.SERVER
export const CENTRIFUGO = process.env.CENTRIFUGO
export const GQLHTTP = process.env.GQLHTTP
export const sessionCongfig = {
  password: 'complex_password_at_least_32_characters_long',
  cookieName: 'alfred',
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
}
