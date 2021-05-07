import { ApolloClient, createHttpLink, HttpLink, InMemoryCache } from "@apollo/client";
import { setContext } from '@apollo/client/link/context';
import { getStorage } from "./utils/localStorage";

const httpLink = createHttpLink({
  uri: 'http://localhost:8081',
});

const authLink = setContext((_, { headers }) => {
  let token=getStorage("token").idToken || ""

  return {
    headers: {
      ...headers,
      Authorization: `Bearer ${token}`,
    }
  }
});
const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default client;

// https://www.apollographql.com/docs/react/networking/authentication/