import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client'
// import { setContext } from '@apollo/client/link/context';
// import { getStorage } from "./utils/localStorage";

const httpLink = createHttpLink({
  uri: 'http://localhost:8081/graphql',
})

// const authLink = setContext((_, { headers }) => {
//   return {
//     headers: {
//       ...headers,
//       Authorization: `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkxkalJVcjE1NGQxb0RSSklielNYbiJ9.eyJuaWNrbmFtZSI6InJqb3NoaS1kZXEiLCJuYW1lIjoiUmFqYXJhbSBKb3NoaSIsInBpY3R1cmUiOiJodHRwczovL2F2YXRhcnMuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3UvNzk5MDEzODI_dj00IiwidXBkYXRlZF9hdCI6IjIwMjEtMDUtMDZUMjA6Mjg6NTAuOTgyWiIsImlzcyI6Imh0dHBzOi8vYWxmcmVkLXNoLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnaXRodWJ8Nzk5MDEzODIiLCJhdWQiOiJCRm5mZGFpYktTZHFrU0FPa3NyM1h1VU5KdUNXOXpiWiIsImlhdCI6MTYyMDMzMzQwNCwiZXhwIjoxNjIwMzY5NDA0fQ.MZxjUBQTkHvwLxkWZNQ90kcjK8q2dsyHSDQ3Zhr7tWmiuHXYUCQJ1UJhrqM2jLzn-lxPLgsXuUA1k_ynC8rBWM1GkGhW6l53okptdyYfybILR9-HApjwvuNEHp291BbMPClYfKikkxljeOi6acbIgtnPNpYH8FEAjMlyvXnaM_kfgn31zCMhvmUPKr2bDE4e3alzSMduj74eP2I_j_kCLMq_x401fGmbNhXLwJyLkrWXq10owc_n6iQJSD8dU0Q2k0GxPcwDAUYfxjdXRN7748w3QDpeQULS2cdlUZMI_TeJKb37PKTvPbINs_15TKeHRNXkoUhqVNCdyLugHl4aPQ`,
//     }
//   }
// });
const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
})

export default client

// https://www.apollographql.com/docs/react/networking/authentication/
