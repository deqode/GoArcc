import { ApolloClient, InMemoryCache } from "@apollo/client";
import cache from "./GraphQL/cache";

const client = new ApolloClient({
  uri: "http://localhost:8081/graphql",
  cache: new InMemoryCache(),
});

export default client;