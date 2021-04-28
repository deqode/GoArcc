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

function MyApp({ Component, pageProps }: AppProps) {
  const [user, setState] = useState({
    accessToken: "",
    idToken: "",
    userId: "",
    provider:""
  });

  useEffect(() => {
    if (getStorage("token"))
      setState(getStorage("token"))
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
      provider:""
    })
  }

  return (
    <ApolloProvider client={client}>
      <UserContext.Provider value={{ user, setUser,removeUser }}>
        <Navbar />
        <Component {...pageProps} />
      </UserContext.Provider>
    </ApolloProvider >
  );
}

export default MyApp
