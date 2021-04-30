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
    state:0,
  });

  useEffect(() => {
    if (getStorage("token"))
      setState(getStorage("token"))
      else setState({
        accessToken: "",
        idToken: "",
        userId: "",
        provider: "",
        state:-1,
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
      state:-1,
    })
  }
  // const router = useRouter();
  // useEffect(() => {
  //   if (user.idToken == "")
  //     router.push("/")
  // }, [])
console.log(user)

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


// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkxkalJVcjE1NGQxb0RSSklielNYbiJ9.eyJuaWNrbmFtZSI6InJqb3NoaS1kZXEiLCJuYW1lIjoiUmFqYXJhbSBKb3NoaSIsInBpY3R1cmUiOiJodHRwczovL2F2YXRhcnMuZ2l0aHVidXNlcmNvbnRlbnQuY29tL3UvNzk5MDEzODI_dj00IiwidXBkYXRlZF9hdCI6IjIwMjEtMDQtMjlUMDc6NDM6MTcuMzQxWiIsImlzcyI6Imh0dHBzOi8vYWxmcmVkLXNoLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnaXRodWJ8Nzk5MDEzODIiLCJhdWQiOiJCRm5mZGFpYktTZHFrU0FPa3NyM1h1VU5KdUNXOXpiWiIsImlhdCI6MTYxOTc2MjkyNSwiZXhwIjoxNjE5Nzk4OTI1fQ.kvD1JoJwu39B9_33Etx56isPnOoLiDKNNZ2YbnADxL3NYCZdlkkGesYRMriqq2UTCOrw_sz7u1cOO-vBO_5m-oKStUmup6_dCaVbtXZ37lp8tEG-kgpyZudFOV4UtAS7GTBV-Eynb5BExjRZTwDMIoibi5a6P_lq75JFnOYadSMTX3DGhKg7rEEVW2izoPW7P_S9UDrobcHYgGip4Fym-i16kxhx2GelBb389VQECRICm-lP6Hft5iirDyX57nmLhmL68C1cIUzFaFUYOcdeMR8S7OpPczqgTjnETvado_R3w4RbMX1o3cMTyKXVONLqxGwzLa7LY-b0-d4hAwem1Q