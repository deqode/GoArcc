import { AppProps } from 'next/app';
import { ApolloProvider } from "@apollo/client";
import client from "../apollo-client";
import '../styles/globals.css';
import '../styles/bootstrap.min.css';
import '../styles/fonts.css';
import '../styles/alfred.css';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default MyApp
