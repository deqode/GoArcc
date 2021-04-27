import { ApolloClient } from "@apollo/client";
import cache from "./GraphQL/cache";

const client = new ApolloClient({
  uri: "http://localhost:8081/graphql",
  cache: cache,
});

export default client;