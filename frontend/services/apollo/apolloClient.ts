import { ApolloClient, InMemoryCache, NormalizedCacheObject, createHttpLink } from '@apollo/client'
import { setContext } from '@apollo/client/link/context'

import { User } from '../../contexts/UserContext'
import { GQLHTTP } from '../../utils/constants'

const httpLink = createHttpLink({
  uri: GQLHTTP,
})

const client = (user: User): ApolloClient<NormalizedCacheObject> => {
  if (user && user.idToken !== '') {
    const authLink = setContext((_, { headers }) => ({
      headers: {
        ...headers,
        Authorization: `Bearer ${user.idToken}`,
      },
    }))
    return new ApolloClient({
      link: authLink.concat(httpLink),
      cache: new InMemoryCache(),
    })
  } else
    return new ApolloClient({
      link: httpLink,
      cache: new InMemoryCache(),
    })
}
export default client
