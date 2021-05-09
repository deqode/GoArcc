import { AppProps } from 'next/app'
import { ApolloProvider } from '@apollo/client'
import client from '../apollo-client'
import '../styles/globals.css'
import '../styles/bootstrap.min.css'
import '../styles/fonts.css'
import '../styles/alfred.css'
import Navbar from '../components/Navbar'
import { UserContext } from '../Contexts/UserContext'
import { AppContext } from '../Contexts/AppContext'
import { useEffect, useState } from 'react'
import { getStorage, removeStorage, setStorage } from '../utils/localStorage'

import Centrifuge from 'centrifuge'
import { CENTRIFUGO } from '../utils/constants'

function MyApp({ Component, pageProps }: AppProps): any {
  const [user, setState] = useState({
    accessToken: '',
    idToken: '',
    userId: '',
    provider: '',
    accounts: [],
    state: 0,
  })

  const [app, setAppState] = useState({
    centrifuge: new Centrifuge(CENTRIFUGO || ''),
    subscribed: false,
  })

  useEffect(() => {
    if (getStorage('token')) setState(getStorage('token'))
    else
      setState({
        accessToken: '',
        idToken: '',
        userId: '',
        provider: '',
        accounts: [],
        state: -1,
      })
  }, [])

  const setUser = (value: any): void => {
    // setStorage("token", value)
    setState(value)
  }

  const removeUser = (): void => {
    removeStorage('token')
    setState({
      accessToken: '',
      idToken: '',
      userId: '',
      provider: '',
      accounts: [],
      state: -1,
    })
  }

  const setApp = (value: any): void => {
    setStorage('app', value)
    setAppState(value)
  }

  const removeApp = (): void => {
    removeStorage('app')
    setAppState({
      centrifuge: new Centrifuge(CENTRIFUGO),
      subscribed: false,
    })
  }

  return (
    <ApolloProvider client={client}>
      <UserContext.Provider value={{ user, setUser, removeUser }}>
        <AppContext.Provider value={{ app, setApp, removeApp }}>
          <Navbar />
          <Component {...pageProps} />
        </AppContext.Provider>
      </UserContext.Provider>
    </ApolloProvider>
  )
}

export default MyApp
