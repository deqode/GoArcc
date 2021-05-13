import { AppProps } from 'next/app'
import { ApolloProvider } from '@apollo/client'
import client from '../apolloClient'
import UserContext, { useUserContext, defaultUser } from '../contexts/UserContext'
import Navbar from '../components/Navbar'
import { ReactElement, useEffect } from 'react'
import { ThemeProvider } from '@material-ui/styles'
import theme from '../styles/Theme'
import Head from 'next/head'

const MyApp = ({ Component, pageProps }: AppProps): ReactElement => {
  const { user, setUser, removeUser } = useUserContext(defaultUser)
  useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    if (jssStyles && jssStyles.parentElement) {
      jssStyles.parentElement.removeChild(jssStyles)
    }
  }, [])

  return (
    <UserContext.Provider value={{ user, setUser, removeUser }}>
      <Head>
        <title>My page</title>
        <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
      </Head>
      <ApolloProvider client={client(user)}>
        <ThemeProvider theme={theme}>
          <Navbar />
          <Component {...pageProps} />
        </ThemeProvider>
      </ApolloProvider>
    </UserContext.Provider>
  )
}

export default MyApp
