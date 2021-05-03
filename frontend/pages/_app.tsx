import { AppProps } from 'next/app';
import { ApolloProvider } from "@apollo/client";
import client from "../apollo-client";
import '../styles/globals.css';
import '../styles/bootstrap.min.css';
import '../styles/fonts.css';
import '../styles/alfred.css';
import Navbar from '../components/Navbar';
import { UserContext } from '../Contexts/UserContext';
import { useEffect, useState } from 'react';
import { getStorage, removeStorage, setStorage } from '../utils/localStorage';
import { useRouter } from 'next/router';
import Login from '.';
import Success from './success';

function MyApp({ Component, pageProps }: AppProps) {
  const [user, setState] = useState({
    accessToken: "",
    idToken: "",
    userId: "",
    provider: "",
    accounts: [],
    state: 0,
  });

  useEffect(() => {
    if (getStorage("token"))
      setState(getStorage("token"))
    else setState({
      accessToken: "",
      idToken: "",
      userId: "",
      provider: "",
      accounts: [],
      state: -1,
    })
  }, [])

  const setUser = (value) => {
    setStorage("token", value)
    setState(value)
  }

  const removeUser = () => {
    removeStorage("token")
    setState({
      accessToken: "",
      idToken: "",
      userId: "",
      provider: "",
      accounts: [],
      state: -1,
    })
  }
  return (
    <ApolloProvider client={client}>
      <UserContext.Provider value={{ user, setUser, removeUser }}>
        <Navbar />
        <Component {...pageProps} />
      </UserContext.Provider>
    </ApolloProvider >
  );
}

export default MyApp


