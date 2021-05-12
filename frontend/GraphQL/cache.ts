import { gql, InMemoryCache, makeVar } from '@apollo/client'

export const token = makeVar(null)

export default new InMemoryCache({
  typePolicies: {
    Query: {
      fields: {
        token: {
          read() {
            return token()
          },
        },
      },
    },
  },
})

export const GET_TOKEN = gql`
  query getToken {
    token @client
  }
`
