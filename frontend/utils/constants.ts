export const SERVER = 'http://localhost:8082/v1'
export const CENTRIFUGO = 'ws://localhost:7070/connection/websocket'
export const GQLHTTP = 'http://localhost:8081/graphql'
export const sessionCongfig = {
  password: 'complex_password_at_least_32_characters_long',
  cookieName: 'alfred',
  cookieOptions: {
    secure: process.env.NODE_ENV === 'production',
  },
}
