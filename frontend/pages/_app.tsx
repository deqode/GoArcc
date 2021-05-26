import { ApolloProvider } from '@apollo/client'
import { ThemeProvider } from '@material-ui/styles'
import { AppProps } from 'next/app'
import { ReactElement, useEffect } from 'react'

import Navbar from '../components/Navbar'
import UserContext, { useUserContext, defaultUser } from '../contexts/UserContext'
import client from '../services/apollo/apolloClient'
import theme from '../styles/Theme'

const MyApp = ({ Component, pageProps }: AppProps): ReactElement => {
  const { user, setUser, removeUser } = useUserContext(defaultUser)
  useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    if (jssStyles && jssStyles.parentElement) {
      jssStyles.parentElement.removeChild(jssStyles)
    }
  }, [])
  useEffect(() => {
    if (pageProps.user?.idToken?.length)
      setUser({
        idToken: pageProps.user.idToken,
        loggedIn: true,
      })
    else
      setUser({
        idToken: '',
        loggedIn: false,
      })
  }, [pageProps])
  return (
    <UserContext.Provider value={{ user, setUser, removeUser }}>
      <ApolloProvider client={client(user)}>
        <ThemeProvider theme={theme}>
          <Navbar user={user} />
          <Component {...pageProps} />
        </ThemeProvider>
      </ApolloProvider>
    </UserContext.Provider>
  )
}

export default MyApp
