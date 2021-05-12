import { AppProps } from 'next/app'
import { ApolloProvider } from '@apollo/client'
import client from '../apolloClient'
import '../styles/globals.css'
import '../styles/bootstrap.min.css'
import '../styles/fonts.css'
import '../styles/alfred.css'
import UserContext, { useUserContext, defaultUser } from '../contexts/UserContext'
import Navbar from '../components/Navbar'
import { ReactElement } from 'react'

const MyApp = ({ Component, pageProps }: AppProps): ReactElement => {
  const { user, setUser, removeUser } = useUserContext(defaultUser)

  return (
    <UserContext.Provider value={{ user, setUser, removeUser }}>
      <ApolloProvider client={client(user)}>
        <Navbar />
        <Component {...pageProps} />
      </ApolloProvider>
    </UserContext.Provider>
  )
}

export default MyApp
